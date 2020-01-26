package server

import (
	"context"
	"todolist/proto"
	"todolist/util"
)

func (s Server) Add(ctx context.Context, in *proto.AddRequest) (*proto.ApiResponse, error) {
	task, err := s.TaskFactory.Create(in.Title, in.Description, in.Status, in.Tags...)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Add] [Create]", err)
	}

	if err = s.Service.Add(task); err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Add]", err)
	}

	util.DebugLog("[Server] [Add] [Success]")
	return &proto.ApiResponse{}, nil
}
