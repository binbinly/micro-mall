package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gateway/internal/app"
	"gateway/pkg/config"
)

var (
	file     string
	StartCmd = &cobra.Command{
		Use:     "config",
		Short:   "Get Application config info",
		Example: "gateway config -c configs/default.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&file, "config", "c", "configs/default.yaml", "config path")
}

func run() {
	config.Load(file, app.Conf, func(v *viper.Viper) {
		app.SetDefaultConf(v)
	})
}
