package config

import (
	"fmt"
	"gateway/internal/app"
	"github.com/binbinly/pkg/config"

	"github.com/spf13/cobra"
)

var (
	cfgDir   string
	env      string
	StartCmd = &cobra.Command{
		Use:     "config",
		Short:   "Get Application config info",
		Example: "mall config -c default.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&cfgDir, "config", "c", "configs", "config path")
	StartCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "Configure Runtime Environment")
}

func run() {
	// init config
	c := config.New(cfgDir, config.WithEnv(env))
	if err := c.Load("app", app.Conf); err != nil {
		panic(err)
	}
	fmt.Printf("config:%+v\n", app.Conf)
}