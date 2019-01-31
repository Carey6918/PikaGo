package client

import (
	"fmt"
	"github.com/Carey6918/PikaRPC/config"
	"github.com/Carey6918/PikaRPC/log"
	"github.com/Carey6918/PikaRPC/tracing"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"
)

func init() {
	resolver.Register(NewBuilder("test")) // consul lb
}

type Client struct {
	sync.RWMutex
	connPool map[string]*grpc.ClientConn
	options  *Option
}

var GClient *Client

func Init(opts ...Options) {
	config.Init()

	var client Client
	client.options = &Option{
		watchInterval: 20 * time.Second,
	}
	for _, opt := range opts {
		opt(client.options)
	}
	client.connPool = make(map[string]*grpc.ClientConn)
	GClient = &client
}

func GetConn(serviceName string) (*grpc.ClientConn, error) {
	GClient.RLock()
	if cli, ok := GClient.connPool[serviceName]; ok {
		GClient.RUnlock()
		return cli, nil
	}
	GClient.RUnlock()

	// 通过consul服务发现
	service, err := discovery(serviceName)
	if err != nil {
		return nil, err
	}

	tracer, err := tracing.NewZipkinTracer(config.ServiceConf.ServiceName)
	if err != nil {
		return nil, err
	}

	entry := logrus.NewEntry(logrus.StandardLogger())
	// 将grpclog替换为logrus
	// grpc_logrus.ReplaceGrpcLogger(entry) （不用这个，原因写在log/logger.go里）
	grpclog.SetLoggerV2(log.NewLogger(entry).WithField("system", "system"))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())),
		grpc.WithUnaryInterceptor(grpc_logrus.UnaryClientInterceptor(entry)),
		grpc.WithUnaryInterceptor(FillContextInterceptor(serviceName)))
	if err != nil {
		return nil, err
	}

	GClient.Lock()
	defer GClient.Unlock()
	GClient.connPool[serviceName] = conn
	return conn, nil
}

func Close(service string) error {
	if conn, ok := GClient.connPool[service]; ok {
		return conn.Close()
	}
	return nil
}
