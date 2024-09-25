// Package errs
/**
* @Project : GenericGo
* @File    : error.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/27 14:49
**/

package errs

import (
	"errors"
	"fmt"
)

func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("index %d is out of bounds, length is %d", index, length)
}

func NewErrEmptySlice() error {
	return errors.New("the provided slice is empty, please ensure a non-empty slice is passed for operation")
}

func NewErrEmptyQueue() error {
	return errors.New("the queue is empty, unable to perform operation")
}

func NewErrOutOfCapacity() error {
	return errors.New("capacity exceeded, unable to add more elements")
}
