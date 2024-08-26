// Package slice
/**
* @Project : GenericGo
* @File    : symmetric_difference.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/10 16:38
**/

package slice

// SymmetricDifference 返回两个切片的对称差集，即两个切片中不重叠的元素组成的切片。
// 注意：已去重，并且只支持 comparable 类型，返回的元素顺序不固定
func SymmetricDifference[T comparable](src1, src2 []T) []T {
	src1Map, src2Map := toMap(src1), toMap(src2)
	for key := range src1Map {
		if _, exists := src2Map[key]; exists {
			delete(src2Map, key)
		} else {
			src2Map[key] = struct{}{}
		}
	}

	res := make([]T, 0, len(src2Map))
	for key := range src2Map {
		res = append(res, key)
	}
	return res
}

// SymmetricDifferenceFunc 返回两个切片的对称差集，使用自定义的相等性比较函数。
// isEqual 是一个函数，用于比较两个元素是否相等。
// 注意：已去重，支持任意类型，返回的元素顺序不固定
func SymmetricDifferenceFunc[T any](src1, src2 []T, isEqual equalFunc[T]) []T {
	var res []T

	// 找出第一个切片中不在第二个切片中的元素
	for _, value := range src1 {
		if !ContainsFunc(src2, func(t T) bool {
			return isEqual(t, value)
		}) {
			res = append(res, value)
		}
	}

	// 找出第二个切片中不在第一个切片中的元素
	for _, value := range src2 {
		if !ContainsFunc(src1, func(t T) bool {
			return isEqual(t, value)
		}) {
			res = append(res, value)
		}
	}

	// 去重
	return removeDuplicatesFunc(res, isEqual)
}
