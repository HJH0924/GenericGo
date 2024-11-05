// Package logger
/**
* @Project : GenericGo
* @File    : zap_logger.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/11/5 09:26
**/

package logger

import "go.uber.org/zap"

type ZapLogger struct {
	zl *zap.Logger
}

func NewZapLogger(zl *zap.Logger) *ZapLogger {
	return &ZapLogger{
		zl: zl,
	}
}

func (Self *ZapLogger) Debug(msg string, args ...Field) {
	Self.zl.Debug(msg, Self.toZapFields(args)...)
}

func (Self *ZapLogger) Info(msg string, args ...Field) {
	Self.zl.Info(msg, Self.toZapFields(args)...)
}

func (Self *ZapLogger) Warn(msg string, args ...Field) {
	Self.zl.Warn(msg, Self.toZapFields(args)...)
}

func (Self *ZapLogger) Error(msg string, args ...Field) {
	Self.zl.Error(msg, Self.toZapFields(args)...)
}

func (Self *ZapLogger) toZapFields(args []Field) []zap.Field {
	res := make([]zap.Field, 0, len(args))
	for _, arg := range args {
		res = append(res, zap.Any(arg.Key, arg.Val))
	}
	return res
}
