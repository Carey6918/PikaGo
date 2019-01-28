package server

import (
	"fmt"
	"github.com/Carey6918/PikaRPC/client"
	"github.com/Carey6918/PikaRPC/helper"
	"github.com/Carey6918/PikaRPC/tracing"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/resolver"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	gServer  *grpc.Server
	option   *Option
	listener net.Listener
}

var GServer *Server // 全局服务
var tracer opentracing.Tracer

func Init() {
	initConfig()
	resolver.Register(client.NewBuilder("test")) // consul lb

	// 通过consul注册服务
	if err := newRegisterContest().Register(); err != nil {
		fmt.Errorf("consul register failed, err= %v", err)
	}

	// 初始化zepkin跟踪器
	var err error
	tracer, err = tracing.NewZipkinTracer(ServiceConf.ServiceName)
	if err != nil {
		fmt.Errorf("init tracing failed, err= %v", err)
	}

	newServer(WithGRPCOpts(grpc.ConnectionTimeout(1*time.Second),
		grpc_middleware.WithUnaryServerChain(otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()))))

	grpc_health_v1.RegisterHealthServer(GetGRPCServer(), &HealthServerImpl{})
}

func newServer(opts ...Options) {
	var server Server
	server.option = new(Option)
	for _, opt := range opts {
		opt(server.option)
	}
	// 初始化gRPC服务
	server.gServer = grpc.NewServer(server.option.gOpts...)
	GServer = &server
}

func Run() error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- GServer.serve()
	}()
	return waitSignal(errCh)
}

func waitSignal(errCh chan error) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

	for {
		select {
		case sig := <-signals:
			switch sig {
			// exit forcely
			case syscall.SIGTERM: // 结束程序(可以被捕获、阻塞或忽略)
				fmt.Printf("stop run, signals= %v", sig.String())
				return nil
			case syscall.SIGHUP, syscall.SIGINT: // 终端连接断开/用户发送(ctrl+c)结束
				GServer.stop()
				fmt.Printf("stop run, signals= %v", sig.String())
				return nil
			}
		case err := <-errCh:
			return err
		}
	}
	return <-errCh
}

func (s *Server) serve() error {
	if err := s.listen(); err != nil {
		return err
	}

	// 注册gRPC服务
	reflection.Register(s.gServer)
	if err := s.gServer.Serve(s.listener); err != nil {
		return err
	}
	return nil
}

func (s *Server) stop() error {
	return s.listener.Close()
}

func (s *Server) listen() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", helper.GetLocalIP(), ServiceConf.ServicePort))
	if err != nil {
		fmt.Errorf("listen tcp failed, err= %v", err)
		return err
	}
	s.listener = listener
	return nil
}

func GetGRPCServer() *grpc.Server {
	return GServer.gServer
}
