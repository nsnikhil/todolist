package cmd

import (
	"github.com/spf13/cobra"
)

func newCommand(name string, description string, run func()) *cobra.Command {
	return &cobra.Command{
		Use:   name,
		Short: description,
		Run:   func(cmd *cobra.Command, args []string) {
			run()
		},
	}
}
