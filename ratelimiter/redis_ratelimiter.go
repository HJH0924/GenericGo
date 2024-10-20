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
	"github.com/redis/go-redis/v9"
)

type RedisRateLimiter struct {
	redisClient redis.Cmdable
}

func (Self *RedisRateLimiter) IsLimit(ctx context.Context, key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
