package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gateway/pkg/registry"
	"gateway/pkg/transport"

	"github.com/binbinly/pkg/logger"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// App global app
type App struct {
	opts   options
	ctx    context.Context
	cancel func()
	mu     sync.Mutex
	svc    *registry.Service
}

// New create a app globally
func New(opts ...Option) *App {
	o := options{
		registerTTL:      registerTTL,
		registerInterval: registerInterval,
		services:         services,
	}
	if id, err := uuid.NewUUID(); err == nil {
		o.id = id.String()
	}
	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		opts:   o,
		ctx:    ctx,
		cancel: cancel,
	}
}

// Run 启动api网关服务
func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)

	// start server
	go func() {
		<-a.ctx.Done()
		if err := a.opts.server.Stop(a.ctx); err != nil {
			log.Printf("[HTTP] stop err: %v", err)
		}
	}()
	go func() {
		if err := a.opts.server.Start(a.ctx); err != nil {
			log.Fatalf("[HTTP] Failed to listen and serve: %v", err)
		}
	}()

	if err := a.Register(); err != nil {
		return err
	}

	a.RegisterHandlers()

	// watch signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	eg.Go(func() error {
		defer log.Println("signal defer")
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case s := <-quit:
				log.Printf("Server receive a quit signal: %s", s.String())
				if err := a.Stop(); err != nil {
					log.Printf("failed to stop app, err: %v", err)
					return err
				}
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

// Register 服务注册
func (a *App) Register() error {
	// register service
	if a.opts.registry != nil {
		var address string
		if r, ok := a.opts.server.(transport.Endpoint); ok {
			e, err := r.Endpoint()
			if err != nil {
				return err
			}
			address, err = parseAddress(e.String())
			if err != nil {
				return err
			}
		}

		node := &registry.Node{
			Id:      a.opts.id,
			Address: address,
		}
		service := &registry.Service{
			Name:     a.opts.name,
			Version:  a.opts.version,
			Metadata: a.opts.metadata,
			Nodes:    []*registry.Node{node},
		}
		if err := a.opts.registry.Register(service); err != nil {
			return err
		}
		a.svc = service

		go func() {
			t := time.NewTicker(a.opts.registerInterval)

			for {
				select {
				case <-t.C:
					if err := a.opts.registry.Register(service); err != nil {
						logger.Warn("Server register error: ", err)
					}
				}
			}
		}()
	}

	return nil
}

// RegisterHandlers 注册代理的全部服务处理器
func (a *App) RegisterHandlers() {
	for _, service := range a.opts.services {
		conn := newRPCClientConn(a.ctx, service.Name)
		if err := service.Handler(a.ctx, a.opts.mux, conn); err != nil {
			log.Fatalf("register %v handler err: %v", service.Name, err)
		}
		log.Printf("register %v handler success", service.Name)
	}
}

// Stop stops the application gracefully.
func (a *App) Stop() error {
	if a.opts.registry != nil {
		if err := a.opts.registry.Deregister(a.svc); err != nil {
			return err
		}
	}

	// cancel app
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

// parseAddress 解析 address
func parseAddress(endpoint string) (string, error) {
	raw, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", raw.Hostname(), raw.Port()), nil
}
