package app

import (
	"context"
	"github.com/urfave/cli/v2"
	"os"
	"pkg/handler"
	pb "pkg/proto/core"
	"pkg/wrap"
	"pkg/wrap/recovery"
	"sync"
	"time"

	ot "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	"github.com/go-micro/cli/debug/trace/jaeger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/consul"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

// App global app
type App struct {
	srv    micro.Service
	opts   options
	ctx    context.Context
	cancel func()
}

// New create a app globally
func New(opts ...Option) *App {
	o := options{}
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

func (a *App) Init(opts ...micro.Option) {
	// Create tracer
	tracer, closer, err := jaeger.NewTracer(
		jaeger.Name(a.opts.name),
		jaeger.FromEnv(true),
		jaeger.GlobalTracer(true),
	)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	mo := []micro.Option{
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.BeforeStart(func() error {
			logger.Infof("Starting service %s", a.opts.name)
			return nil
		}),
		micro.BeforeStop(func() error {
			logger.Infof("Shutting down service %s", a.opts.name)
			closer.Close()
			a.cancel()
			return nil
		}),
		micro.AfterStop(func() error {
			wg.Wait()
			return nil
		}),
		micro.WrapCall(ot.NewCallWrapper(tracer)),
		micro.WrapClient(ot.NewClientWrapper(tracer)),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer),
			wrap.Auth(a.opts.authFunc),
			wrap.Validator(),
			recovery.NewHandlerWrapper()),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(tracer)),
	}
	mo = append(mo, opts...)

	// Create service
	a.srv = micro.NewService(mo...)
	a.srv.Init(
		micro.Name(a.opts.name),
		micro.Version(a.opts.version),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*30),
		micro.Flags(
			&cli.BoolFlag{
				Name:  "migrate",
				Usage: "Execute db migrate",
			},
		),
		micro.Action(func(c *cli.Context) error {
			//迁移数据库
			if c.Bool("migrate") {
				a.opts.migrate()
				os.Exit(1)
			}
			return nil
		}),
	)
	if err = a.srv.Server().Init(server.Wait(&wg)); err != nil {
		panic(err)
	}
}

func (a *App) Service() micro.Service {
	return a.srv
}

func (a *App) Run() {
	// Register handler
	if err := pb.RegisterHealthHandler(a.srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := a.srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
