package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"todolist/proto"
	"todolist/util"
)

func (s Server) Remove(ctx context.Context, in *proto.RemoveRequest) (*proto.ApiResponse, error) {
	count, err := s.Service.Remove(in.Id, in.Ids...)
	if err != nil {
		return &proto.ApiResponse{}, util.LogAndGetError("[Server] [Remove]", err)
	}

	util.DebugLog("[Server] [Remove] [Success]")
	return &proto.ApiResponse{Data: &any.Any{Value: toByteSlice(count)}}, nil
}
