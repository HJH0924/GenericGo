// Package lru
/**
* @Project : GenericGo
* @File    : cache.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/11 14:33
**/

package lru

import (
	"context"
	"time"

	"github.com/HJH0924/GenericGo/cache"
)

var (
	_ cache.Cache = (*Cache)(nil)
)

// Cache 是 cache.Cache 接口的实现，用于操作 LRU 缓存。
// LRU - Least Recently Used
type Cache struct {
}

func (Self *Cache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) SetNX(ctx context.Context, key string, val any, expiration time.Duration) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) Get(ctx context.Context, key string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) GetSet(ctx context.Context, key string, val any) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) Delete(ctx context.Context, keys ...string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) LPush(ctx context.Context, key string, vals ...any) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) LPop(ctx context.Context, key string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) SAdd(ctx context.Context, key string, members ...any) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) SRem(ctx context.Context, key string, members ...any) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *Cache) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	//TODO implement me
	panic("implement me")
}
