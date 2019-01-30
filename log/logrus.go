package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	Service   = "service"
	RequestID = "request_id"
)

func init() {
	// 输出样式
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// 设置output
	logrus.SetOutput(os.Stdout)

	// 设置最低loglevel
	logrus.SetLevel(logrus.DebugLevel)
}

// todo 补充ctx的to_cluster & from_cluster
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
