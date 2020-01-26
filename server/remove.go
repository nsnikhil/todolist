package server

import (
	"context"
	"todolist/proto"
	"todolist/util"
)

func (s Server) Remove(ctx context.Context, in *proto.RemoveRequest) (*proto.ApiResponse, error) {
	err := s.Service.Remove(in.Id, in.Ids...)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Remove]", err)
	}

	util.DebugLog("[Server] [Remove] [Success]")
	return &proto.ApiResponse{}, nil
}
