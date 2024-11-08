// Package ratelimit
/**
* @Project : GenericGo
* @File    : builder_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/11/8 12:13
**/

package ratelimit

import (
	"errors"
	"github.com/HJH0924/GenericGo/ratelimiter"
	"github.com/HJH0924/GenericGo/ratelimiter/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBuilder_Build(t *testing.T) {
	tests := []struct {
		name     string
		mock     func(ctrl *gomock.Controller, key string) ratelimiter.RateLimiter
		wantCode int
	}{
		{
			name: "触发限流",
			mock: func(ctrl *gomock.Controller, key string) ratelimiter.RateLimiter {
				limiter := mocks.NewMockRateLimiter(ctrl)
				limiter.EXPECT().IsLimit(gomock.Any(), key).Return(true, nil)
				return limiter
			},
			wantCode: http.StatusTooManyRequests,
		},
		{
			name: "没有触发限流",
			mock: func(ctrl *gomock.Controller, key string) ratelimiter.RateLimiter {
				limiter := mocks.NewMockRateLimiter(ctrl)
				limiter.EXPECT().IsLimit(gomock.Any(), key).Return(false, nil)
				return limiter
			},
			wantCode: http.StatusOK,
		},
		{
			name: "系统错误",
			mock: func(ctrl *gomock.Controller, key string) ratelimiter.RateLimiter {
				limiter := mocks.NewMockRateLimiter(ctrl)
				limiter.EXPECT().IsLimit(gomock.Any(), key).Return(false, errors.New("系统错误"))
				return limiter
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			req, err := http.NewRequest(http.MethodGet, "/rateLimit", nil)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			ctx, server := gin.CreateTestContext(resp)
			ctx.Request = req

			builder := NewBuilder(tt.mock(ctrl, keyFunc(ctx))).SetKeyFunc(keyFunc).SetLogFunc(logFunc)
			server.Use(builder.Build())
			server.GET("/rateLimit", func(ctx *gin.Context) {
				ctx.Status(http.StatusOK)
			})
			server.ServeHTTP(resp, req)

			assert.Equal(t, tt.wantCode, resp.Code)
		})
	}
}

func keyFunc(ctx *gin.Context) string {
	var key strings.Builder
	key.WriteString("ip-limiter")
	key.WriteString(":")
	key.WriteString(ctx.ClientIP())
	return key.String()
}

func logFunc(msg any, args ...any) {
	l := make([]any, 0, len(args)+1)
	l = append(l, msg)
	l = append(l, args...)
	log.Println(l...)
}
