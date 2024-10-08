// Package redis
/**
* @Project : GenericGo
* @File    : cache.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/8 16:17
**/

package redis

import (
	"context"
	"errors"
	"github.com/HJH0924/GenericGo/cache"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	_ cache.Cache = (*Cache)(nil)
)

// Cache 是 cache.Cache 接口的实现，用于操作 Redis 缓存。
type Cache struct {
	client redis.Cmdable
}

// Set 设置缓存中的键值对，并可设置过期时间。
func (Self *Cache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	return Self.client.Set(ctx, key, val, expiration).Err()
}

// SetNX (Set if Not eXists) 设置一个键值对，如果键已存在，则返回false，但不会有error，键不存在则可以成功设置，此时返回true
func (Self *Cache) SetNX(ctx context.Context, key string, val any, expiration time.Duration) (bool, error) {
	return Self.client.SetNX(ctx, key, val, expiration).Result()
}

// Get 获取缓存中的值。
func (Self *Cache) Get(ctx context.Context, key string) (any, error) {
	res, err := Self.client.Get(ctx, key).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		err = cache.NewErrKeyNotExist
	}
	return res, err
}

// GetSet 设置缓存中的键值对，并返回旧值。
func (Self *Cache) GetSet(ctx context.Context, key string, val any) (any, error) {
	res, err := Self.client.GetSet(ctx, key, val).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		err = cache.NewErrKeyNotExist
	}
	return res, err
}

// Delete 删除缓存中的一个或多个键。
func (Self *Cache) Delete(ctx context.Context, keys ...string) (int64, error) {
	return Self.client.Del(ctx, keys...).Result()
}

// LPush 将一个或多个值插入到列表的头部。
func (Self *Cache) LPush(ctx context.Context, key string, vals ...any) (int64, error) {
	return Self.client.LPush(ctx, key, vals...).Result()
}

// LPop 从列表头部弹出一个元素。
func (Self *Cache) LPop(ctx context.Context, key string) (any, error) {
	res, err := Self.client.LPop(ctx, key).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		err = cache.NewErrListEmpty
	}
	return res, err
}

// SAdd 将一个或多个成员添加到集合中。
func (Self *Cache) SAdd(ctx context.Context, key string, members ...any) (int64, error) {
	return Self.client.SAdd(ctx, key, members...).Result()
}

// SRem 从集合中移除一个或多个成员。
func (Self *Cache) SRem(ctx context.Context, key string, members ...any) (int64, error) {
	return Self.client.SRem(ctx, key, members...).Result()
}

// IncrBy 增加键对应的整数值。
func (Self *Cache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return Self.client.IncrBy(ctx, key, value).Result()
}

// DecrBy 减少键对应的整数值。
func (Self *Cache) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	return Self.client.DecrBy(ctx, key, decrement).Result()
}

// IncrByFloat 增加键对应的浮点数值。
func (Self *Cache) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	return Self.client.IncrByFloat(ctx, key, value).Result()
}

// NewCache 创建一个新的 Cache 实例。
func NewCache(client redis.Cmdable) *Cache {
	return &Cache{
		client: client,
	}
}
