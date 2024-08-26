// Package slice
/**
* @Project : GenericGo
* @File    : reverse.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/5 18:48
**/

package slice

// Reverse 函数接受一个类型为T的切片，并返回这个切片的逆序副本
func Reverse[T any](src []T) []T {
	res := make([]T, 0, len(src))
	for i := len(src) - 1; i >= 0; i-- {
		res = append(res, src[i])
	}
	return res
}

// ReverseInPlace 函数接受一个类型为T的切片，并在原地逆序该切片
func ReverseInPlace[T any](src []T) {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}
