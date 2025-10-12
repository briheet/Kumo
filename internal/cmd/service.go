package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func ServiceCmd(context context.Context) *cobra.Command {

	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "Exchange injestion service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hi")

			return nil
		},
	}

	return serviceCmd
}
