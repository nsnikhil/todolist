package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/health/grpc_health_v1"
	"todolist/proto"
	"todolist/util"
)

var (
	notServing = &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
	}

	serving = &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}
)

func (hs *HealthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {

	res, err := hs.Ping(ctx, &proto.PingRequest{})
	if err != nil {
		return notServing, util.LogAndGetError("[HealthServer] [Check] [Ping]", err)
	}

	if res == nil || res.Data == nil || string(res.Data.Value) != "pong" {
		return notServing, util.LogAndGetError("[HealthServer] [Check] [res]", fmt.Errorf("invalid response :%v", res))
	}

	util.DebugLog("[HealthServer] [Check] [Success]")
	return serving, nil

}
