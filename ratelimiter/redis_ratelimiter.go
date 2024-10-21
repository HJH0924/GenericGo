// Package ratelimiter
/**
* @Project : GenericGo
* @File    : redis_ratelimiter.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/19 12:12
**/

package ratelimiter

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisRateLimiter 基于 Redis 的限流器实现。
type RedisRateLimiter struct {
	redisClient redis.Cmdable // 用于执行 Redis 命令的客户端
	window      time.Duration // 限流的时间窗口
	threshold   int           // 在时间窗口内允许的最大请求次数
}

// NewRedisRateLimiter 函数创建并返回一个新的 RedisRateLimiter 实例。
func NewRedisRateLimiter(redisClient redis.Cmdable, window time.Duration, threshold int) *RedisRateLimiter {
	return &RedisRateLimiter{
		redisClient: redisClient,
		window:      window,
		threshold:   threshold,
	}
}

// slideWindowLua 是嵌入的 Lua 脚本，用于实现滑动窗口限流算法。
//
//go:embed slide_window.lua
var slideWindowLua string

// IsLimit 方法检查给定的 key 是否在限流窗口内超过了阈值。
// ctx 是上下文，用于控制限流操作的取消和超时。
// key 是用于限流的唯一标识符，通常是用户 IP 或其他唯一标识。
// 方法返回一个布尔值，指示是否触发限流，以及一个可能的错误。
func (Self *RedisRateLimiter) IsLimit(ctx context.Context, key string) (bool, error) {
	return Self.redisClient.Eval(ctx, slideWindowLua, []string{key}, Self.window.Milliseconds(), Self.threshold, time.Now().UnixMilli()).Bool()
}
