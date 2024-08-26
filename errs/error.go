// Package errs
/**
* @Project : GenericGo
* @File    : error.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/27 14:49
**/

package errs

import "fmt"

func WrapError(message string) error {
	errorPrefix := "GenericGo"
	return fmt.Errorf("[%s]: %s", errorPrefix, message)
}

func NewErrIndexOutOfRange(length int, index int) error {
	return WrapError(fmt.Sprintf("索引越界错误：索引 %d 超出了范围，长度为 %d", index, length))
}

func NewErrEmptySlice() error {
	return WrapError("提供的切片为空，请确保传递非空切片以进行操作")
}

func NewErrEmptyQueue() error {
	return WrapError("队列为空，无法执行操作")
}

func NewErrOutOfCapacity() error {
	return WrapError("超出容量限制，无法添加更多元素")
}
