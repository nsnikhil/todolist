package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"todolist/applogger"
	"todolist/constants"
)

type ServerInterface interface {
	Serve(address string)
}

type Server struct {
	apiServer *http.Server
}

func NewServer(router http.Handler) *Server {
	server := &Server{
		apiServer: &http.Server{
			Handler: router,
		},
	}
	applogger.Infof(constants.NewServerCreated, "[Server] [NewServer]", server)
	return server
}

func (srv Server) Serve(address string) {
	applogger.Infof(constants.ServerStartedListeningOn, "[Server] [Serve]", address)
	srv.apiServer.Addr = address
	go listenServer(srv.apiServer)
	waitForShutdown(srv.apiServer)
}

//TODO FIX THE ERROR NIL POINTER CHECK
func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		applogger.Fatalf(err.Error())
	}
}

func waitForShutdown(apiServer *http.Server) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh
	applogger.Infof(constants.ServerShuttingDown, sig)
	if err := apiServer.Shutdown(context.Background()); err != nil {
		applogger.Errorf(constants.ErrorFailedToShutdownServer, "[Server] [waitForShutdown]", err)
	}
	applogger.Infof(constants.ServerShutDownComplete)
}
