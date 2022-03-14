package logger

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// Log 获取log对象
func Log(name ...string) *glog.Logger {
	if len(name) > 0 && name[0] != "" {
		return g.Log(name[0]).Skip(1).Line()
	}
	return g.Log().Skip(1).Line()
}

func Print(ctx context.Context, v ...interface{}) {
	Log("access").Print(ctx, v)
}

func Printf(ctx context.Context, format string, v ...interface{}) {
	Log("access").Printf(ctx, format, v)
}

func Info(ctx context.Context, v ...interface{}) {
	Log("access").Info(ctx, v)
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	Log("access").Infof(ctx, format, v)
}

func Debug(ctx context.Context, v ...interface{}) {
	Log("access").Debug(ctx, v)
}

func Debugf(ctx context.Context, format string, v ...interface{}) {
	Log("access").Debugf(ctx, format, v)
}

func Notice(ctx context.Context, v ...interface{}) {
	Log("access").Notice(ctx, v)
}

func Noticef(ctx context.Context, format string, v ...interface{}) {
	Log("access").Noticef(ctx, format, v)
}

func Warning(ctx context.Context, v ...interface{}) {
	Log("error").Warning(ctx, v)
}

func Warningf(ctx context.Context, format string, v ...interface{}) {
	Log("error").Warningf(ctx, format, v)
}

func Error(ctx context.Context, v ...interface{}) {
	Log("error").Error(ctx, v)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	Log("error").Errorf(ctx, format, v)
}

func Fatal(ctx context.Context, v ...interface{}) {
	Log("error").Fatal(ctx, v)
}

func Fatalf(ctx context.Context, format string, v ...interface{}) {
	Log("error").Fatalf(ctx, format, v)
}

func Critical(ctx context.Context, v ...interface{}) {
	Log("error").Critical(ctx, v)
}

func Criticalf(ctx context.Context, format string, v ...interface{}) {
	Log("error").Criticalf(ctx, format, v)
}

func Panic(ctx context.Context, v ...interface{}) {
	Log("error").Panic(ctx, v)
}

func Panicf(ctx context.Context, format string, v ...interface{}) {
	Log("error").Panicf(ctx, format, v)
}
