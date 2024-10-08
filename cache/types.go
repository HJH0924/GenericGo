// Package cache
/**
* @Project : GenericGo
* @File    : types.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/8 15:48
**/

package cache

import (
	"context"
	"errors"
	"time"
)

var (
	NewErrKeyNotExist = errors.New("key不存在")
	NewErrListEmpty   = errors.New("列表为空")
)

// Cache 定义了缓存操作的接口
type Cache interface {
	// Set 设置一个键值对，并可以设置过期时间。
	// 如果过期时间为0，则表示永不过期。
	Set(ctx context.Context, key string, val any, expiration time.Duration) error

	// SetNX (Set if Not eXists) 设置一个键值对，如果键已存在，则返回false，但不会有error，键不存在则可以成功设置，此时返回true
	// 可以设置过期时间，如果过期时间为0，则表示永不过期。
	SetNX(ctx context.Context, key string, val any, expiration time.Duration) (bool, error)

	// Get 获取键对应的值。
	// 如果键不存在，将返回 errs.ErrKeyNotExist 错误。
	Get(ctx context.Context, key string) (any, error)

	// GetSet 设置一个新的值，并返回旧的值。
	// 如果键不存在，将返回 errs.ErrKeyNotExist 错误。
	GetSet(ctx context.Context, key string, val any) (any, error)

	// Delete 删除一个或多个键。
	// 如果键不存在，则不会计入删除数量，也不会返回错误。
	Delete(ctx context.Context, keys ...string) (int64, error)

	// LPush 将一个或多个值插入到键对应的列表的头部。
	// 如果键不存在，则创建一个空列表。
	// 如果键的值不是列表，则返回错误。
	// 默认返回列表中元素的个数。
	LPush(ctx context.Context, key string, vals ...any) (int64, error)

	// LPop 移除并返回列表的第一个元素。
	LPop(ctx context.Context, key string) (any, error)

	// SAdd 将一个或多个成员添加到集合中。
	// 如果成员已存在，则忽略。
	// 返回添加成功的成员数量。
	SAdd(ctx context.Context, key string, members ...any) (int64, error)

	// SRem 移除集合中的一个或多个成员。
	// 如果成员不存在，则忽略。
	// 返回实际删除的成员数量。
	SRem(ctx context.Context, key string, members ...any) (int64, error)

	// IncrBy 将键对应的值增加指定的整数值。
	// 返回增加后的新值。
	IncrBy(ctx context.Context, key string, value int64) (int64, error)

	// DecrBy 将键对应的值减少指定的整数值。
	// 返回减少后的新值。
	DecrBy(ctx context.Context, key string, decrement int64) (int64, error)

	// IncrByFloat 为键中所储存的值加上指定的浮点数增量值。
	// 返回增加后的新值。
	IncrByFloat(ctx context.Context, key string, value float64) (float64, error)
}
