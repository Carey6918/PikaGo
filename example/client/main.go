package main

import (
	"context"
	"github.com/Carey6918/PikaRPC/client"
	"github.com/Carey6918/PikaRPC/config"
	"github.com/Carey6918/PikaRPC/example/proto"
	"github.com/Carey6918/PikaRPC/log"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	ServiceName = "carey.is.genius"
)

func main() {
	client.Init(client.WithWatchInterval(10 * time.Second))
	conn, err := client.GetConn(ServiceName)
	defer client.Close(ServiceName)
	if err != nil {
		logrus.Fatalf("get conn failed, err= %v", err)
	}

	cli := add.NewAddServiceClient(conn)
	req := &add.AddRequest{
		A: 1,
		B: 1,
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, log.RequestID, "test")
	ctx = context.WithValue(ctx, log.Service, config.ServiceConf.ServiceName)
	resp, err := cli.Add(ctx, req)
	if err != nil {
		log.Logger(ctx).Fatalf("add failed, err= %v", err)
	}
	log.Logger(ctx).Infof("resp= %v", resp)
}
