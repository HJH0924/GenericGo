// Package genericgo
/**
* @Project : GenericGo
* @File    : constrain.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/31 20:46
**/

package genericgo

// RealNumber 实数
// 绝大多数情况下，都应该用这个来表达数字的含义
type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

// Number 包含实数和复数
type Number interface {
	RealNumber | ~complex64 | ~complex128
}
