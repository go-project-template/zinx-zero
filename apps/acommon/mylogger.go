package acommon

import (
	"context"

	"github.com/aceld/zinx/ziface"
	"github.com/zeromicro/go-zero/core/logx"
)

// 重置 zinx 日志为 go-zero -> logx
// 可以通过自身业务的日志方式，来重置zinx内部引擎的日志打印方式
type MyLogger struct {
	ziface.ILogger
}

// 没有context的日志接口
func (l *MyLogger) DebugF(format string, v ...interface{}) {
	logx.WithCallerSkip(1).Debugf(format, v...)
}
func (l *MyLogger) InfoF(format string, v ...interface{}) {
	logx.WithCallerSkip(1).Infof(format, v...)
}
func (l *MyLogger) ErrorF(format string, v ...interface{}) {
	logx.WithCallerSkip(1).Errorf(format, v...)
}

// 携带context的日志接口
func (l *MyLogger) DebugFX(ctx context.Context, format string, v ...interface{}) {
	logx.WithContext(ctx).WithCallerSkip(1).Debugf(format, v...)
}
func (l *MyLogger) InfoFX(ctx context.Context, format string, v ...interface{}) {
	logx.WithContext(ctx).WithCallerSkip(1).Infof(format, v...)
}
func (l *MyLogger) ErrorFX(ctx context.Context, format string, v ...interface{}) {
	logx.WithContext(ctx).WithCallerSkip(1).Errorf(format, v...)
}
