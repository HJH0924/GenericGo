// Package genericgo
/**
* @Project : GenericGo
* @File    : types.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/17 10:32
**/

package genericgo

// Comparator 用于比较两个对象的大小
// left 	<	right	--- 	-1
// left 	==	right	--- 	0
// left 	>	right	--- 	1
type Comparator[T any] func(left T, right T) int

// Zero 根据类型参数 T 返回其零值。
func Zero[T any]() T {
	var zero T
	return zero
}
