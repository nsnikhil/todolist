package cmd

import (
	"todolist/app"
	"todolist/config"
	"todolist/constants"
	"todolist/handler"
)

var serveCmd = newCommand(constants.ServeCommandName, constants.ServeCommandDescription, bootStrap)

func bootStrap() {
	dependencies := app.SetUpDependencies()
	router := handler.NewRouter(dependencies)
	server := app.NewServer(router)
	server.Serve(config.GetServerConfig().Address())
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
