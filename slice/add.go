// Package slice
/**
* @Project : GenericGo
* @File    : add.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/27 15:05
**/

package slice

import (
	genericgo "github.com/HJH0924/GenericGo"
	"github.com/HJH0924/GenericGo/errs"
)

// Add 在index处添加元素
// index 范围应为[0, len(src)]
// 如果 index == len(src) 则表示往末尾添加元素
func Add[T any](src []T, index int, element T) ([]T, error) {
	if length := len(src); index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}

	// 先往尾部添加一个零元素
	src = append(src, genericgo.Zero[T]())

	// 元素后移
	for i := len(src) - 1; i > index; i-- {
		src[i] = src[i-1]
	}

	// 最后直接通过索引赋值
	src[index] = element
	return src, nil
}
