// Package accesslog
/**
* @Project : GenericGo
* @File    : builder.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/11/8 08:45
**/

package accesslog

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/atomic"
	"io"
	"time"
)

// Builder 用于构建并配置 Gin 的日志中间件。
type Builder struct {
	allowReqBody  *atomic.Bool  // 是否允许记录请求体
	allowRespBody *atomic.Bool  // 是否允许记录响应体
	maxLength     *atomic.Int64 // 请求体和响应体的最大记录长度
	logFunc       LogFunc       // 日志记录函数
}

// NewBuilder 创建一个新的 Builder 实例，并初始化配置。
func NewBuilder(fn LogFunc) *Builder {
	return &Builder{
		allowReqBody:  atomic.NewBool(false),
		allowRespBody: atomic.NewBool(false),
		maxLength:     atomic.NewInt64(1024),
		logFunc:       fn,
	}
}

// AllowReqBody 设置是否允许记录请求体。
func (Self *Builder) AllowReqBody(ok bool) *Builder {
	Self.allowReqBody.Store(ok)
	return Self
}

// AllowRespBody 设置是否允许记录响应体。
func (Self *Builder) AllowRespBody(ok bool) *Builder {
	Self.allowRespBody.Store(ok)
	return Self
}

// MaxLength 设置请求体和响应体的最大记录长度
func (Self *Builder) MaxLength(l int64) *Builder {
	Self.maxLength.Store(l)
	return Self
}

// Build 构建并返回一个 Gin 中间件，该中间件在请求处理前后记录日志，并截取请求和响应体。
func (Self *Builder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		url := ctx.Request.URL.String()
		urlLen := int64(len(url))
		maxLength := Self.maxLength.Load()
		allowReqBody := Self.allowReqBody.Load()
		allowRespBody := Self.allowRespBody.Load()

		if urlLen >= maxLength {
			url = url[:maxLength]
		}

		al := &AccessLog{
			Method: ctx.Request.Method,
			URL:    url,
		}

		if allowReqBody && ctx.Request.Body != nil {
			body, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
			al.ReqBody = string(body)
		}

		if allowRespBody {
			ctx.Writer = respWriter{
				al:             al,
				maxLength:      maxLength,
				ResponseWriter: ctx.Writer,
			}
		}

		defer func() {
			al.Duration = time.Since(start).String()
			Self.logFunc(ctx, al)
		}()

		// 执行业务逻辑
		ctx.Next()
	}
}

// respWriter 是一个自定义的 ResponseWriter，用于截取响应并记录响应体。
type respWriter struct {
	al        *AccessLog
	maxLength int64
	gin.ResponseWriter
}

// WriteHeader 重写 WriteHeader 方法，以便记录响应状态码。
func (Self respWriter) WriteHeader(statusCode int) {
	Self.al.StatusCode = statusCode
	Self.ResponseWriter.WriteHeader(statusCode)
}

// Write 重写 Write 方法，以便记录响应体。
func (Self respWriter) Write(body []byte) (int, error) {
	curLen := int64(len(body))
	if curLen >= Self.maxLength {
		body = body[:Self.maxLength]
	}
	Self.al.RespBody = string(body)
	return Self.ResponseWriter.Write(body)
}

// AccessLog 记录访问日志的结构体。
type AccessLog struct {
	Method     string
	URL        string
	Duration   string
	ReqBody    string
	RespBody   string
	StatusCode int
}

// LogFunc 定义日志记录函数的签名。
type LogFunc func(ctx context.Context, al *AccessLog)
