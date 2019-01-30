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
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the debug severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

func Logger(ctx context.Context) *logrus.Entry {
	fields := logrus.Fields{}
	if service, ok := ctx.Value(Service).(string); ok {
		fields[Service] = service
	}
	if requestID, ok := ctx.Value(RequestID).(string); ok {
		fields[RequestID] = requestID
	}
	return logrus.WithFields(fields)
}
