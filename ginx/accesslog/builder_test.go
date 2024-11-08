// Package accesslog
/**
* @Project : GenericGo
* @File    : builder_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/11/8 09:37
**/

package accesslog

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuilder_Build(t *testing.T) {
	tests := []struct {
		name              string
		middlewareBuilder func(fn LogFunc) gin.HandlerFunc
		wantAl            AccessLog
	}{
		{
			name: "打印全部请求体跟响应体",
			middlewareBuilder: func(fn LogFunc) gin.HandlerFunc {
				return NewBuilder(fn).AllowReqBody(true).AllowRespBody(true).MaxLength(1024).Build()
			},
			wantAl: AccessLog{
				Method:     "POST",
				URL:        "/accessLog",
				ReqBody:    `{"name":"Tom"}`,
				RespBody:   `{"msg":"hello Tom"}`,
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "打印全部请求体，不打印响应体",
			middlewareBuilder: func(fn LogFunc) gin.HandlerFunc {
				return NewBuilder(fn).AllowReqBody(true).AllowRespBody(false).MaxLength(1024).Build()
			},
			wantAl: AccessLog{
				Method:  "POST",
				URL:     "/accessLog",
				ReqBody: `{"name":"Tom"}`,
			},
		},
		{
			name: "不打印请求体，打印全部响应体",
			middlewareBuilder: func(fn LogFunc) gin.HandlerFunc {
				return NewBuilder(fn).AllowReqBody(false).AllowRespBody(true).MaxLength(1024).Build()
			},
			wantAl: AccessLog{
				Method:     "POST",
				URL:        "/accessLog",
				RespBody:   `{"msg":"hello Tom"}`,
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "不打印请求体和响应体",
			middlewareBuilder: func(fn LogFunc) gin.HandlerFunc {
				return NewBuilder(fn).AllowReqBody(false).AllowRespBody(false).MaxLength(1024).Build()
			},
			wantAl: AccessLog{
				Method: "POST",
				URL:    "/accessLog",
			},
		},
		{
			name: "打印请求体和响应体 - 限制长度",
			middlewareBuilder: func(fn LogFunc) gin.HandlerFunc {
				return NewBuilder(fn).AllowReqBody(true).AllowRespBody(true).MaxLength(5).Build()
			},
			wantAl: AccessLog{
				Method:     "POST",
				URL:        "/acce",
				ReqBody:    `{"name":"Tom"}`,
				RespBody:   `{"msg`,
				StatusCode: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/accessLog", bytes.NewBuffer([]byte(`{"name":"Tom"}`)))
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			ctx, server := gin.CreateTestContext(resp)
			ctx.Request = req

			resAl := new(AccessLog)
			server.Use(tt.middlewareBuilder(func(ctx context.Context, al *AccessLog) {
				copyTo(al, resAl)
				fmt.Printf("%+v\n", al)
			}))
			server.POST("/accessLog", func(ctx *gin.Context) {
				type Req struct {
					Name string `json:"name"`
				}
				req := new(Req)
				err := ctx.BindJSON(req)
				assert.NoError(t, err)
				msg := fmt.Sprintf("hello %s", req.Name)
				ctx.JSON(http.StatusOK, map[string]any{
					"msg": msg,
				})
			})
			server.ServeHTTP(resp, req)

			assert.Equal(t, tt.wantAl.Method, resAl.Method)
			assert.Equal(t, tt.wantAl.URL, resAl.URL)
			assert.Equal(t, tt.wantAl.ReqBody, resAl.ReqBody)
			assert.Equal(t, tt.wantAl.RespBody, resAl.RespBody)
			assert.Equal(t, tt.wantAl.StatusCode, resAl.StatusCode)
		})
	}
}

func copyTo(src, dst *AccessLog) {
	dst.Method = src.Method
	dst.URL = src.URL
	dst.Duration = src.Duration
	dst.ReqBody = src.ReqBody
	dst.RespBody = src.RespBody
	dst.StatusCode = src.StatusCode
}
