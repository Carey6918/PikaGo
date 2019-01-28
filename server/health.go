package server

import (
	"context"
	"github.com/Carey6918/PikaRPC/client"
	"google.golang.org/grpc/grpclog"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"time"
)

const healthCheckInterval = 30 * time.Second

// gRPC健康检查，实现了grpc_health_v1.HealthServer接口
type HealthServerImpl struct{}

func (s *HealthServerImpl) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	client.Init(client.WithWatchInterval(healthCheckInterval))
	_, err := client.GetConn(req.GetService())
	defer client.Close(req.GetService())
	if err != nil {
		grpclog.Errorf("health check to %v failed, err= %v", req.GetService(), err)
		return &health.HealthCheckResponse{
			Status: health.HealthCheckResponse_NOT_SERVING,
		}, nil
	}
	grpclog.Infof("health check to %v success", req.GetService())
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthServerImpl) Watch(req *health.HealthCheckRequest, server health.Health_WatchServer) error {
	return nil
}
