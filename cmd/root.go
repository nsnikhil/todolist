package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"todolist/constants"
)

var rootCmd = &cobra.Command{
	Use:     constants.AppName,
	Short:   constants.AppDescription,
	Version: constants.AppVersion,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
