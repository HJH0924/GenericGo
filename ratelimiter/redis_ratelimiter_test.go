// Package ratelimiter
/**
* @Project : GenericGo
* @File    : redis_ratelimiter_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/11/8 14:20
**/

package ratelimiter

import (
	"context"
	"errors"
	"github.com/HJH0924/GenericGo/ratelimiter/redismocks"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestRedisRateLimiter_IsLimit(t *testing.T) {
	tests := []struct {
		name      string
		mock      func(ctrl *gomock.Controller, key string, window time.Duration, threshold int) redis.Cmdable
		window    time.Duration
		threshold int
		wantRes   bool
		wantErr   error
	}{
		{
			name: "没有触发限流",
			mock: func(ctrl *gomock.Controller, key string, window time.Duration, threshold int) redis.Cmdable {
				mockRedis := redismocks.NewMockCmdable(ctrl)
				mockRedisRes := redis.NewCmdResult(false, nil)
				mockRedis.EXPECT().Eval(gomock.Any(), slideWindowLua, []string{key}, window.Milliseconds(), threshold, time.Now().UnixMilli()).Return(mockRedisRes)
				return mockRedis
			},
			window:    time.Minute,
			threshold: 10,
			wantRes:   false,
		},
		{
			name: "触发限流",
			mock: func(ctrl *gomock.Controller, key string, window time.Duration, threshold int) redis.Cmdable {
				mockRedis := redismocks.NewMockCmdable(ctrl)
				mockRedisRes := redis.NewCmdResult(true, nil)
				mockRedis.EXPECT().Eval(gomock.Any(), slideWindowLua, []string{key}, window.Milliseconds(), threshold, time.Now().UnixMilli()).Return(mockRedisRes)
				return mockRedis
			},
			window:    time.Minute,
			threshold: 10,
			wantRes:   true,
		},
		{
			name: "系统错误",
			mock: func(ctrl *gomock.Controller, key string, window time.Duration, threshold int) redis.Cmdable {
				mockRedis := redismocks.NewMockCmdable(ctrl)
				mockRedisRes := redis.NewCmdResult(false, errors.New("系统错误"))
				mockRedis.EXPECT().Eval(gomock.Any(), slideWindowLua, []string{key}, window.Milliseconds(), threshold, time.Now().UnixMilli()).Return(mockRedisRes)
				return mockRedis
			},
			window:    time.Minute,
			threshold: 10,
			wantRes:   false,
			wantErr:   errors.New("系统错误"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			limiter := NewRedisRateLimiter(tt.mock(ctrl, "key", tt.window, tt.threshold), tt.window, tt.threshold)
			limited, err := limiter.IsLimit(context.Background(), "key")
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantRes, limited)
		})
	}
}
