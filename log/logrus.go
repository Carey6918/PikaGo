package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	FromService = "from_service"
	ToService   = "to_service"
	RequestID   = "request_id"
)

func init() {
	// 输出样式
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// 设置output
	logrus.SetOutput(os.Stdout)

	// 设置最低loglevel
	logrus.SetLevel(logrus.DebugLevel)
}

func With(ctx context.Context) *logrus.Entry {
	fields := logrus.Fields{}
	if fromService, ok := ctx.Value(FromService).(string); ok {
		fields[FromService] = fromService
	}
	if toService, ok := ctx.Value(ToService).(string); ok {
		fields[ToService] = toService
	}
	if requestID, ok := ctx.Value(RequestID).(string); ok {
		fields[RequestID] = requestID
	}
	return logrus.WithFields(fields)
}
