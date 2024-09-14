// Package tuple
/**
* @Project : GenericGo
* @File    : pair.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/14 10:35
**/

package tuple

import (
	"fmt"

	"github.com/HJH0924/GenericGo/errs"
)

// Pair 键值对
type Pair[K any, V any] struct {
	Key K
	Val V
}

// ToString 方法返回 Pair 的字符串表示形式。
func (Self *Pair[K, V]) ToString() string {
	return fmt.Sprintf("<%#v, %#v>", Self.Key, Self.Val)
}

// Split 方法将 Pair 的 Key 和 Value 作为返回参数传出。
func (Self *Pair[K, V]) Split() (K, V) {
	return Self.Key, Self.Val
}

// NewPair 创建一个新的 Pair 实例。
func NewPair[K any, V any](key K, val V) Pair[K, V] {
	return Pair[K, V]{
		Key: key,
		Val: val,
	}
}

// NewPairs 根据传入的 keys 和 values 切片创建并返回一个 Pair 数组。
// 如果 keys 或 values 为 nil，或者它们的长度不相等，将返回错误。
func NewPairs[K any, V any](keys []K, vals []V) ([]Pair[K, V], error) {
	if keys == nil || vals == nil {
		return nil, NewErrNilKeysVals()
	}
	keysLen := len(keys)
	valsLen := len(vals)
	if keysLen != valsLen {
		return nil, NewErrUnequalLengthsOfKeysVals(keysLen, valsLen)
	}
	pairs := make([]Pair[K, V], keysLen)
	for i := 0; i < keysLen; i++ {
		pairs[i] = NewPair(keys[i], vals[i])
	}
	return pairs, nil
}

// SplitPairs 接受一个 Pair 切片并返回两个切片，分别包含所有的 Key 和 Value。
func SplitPairs[K any, V any](pairs []Pair[K, V]) ([]K, []V) {
	if pairs == nil {
		return nil, nil
	}
	n := len(pairs)
	keys := make([]K, n)
	vals := make([]V, n)
	for i, pair := range pairs {
		keys[i], vals[i] = pair.Split()
	}
	return keys, vals
}

// FlattenPairs 将一个 Pair 切片转换为一个单一的切片，其中交替包含所有的 key 和 val。
// 与 PackPairs 配套使用。
func FlattenPairs[K any, V any](pairs []Pair[K, V]) []any {
	if pairs == nil {
		return nil
	}
	flatPairs := make([]any, 0, len(pairs)*2)
	for _, pair := range pairs {
		flatPairs = append(flatPairs, pair.Key, pair.Val)
	}
	return flatPairs
}

// PackPairs 接受一个交替包含 Key 和 Value 的切片，并返回一个 Pair 切片。
// 如果 flatPairs 不满足特定的条件，将返回错误。
// 与 FlattenPairs 配套使用。
func PackPairs[K any, V any](flatPairs []any) ([]Pair[K, V], error) {
	if flatPairs == nil {
		return nil, nil
	}
	n := len(flatPairs)
	if n%2 != 0 {
		return nil, NewErrInvalidFlatPairsLength()
	}
	pairs := make([]Pair[K, V], n/2)
	for i := 0; i < n; i += 2 {
		key, ok := flatPairs[i].(K)
		if !ok {
			return nil, NewErrTypeAssertionForKey(i)
		}
		val, ok := flatPairs[i+1].(V)
		if !ok {
			return nil, NewErrTypeAssertionForVal(i + 1)
		}
		pairs[i/2] = NewPair(key, val)
	}
	return pairs, nil
}

// NewErrNilKeysVals 创建一个 keys 和 vals 不能为 nil 的错误。
func NewErrNilKeysVals() error {
	return errs.WrapError("keys and vals must not be nil")
}

// NewErrUnequalLengthsOfKeysVals 创建一个 keys 和 vals 长度必须相等的错误。
func NewErrUnequalLengthsOfKeysVals(keysLen int, valsLen int) error {
	return errs.WrapError(fmt.Sprintf("length of keys and vals must be equal, len(keys)=%d, len(vals)=%d", keysLen, valsLen))
}

// NewErrInvalidFlatPairsLength 创建一个 flatPairs 长度必须为偶数的错误。
func NewErrInvalidFlatPairsLength() error {
	return errs.WrapError("flatPairs length must be even")
}

// NewErrTypeAssertionForKey 创建一个 key 类型断言失败的错误。
func NewErrTypeAssertionForKey(index int) error {
	return errs.WrapError(fmt.Sprintf("type assertion failed for key at index %d", index))
}

// NewErrTypeAssertionForVal 创建一个 val 类型断言失败的错误。
func NewErrTypeAssertionForVal(index int) error {
	return errs.WrapError(fmt.Sprintf("type assertion failed for val at index %d", index))
}
