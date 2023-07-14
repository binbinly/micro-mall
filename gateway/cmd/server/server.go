package server

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gateway/internal/app"
	"gateway/internal/router"
	"gateway/internal/server"
	"gateway/pkg/config"
	"gateway/pkg/registry"
	"gateway/pkg/registry/consul"
	"gateway/pkg/trace"
)

var (
	file     string
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API Gateway Server",
		Example:      "gateway server -c configs/default.yaml",
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
	StartCmd.PersistentFlags().StringVarP(&file, "config", "c", "configs/default.yaml", "config path")
}

func setup() {
	// load config
	config.Load(file, app.Conf, func(v *viper.Viper) {
		app.SetDefaultConf(v)
	})

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
