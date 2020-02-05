package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"todolist/proto"
	"todolist/util"
)

func (s Server) Update(ctx context.Context, in *proto.UpdateRequest) (*proto.ApiResponse, error) {
	task, err := s.TaskFactory.CreateWithID(in.Id, in.Title, in.Description, in.Status, in.Tags...)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Update] [CreateWithID]", err)
	}

	count, err := s.Service.Update(task)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Update]", err)
	}

	util.DebugLog("[Server] [Update] [Success]")
	return &proto.ApiResponse{Data: &any.Any{Value: toByteSlice(count)}}, nil
}
