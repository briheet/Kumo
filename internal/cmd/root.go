package cmd

import (
	"context"

	"github.com/briheet/kumo/internal/cmdutils/config"
	"github.com/spf13/cobra"
)

func Execute(ctx context.Context) int {

	rootCmd := &cobra.Command{
		Use:   "service",
		Short: "Exchange injestion service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			// load config
			cfg := config.InitViperConfig()

			// opentelemetry otel
			// tracing, metrics, runtime

			// Database pool, redis, natsconfig
			// Exchange config := derbit, bybit

			return nil
		},
	}

	return 0
}
