package server

import (
	"todolist/app"
)

type Server struct {
	app.Dependencies
}

func NewServer(dependencies app.Dependencies) Server {
	return Server{Dependencies: dependencies}
}
