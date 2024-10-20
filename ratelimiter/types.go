// Package ratelimiter
/**
* @Project : GenericGo
* @File    : types.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/19 12:09
**/

package ratelimiter

import "context"

type RateLimiter interface {
	IsLimit(ctx context.Context, key string) (bool, error)
}
