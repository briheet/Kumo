package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func GUICmd(context context.Context) *cobra.Command {

	guiCmd := &cobra.Command{
		Use:   "gui",
		Short: "GUI for kumo",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("HI")
			return nil
		},
	}

	return guiCmd

}
