package server

import (
	"todolist/app"
	"todolist/proto"
)

type Server struct {
	app.Dependencies
}

func NewServer(dependencies app.Dependencies) Server {
	return Server{Dependencies: dependencies}
}

type HealthServer struct {
	proto.ApiServer
}

func NewHealthServer(server proto.ApiServer) *HealthServer {
	return &HealthServer{ApiServer: server}
}
