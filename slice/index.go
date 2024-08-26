// Package slice
/**
* @Project : GenericGo
* @File    : index.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/9 14:11
**/

package slice

// Index 在给定的可比较类型的切片中搜索第一个与目标值相等的元素的索引。
// 如果没有找到，则返回-1。
func Index[T comparable](src []T, target T) int {
	return IndexFunc[T](src, func(t T) bool {
		return t == target
	})
}

// IndexFunc 在给定的切片中搜索第一个满足匹配函数条件的元素的索引。
// 如果没有找到，则返回-1。
// 优先使用 Index
func IndexFunc[T any](src []T, match matchFunc[T]) int {
	for index, value := range src {
		if match(value) {
			return index
		}
	}
	return -1
}

// LastIndex 在给定的可比较类型的切片中搜索最后一个与目标值相等的元素的索引。
// 如果没有找到，则返回-1。
func LastIndex[T comparable](src []T, target T) int {
	return LastIndexFunc[T](src, func(t T) bool {
		return t == target
	})
}

// LastIndexFunc 在给定的切片中搜索最后一个满足匹配函数条件的元素的索引。
// 如果没有找到，则返回-1。
// 优先使用 LastIndex
func LastIndexFunc[T any](src []T, match matchFunc[T]) int {
	for i := len(src) - 1; i >= 0; i-- {
		if match(src[i]) {
			return i
		}
	}
	return -1
}

// IndexAll 在给定的切片中搜索所有满足匹配函数条件的元素的索引。
func IndexAll[T comparable](src []T, target T) []int {
	return IndexAllFunc[T](src, func(t T) bool {
		return t == target
	})
}

// IndexAllFunc 在给定的切片中搜索所有满足匹配函数条件的元素的索引。
// 优先使用 IndexAll
func IndexAllFunc[T any](src []T, match matchFunc[T]) []int {
	indexes := make([]int, 0, len(src))
	for index, value := range src {
		if match(value) {
			indexes = append(indexes, index)
		}
	}
	return indexes
}
