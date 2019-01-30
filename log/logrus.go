package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/grpclog"
	"os"
)

const (
	Service   = "service"
	RequestID = "request_id"
)

func init() {
	// 将grpclog替换为logrus
	grpclog.SetLoggerV2(NewLogger().WithField("system", "system"))

	// 输出样式
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// 设置output
	logrus.SetOutput(os.Stdout)

	// 设置最低loglevel
	logrus.SetLevel(logrus.DebugLevel)
}

func With(ctx context.Context) *logrus.Entry {
	fields := logrus.Fields{}
	if service, ok := ctx.Value(Service).(string); ok {
		fields[Service] = service
	}
	if requestID, ok := ctx.Value(RequestID).(string); ok {
		fields[RequestID] = requestID
	}
	return logrus.WithFields(fields)
}
