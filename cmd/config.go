package cmd

import (
	"encoding/json"
	"os"
	"todolist/applogger"
	"todolist/config"
	"todolist/constants"
)

var configCmd = newCommand(constants.ConfigCommandName, constants.ConfigCommandDescription, printConfigurations)

//TODO FIND A WAY TO GET ALL CONFIGS AND PRINT
func printConfigurations() {
	if err := config.Load(); err != nil {
		applogger.Errorf(constants.ErrorFailedToLoadConfig, "[setupDBConnection] [Load]", err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(config.GetServerConfig().Address()); err != nil {
		applogger.Errorf(constants.ErrorFailedToDisplayConfig, "[printConfigurations]")
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
