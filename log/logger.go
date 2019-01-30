package log

import "github.com/sirupsen/logrus"

/**
	grpc_logrus的ReplaceGrpcLogger只继承了旧版本的grpc.Logger
	旧版本的Logger只提供了Fatal和Print两种级别的日志，缺失Debug、Warn、Error级别
	并且logrus没有实现 func V(l int)bool，无法继承grpc.LoggerV2
	因此这里封装了logrus，并实现了LoggerV2的所有方法，以继承新版grpclog
 **/

type Logger struct {
	e *logrus.Entry
}

func NewLogger() *Logger {
	return &Logger{
		e: logrus.NewEntry(logrus.StandardLogger()),
	}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	l.e = l.e.WithFields(logrus.Fields{key: value})
	return l
}

func (l *Logger) WithFields(fields logrus.Fields) *Logger {
	l.e = l.e.WithFields(fields)
	return l
}

func (l *Logger) Info(args ...interface{}) {
	l.e.Info(args)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.e.Infoln(args)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.e.Infof(format, args)
}

func (l *Logger) Warning(args ...interface{}) {
	l.e.Warning(args)
}

func (l *Logger) Warningln(args ...interface{}) {
	l.e.Warningln(args)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.e.Warningf(format, args)
}

func (l *Logger) Error(args ...interface{}) {
	l.e.Error(args)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.e.Errorln(args)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.e.Errorf(format, args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.e.Fatal(args)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.e.Fatalln(args)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.e.Fatalf(format, args)
}

func (l *Logger) V(i int) bool {
	return l.e.Logger.IsLevelEnabled(logrus.Level(i))
}
