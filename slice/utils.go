// Package slice
/**
* @Project : GenericGo
* @File    : utils.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/10 11:53
**/

package slice

// toMap 创建一个映射，使用源切片中的元素作为键，每个键对应的值是一个空结构体。
// 使用空结构体作为值是为了减少内存消耗，同时利用映射来快速检查元素的存在性。
// 函数返回一个映射，其中包含切片中所有元素作为键。
func toMap[T comparable](src []T) map[T]struct{} {
	resMap := make(map[T]struct{}, len(src))
	for _, value := range src {
		resMap[value] = struct{}{}
	}
	return resMap
}

// removeDuplicates 函数去除给定切片中的重复元素，并返回一个新的不包含重复元素的切片。
func removeDuplicates[T comparable](src []T) []T {
	srcMap := toMap(src)
	uniqueSrc := make([]T, 0, len(srcMap))
	for key := range srcMap {
		uniqueSrc = append(uniqueSrc, key)
	}
	return uniqueSrc
}

// removeDuplicatesFunc 使用自定义的相等性比较函数去除切片中的重复元素，并返回一个新的不包含重复元素的切片。
// isEqual 函数定义了如何比较两个元素是否相等。
func removeDuplicatesFunc[T any](src []T, isEqual equalFunc[T]) []T {
	uniqueSrc := make([]T, 0, len(src))
	for index, value := range src {
		isContains := ContainsFunc(src[index+1:], func(t T) bool {
			return isEqual(t, value)
		})
		if !isContains {
			uniqueSrc = append(uniqueSrc, value)
		}
	}
	return uniqueSrc
}
