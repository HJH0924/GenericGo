// Package logger
/**
* @Project : GenericGo
* @File    : zap_logger_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/11/5 09:57
**/

package logger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

func TestZapLogger_Debug(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		key  string
		val  any
	}{
		{
			name: "Debug",
			msg:  "输出用户名",
			key:  "name",
			val:  "Tvux",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core, obs := observer.New(zap.DebugLevel)
			zapLogger := zap.New(core)
			zl := NewZapLogger(zapLogger)

			zl.Debug(tt.msg, Field{Key: tt.key, Val: tt.val})
			entry := obs.AllUntimed()[0]
			assert.Equal(t, tt.msg, entry.Message)
			for _, field := range entry.Context {
				assert.Equal(t, tt.key, field.Key)
				assert.Equal(t, tt.val, field.String)
			}
		})
	}
}

func TestZapLogger_Info(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		key  string
		val  any
	}{
		{
			name: "Info",
			msg:  "输出用户名",
			key:  "name",
			val:  "Tvux",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core, obs := observer.New(zap.DebugLevel)
			zapLogger := zap.New(core)
			zl := NewZapLogger(zapLogger)

			zl.Info(tt.msg, Field{Key: tt.key, Val: tt.val})
			entry := obs.AllUntimed()[0]
			assert.Equal(t, tt.msg, entry.Message)
			for _, field := range entry.Context {
				assert.Equal(t, tt.key, field.Key)
				assert.Equal(t, tt.val, field.String)
			}
		})
	}
}

func TestZapLogger_Warn(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		key  string
		val  any
	}{
		{
			name: "Warn",
			msg:  "输出用户名",
			key:  "name",
			val:  "Tvux",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core, obs := observer.New(zap.DebugLevel)
			zapLogger := zap.New(core)
			zl := NewZapLogger(zapLogger)

			zl.Warn(tt.msg, Field{Key: tt.key, Val: tt.val})
			entry := obs.AllUntimed()[0]
			assert.Equal(t, tt.msg, entry.Message)
			for _, field := range entry.Context {
				assert.Equal(t, tt.key, field.Key)
				assert.Equal(t, tt.val, field.String)
			}
		})
	}
}

func TestZapLogger_Error(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		key  string
		val  any
	}{
		{
			name: "Warn",
			msg:  "输出用户名",
			key:  "name",
			val:  "Tvux",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core, obs := observer.New(zap.DebugLevel)
			zapLogger := zap.New(core)
			zl := NewZapLogger(zapLogger)

			zl.Error(tt.msg, Field{Key: tt.key, Val: tt.val})
			entry := obs.AllUntimed()[0]
			assert.Equal(t, tt.msg, entry.Message)
			for _, field := range entry.Context {
				assert.Equal(t, tt.key, field.Key)
				assert.Equal(t, tt.val, field.String)
			}
		})
	}
}
