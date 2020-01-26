package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"todolist/proto"
)

func (s Server) Ping(ctx context.Context, in *proto.PingRequest) (*proto.ApiResponse, error) {
	return &proto.ApiResponse{Data: &any.Any{TypeUrl: "string", Value: []byte("pong")}}, nil
}
