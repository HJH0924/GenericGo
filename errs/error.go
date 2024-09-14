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
	return WrapError(fmt.Sprintf("Index %d is out of bounds, length is %d", index, length))
}

func NewErrEmptySlice() error {
	return WrapError("The provided slice is empty, please ensure a non-empty slice is passed for operation")
}

func NewErrEmptyQueue() error {
	return WrapError("The queue is empty, unable to perform operation")
}

func NewErrOutOfCapacity() error {
	return WrapError("Capacity exceeded, unable to add more elements")
}

func NewErrRBTreeDuplicateNode() error {
	return WrapError("RBTree cannot add duplicate node key value")
}

func NewErrRBTreeNodeNotFound() error {
	return WrapError("The node key value does not exist in the RBTree")
}
