// Package slice
/**
* @Project : GenericGo
* @File    : intersect.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/10 15:11
**/

package slice

// Intersection 返回两个切片的交集，即在两个切片中都存在的元素。
// 注意：已去重，并且只支持 comparable 类型，返回的元素顺序不固定
func Intersection[T comparable](src1, src2 []T) []T {
	src1Map := toMap(src1)
	res := make([]T, 0, len(src1))
	for _, value := range src2 {
		if _, exists := src1Map[value]; exists {
			res = append(res, value)
		}
	}
	return removeDuplicates(res)
}

// IntersectionFunc 返回两个切片的交集，使用自定义的相等性比较函数。
// isEqual 是一个函数，用于比较两个元素是否相等。
// 注意：已去重，支持任意类型，返回的元素顺序不固定
func IntersectionFunc[T any](src1, src2 []T, isEqual equalFunc[T]) []T {
	res := make([]T, 0, len(src1))
	for _, value := range src2 {
		isContains := ContainsFunc(src1, func(t T) bool {
			return isEqual(t, value)
		})
		if isContains {
			res = append(res, value)
		}
	}
	return removeDuplicatesFunc(res, isEqual)
}
