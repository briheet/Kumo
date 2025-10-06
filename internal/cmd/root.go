package cmd

import (
	"context"
	"fmt"

	"github.com/briheet/kumo/internal/cmdutils/config"
	"github.com/spf13/cobra"
)

func Execute(ctx context.Context) int {

	rootCmd := &cobra.Command{
		Use:   "kumo",
		Short: "Exchange injestion service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			// load config
			cfg := config.InitViperConfig(cmd)

			fmt.Println(cfg.Get("title"))

			// Init of things at root level

			// opentelemetry otel
			// tracing, metrics, runtime

			// Database pool, redis, natsconfig
			// Exchange config := derbit, bybit

			return nil
		},
	}

	rootCmd.PersistentFlags().String("config", "config.toml", "config file path")

	rootCmd.AddCommand(ServiceCmd(ctx))
	rootCmd.AddCommand(GUICmd(ctx))

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}
