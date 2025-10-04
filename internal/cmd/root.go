package cmd

import (
	"context"
	"fmt"

	"github.com/briheet/kumo/internal/cmdutils/config"
	"github.com/spf13/cobra"
)

func Execute(ctx context.Context) int {

	rootCmd := &cobra.Command{
		Use:   "service",
		Short: "Exchange injestion service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			// load config
			cfg := config.InitViperConfig(cmd)

			fmt.Println(cfg.Get("title"))

			// opentelemetry otel
			// tracing, metrics, runtime

			// Database pool, redis, natsconfig
			// Exchange config := derbit, bybit

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Main service logic goes here
			// Just a print here for now

			return nil

		},
	}

	rootCmd.PersistentFlags().String("config", "config.toml", "config file path")

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}
