package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"todolist/proto"
	"todolist/util"
)

func (s Server) Add(ctx context.Context, in *proto.AddRequest) (*proto.ApiResponse, error) {
	task, err := s.TaskFactory.Create(in.Title, in.Description, in.Status, in.Tags...)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Add] [Create]", err)
	}

	id, err := s.Service.Add(task)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Add]", err)
	}

	util.DebugLog("[Server] [Add] [Success]")
	return &proto.ApiResponse{Data: &any.Any{TypeUrl: "string", Value: []byte(id)}}, nil
}
