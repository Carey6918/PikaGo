package main

import (
	"context"
	"github.com/Carey6918/PikaRPC/example/proto"
	"github.com/Carey6918/PikaRPC/log"
	"github.com/Carey6918/PikaRPC/server"
	"github.com/sirupsen/logrus"
)

type AddServerImpl struct{}

func main() {
	server.Init()
	add.RegisterAddServiceServer(server.GetGRPCServer(), &AddServerImpl{})
	if err := server.Run(); err != nil {
		logrus.Errorf("server run failed, err= %v", err)
	}
}

func (s *AddServerImpl) Add(ctx context.Context, req *add.AddRequest) (*add.AddResponse, error) {
	a := req.GetA()
	b := req.GetB()
	sum := a + b

	log.With(ctx).Infof("received a= %v, b= %v, sum= %v", a, b, sum)
	return &add.AddResponse{
		Sum: sum,
	}, nil
}
