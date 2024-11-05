// Package logger
/**
* @Project : GenericGo
* @File    : types.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/11/5 09:24
**/

package logger

type Logger interface {
	Debug(msg string, args ...Field)
	Info(msg string, args ...Field)
	Warn(msg string, args ...Field)
	Error(msg string, args ...Field)
}

type Field struct {
	Key string
	Val any
}
