// Package slice
/**
* @Project : GenericGo
* @File    : delete.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/27 17:08
**/

package slice

import (
	genericgo "github.com/HJH0924/GenericGo"
	"github.com/HJH0924/GenericGo/errs"
)

// Delete 删除 index 处的元素
// 返回删除后的切片，删除的元素，以及错误信息
func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return nil, genericgo.Zero[T](), errs.NewErrIndexOutOfRange(length, index)
	}

	// 保存删除的元素
	deleteElement := src[index]

	// 元素前移
	for i := index; i < length-1; i++ {
		src[i] = src[i+1]
	}

	// 删除最后一个元素，可能需要进行缩容
	return ShrinkSlice[T](src[:length-1]), deleteElement, nil
}

// DeleteIf 删除满足条件(condition)的元素
// 所有操作都会在原切片上进行，以提高性能。
// 被删除元素之后的元素会向前移动，有且只会移动一次。
// condition func(index int, value T) bool：一个函数，用于判断切片中的元素是否应该被删除。
// 用户可以根据索引或元素的具体值来判断是否删除
func DeleteIf[T any](src []T, condition func(index int, value T) bool) []T {
	// 记录没有被删除元素的索引位置
	keepIndex := 0

	for index, value := range src {
		if condition(index, value) {
			// 满足删除条件，跳过，不保留
			continue
		}
		// 不满足删除条件，保留
		src[keepIndex] = value
		keepIndex++
	}

	// 返回未被删除元素构成的新切片，可能需要进行缩容
	return ShrinkSlice[T](src[:keepIndex])
}
