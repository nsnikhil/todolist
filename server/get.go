package server

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"todolist/proto"
	"todolist/util"
)

func (s Server) Get(ctx context.Context, req *proto.GetRequest) (*proto.ApiResponse, error) {
	tasks, err := s.Service.GetTasks(req.Id...)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Get] [GetTasks]", err)
	}

	b, err := json.Marshal(tasks)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Get] [Marshal]", err)
	}

	util.DebugLog("[Server] [Get] [Success]")
	return &proto.ApiResponse{Data: &any.Any{Value: b}}, nil
}
