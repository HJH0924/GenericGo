// Package slice
/**
* @Project : GenericGo
* @File    : find.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/5 21:56
**/

package slice

import genericgo "github.com/HJH0924/GenericGo"

// Find 函数在给定的切片中查找第一个满足 matchFunc 函数条件的元素。
// 如果找到这样的元素，函数返回该元素和 true；
// 如果没有找到，返回类型 T 的零值和 false。
func Find[T any](src []T, match matchFunc[T]) (T, bool) {
	for _, val := range src {
		if match(val) {
			return val, true
		}
	}
	return genericgo.Zero[T](), false
}

// FindAll 函数在给定的切片中查找所有满足 matchFunc 函数条件的元素。
// 函数始终返回一个切片（可能是空切片），如果至少有一个元素满足条件，返回 true；否则返回 false。
func FindAll[T any](src []T, match matchFunc[T]) ([]T, bool) {
	// 我们认为符合条件元素应该是少数
	// 所以会除以 8
	// 也就是触发扩容的情况下，最多三次就会和原本的容量一样
	// 这样做的目的是减少在 append 过程中可能发生的切片扩容次数
	// +1 是为了确保即使没有元素符合条件，也能存储至少一个元素
	res := make([]T, 0, len(src)>>3+1)
	for _, val := range src {
		if match(val) {
			res = append(res, val)
		}
	}
	return res, len(res) > 0
}
