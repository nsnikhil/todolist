package cmd

import (
	"encoding/json"
	"os"
	"todolist/config"
	"todolist/util"
)

var configCmd = newCommand(configCommandName, configCommandDescription, printConfigurations)

//TODO FIND A WAY TO GET ALL CONFIGS AND PRINT
func printConfigurations() {
	if err := config.Load(); err != nil {
		util.LogError("[printConfigurations] [config.Load]", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(config.GetServerConfig().Address()); err != nil {
		util.LogError("[printConfigurations] [json.NewEncoder]", err)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
