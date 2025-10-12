package cmd

import (
	"context"

	"github.com/briheet/kumo/internal/gui"
	"github.com/spf13/cobra"
)

func GUICmd(ctx context.Context) *cobra.Command {

	guiCmd := &cobra.Command{
		Use:   "gui",
		Short: "GUI for kumo",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			err := gui.Start(ctx)
			return err
		},
	}

	return guiCmd
}
