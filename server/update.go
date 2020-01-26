package server

import (
	"context"
	"todolist/proto"
	"todolist/util"
)

func (s Server) Update(ctx context.Context, in *proto.UpdateRequest) (*proto.ApiResponse, error) {
	task, err := s.TaskFactory.CreateWithID(in.Id, in.Title, in.Description, in.Status, in.Tags...)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Update] [CreateWithID]", err)
	}

	if err := s.Service.Update(task); err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Update]", err)
	}

	return &proto.ApiResponse{}, nil
}
