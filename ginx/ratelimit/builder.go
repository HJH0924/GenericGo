// Package ratelimit
/**
* @Project : GenericGo
* @File    : builder.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/11/8 11:32
**/

package ratelimit

import (
	"github.com/HJH0924/GenericGo/ratelimiter"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// Builder 用于为Gin框架创建一个限流中间件的结构体。
type Builder struct {
	limiter ratelimiter.RateLimiter // 限流器实例。
	keyFunc KeyFunc                 // 生成唯一键以进行限流的函数。
	logFunc LogFunc                 // 用于日志记录的函数。
}

// NewBuilder 使用给定的限流器创建一个新的Builder实例。
// keyFunc默认使用客户端的IP地址创建一个唯一键。
// logFunc默认使用log.Println。
func NewBuilder(limiter ratelimiter.RateLimiter) *Builder {
	return &Builder{
		limiter: limiter,
		keyFunc: func(ctx *gin.Context) string {
			var key strings.Builder
			key.WriteString("ip-limiter")
			key.WriteString(":")
			key.WriteString(ctx.ClientIP())
			return key.String()
		},
		logFunc: func(msg any, args ...any) {
			l := make([]any, 0, len(args)+1)
			l = append(l, msg)
			l = append(l, args...)
			log.Println(l...)
		},
	}
}

// SetKeyFunc 设置自定义的键函数以进行限流。
func (Self *Builder) SetKeyFunc(fn KeyFunc) *Builder {
	Self.keyFunc = fn
	return Self
}

// SetLogFunc 设置自定义的日志函数以进行限流。
func (Self *Builder) SetLogFunc(fn LogFunc) *Builder {
	Self.logFunc = fn
	return Self
}

func (Self *Builder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limited, err := Self.limit(ctx)
		if err != nil {
			Self.logFunc("Error in rate limiting", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if limited {
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		ctx.Next()
	}
}

// limit 根据提供的键检查请求是否被限流。
func (Self *Builder) limit(ctx *gin.Context) (bool, error) {
	return Self.limiter.IsLimit(ctx, Self.keyFunc(ctx))
}

// KeyFunc 是生成唯一键以进行限流的函数的类型别名。
type KeyFunc func(ctx *gin.Context) string

// LogFunc 是可以处理多个参数的日志函数的类型别名。
type LogFunc func(msg any, args ...any)
