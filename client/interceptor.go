package client

import (
	"context"
	"github.com/Carey6918/PikaRPC/config"
	"github.com/Carey6918/PikaRPC/log"
	"google.golang.org/grpc"
)

func FillContextInterceptor(toService string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.With(ctx).Infof("FillContextInterceptor.....")
		ctx = context.WithValue(ctx, log.FromService, config.ServiceConf.ServiceName)
		ctx = context.WithValue(ctx, log.ToService, toService)
		log.With(ctx).Infof("FillContextInterceptor.....")
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}
