// Package slice
/**
* @Project : GenericGo
* @File    : union.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/10 15:44
**/

package slice

// Union 计算两个切片的并集，返回一个包含所有不重复元素的新切片。
// 注意：已去重，并且只支持 comparable 类型，返回的元素顺序不固定
func Union[T comparable](src1, src2 []T) []T {
	src1Map := toMap(src1)

	// 将第二个切片的元素添加到映射中。
	for _, value := range src2 {
		src1Map[value] = struct{}{}
	}

	res := make([]T, 0, len(src2))
	for key := range src1Map {
		res = append(res, key)
	}
	return res
}

// UnionFunc 计算两个切片的并集，使用自定义的相等性比较函数。
// 注意：已去重，支持任意类型，返回的元素顺序不固定
func UnionFunc[T any](src1, src2 []T, isEqual equalFunc[T]) []T {
	res := make([]T, 0, len(src1)+len(src2))
	res = append(res, src1...)
	res = append(res, src2...)
	return removeDuplicatesFunc(res, isEqual)
}
