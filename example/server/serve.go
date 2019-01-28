package main

import (
	"context"
	"github.com/Carey6918/PikaRPC/example/proto"
	"github.com/Carey6918/PikaRPC/server"
	"google.golang.org/grpc/grpclog"
)

type AddServerImpl struct{}

func main() {
	server.Init()
	add.RegisterAddServiceServer(server.GetGRPCServer(), &AddServerImpl{})
	if err := server.Run(); err != nil {
		grpclog.Errorf("server run failed, err= %v", err)
	}
}

func (s *AddServerImpl) Add(ctx context.Context, req *add.AddRequest) (*add.AddResponse, error) {
	a := req.GetA()
	b := req.GetB()
	sum := a + b
	return &add.AddResponse{
		Sum: sum,
	}, nil
}
