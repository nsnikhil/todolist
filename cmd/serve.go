package cmd

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"todolist/app"
	"todolist/applogger"
	"todolist/config"
	"todolist/proto"
	"todolist/server"
	"todolist/util"
)

var serveCmd = newCommand(serveCommandName, serveCommandDescription, serve)


func serve() {
	s := server.NewServer(app.SetUpDependencies())
	gs := grpc.NewServer()
	proto.RegisterApiServer(gs, &s)

	go listenAndServe(gs, setUpListener())
	waitForShutdown(gs)
}

func setUpListener() net.Listener {
	listener, err := net.Listen(config.GetServerConfig().Protocol(), config.GetServerConfig().Address())
	if err != nil {
		util.LogError("[Server] [setUpListener]", err)
	}
	return listener
}

func listenAndServe(gs *grpc.Server, c net.Listener) {
	applogger.Infof("%s %s", "listening on ", config.GetServerConfig().Address())
	if err := gs.Serve(c); err != nil {
		applogger.Fatalf("failed to serve: %s", err)
	}
}

func waitForShutdown(gs *grpc.Server) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	gs.GracefulStop()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
