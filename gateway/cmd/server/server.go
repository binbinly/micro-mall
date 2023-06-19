package server

import (
	"gateway/internal/app"
	"gateway/pkg/trace"
	"log"
	"time"

	"gateway/pkg/config"
	"github.com/spf13/cobra"

	"gateway/internal/router"
	"gateway/internal/server"
	"gateway/pkg/registry"
	"gateway/pkg/registry/consul"
)

var (
	cfgDir   string
	env      string
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API Gateway Server",
		Example:      "gateway server -c configs/default.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&cfgDir, "config", "c", "configs", "config path")
	StartCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "Configure Runtime Environment")
}

func setup() {
	// init config
	c := config.New(cfgDir, config.WithEnv(env))
	if err := c.Load("app", app.Conf); err != nil {
		panic(err)
	}

	// init tracer
	_, err := trace.InitTracerProvider(
		trace.WithServiceName(app.Conf.App.Name),
		trace.WithEndpoint(app.Conf.Trace.Endpoint))
	if err != nil {
		log.Fatalf("Failed to init tracer: %v", err)
	}

}

func run() {
	mux := router.NewServeMux()

	http := server.NewHTTPServer(&app.Conf.HTTP)
	http.Handler = mux

	rs := consul.NewRegistry(registry.Addrs(app.Conf.Registry.Addr),
		consul.TCPCheck(time.Second*30))
	// init consul resolver
	consul.Init(rs)

	a := app.New(
		app.WithName(app.Conf.App.Name),
		app.WithVersion(app.Conf.App.Version),
		app.WithRegistry(rs),
		app.WithMux(mux),
		app.WithServer(http),
	)

	//start server
	if err := a.Run(); err != nil {
		log.Fatalf("App Run err: %v", err)
	}
}
