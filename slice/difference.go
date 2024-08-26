// Package slice
/**
* @Project : GenericGo
* @File    : difference.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/10 16:18
**/

package slice

// Difference 返回两个切片的差集，即在第一个切片中但不在第二个切片中的元素。
// 注意：已去重，并且只支持 comparable 类型，返回的元素顺序不固定
func Difference[T comparable](src1, src2 []T) []T {
	src1Map := toMap(src1)
	for _, value := range src2 {
		delete(src1Map, value)
	}

	res := make([]T, 0, len(src1Map))
	for key := range src1Map {
		res = append(res, key)
	}
	return res
}

// DifferenceFunc 返回两个切片的差集，使用自定义的相等性比较函数。
// isEqual 是一个函数，用于比较两个元素是否相等。
// 注意：已去重，支持任意类型，返回的元素顺序不固定
func DifferenceFunc[T any](src1, src2 []T, isEqual equalFunc[T]) []T {
	res := make([]T, 0, len(src1))
	for _, value := range src1 {
		isContains := ContainsFunc(src2, func(t T) bool {
			return isEqual(t, value)
		})
		if !isContains {
			res = append(res, value)
		}
	}
	return removeDuplicatesFunc(res, isEqual)
}
