// Package slice
/**
* @Project : GenericGo
* @File    : aggregate.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/31 20:47
**/

package slice

import (
	genericgo "github.com/HJH0924/GenericGo"
	"github.com/HJH0924/GenericGo/errs"
)

// Max 函数返回给定切片中的最大值
// 如果切片为空，将返回 NewErrEmptySlice 错误
// 在使用 float32 或者 float64 的时候要小心精度问题
func Max[T genericgo.RealNumber](src []T) (T, error) {
	if len(src) == 0 {
		return genericgo.Zero[T](), errs.NewErrEmptySlice()
	}
	res := src[0]
	for i := 1; i < len(src); i++ {
		if src[i] > res {
			res = src[i]
		}
	}
	return res, nil
}

// Min 函数返回给定切片中的最小值
// 如果切片为空，将返回 NewErrEmptySlice 错误
// 在使用 float32 或者 float64 的时候要小心精度问题
func Min[T genericgo.RealNumber](src []T) (T, error) {
	if len(src) == 0 {
		return genericgo.Zero[T](), errs.NewErrEmptySlice()
	}
	res := src[0]
	for i := 1; i < len(src); i++ {
		if src[i] < res {
			res = src[i]
		}
	}
	return res, nil
}

// Sum 函数返回给定切片中所有元素的总和
// 如果切片为空，将返回 0 值和 NewErrEmptySlice 错误
// 在使用 float32 或者 float64 的时候要小心精度问题
func Sum[T genericgo.RealNumber](src []T) (T, error) {
	var res T
	if len(src) == 0 {
		return res, errs.NewErrEmptySlice()
	}
	for _, num := range src {
		res += num
	}
	return res, nil
}
