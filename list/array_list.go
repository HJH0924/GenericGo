// Package list
/**
* @Project : GenericGo
* @File    : array_list.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/12 9:35
**/

package list

import (
	"github.com/HJH0924/GenericGo/errs"
	"github.com/HJH0924/GenericGo/slice"
)

// 该行代码是一个接口实现的编译时断言
// 它用于验证 ArrayList 的指针是否实现了 List 的所有接口
// 这个断言是编译时检查，确保 ArrayList 的实现符合 List 接口规范
var (
	// _ 是一个特殊的变量名，用于忽略未使用的变量值。
	_ List[any] = &ArrayList[any]{}
)

// ArrayList 基于切片的简单封装
type ArrayList[T any] struct {
	vals []T
}

// Append 在 ArrayList 末尾追加一个或多个元素。
func (al *ArrayList[T]) Append(vals ...T) {
	al.vals = append(al.vals, vals...)
}

// Add 在特定下标处增加一个新元素。
// 如果下标超出合法范围，返回错误。
// 如果 idx 等于 ArrayList 长度，则表示往 ArrayList 末端增加元素。
func (al *ArrayList[T]) Add(idx int, val T) (err error) {
	al.vals, err = slice.Add(al.vals, idx, val)
	return err
}

// Delete 删除指定下标的元素，并返回被删除的元素。
// 如果下标超出合法范围，返回错误。
// 可能会发生缩容
func (al *ArrayList[T]) Delete(idx int) (T, error) {
	res, deleteElement, err := slice.Delete(al.vals, idx)
	al.vals = res
	return deleteElement, err
}

// Set 重置指定下标位置的元素为 val。
// 如果下标超出合法范围，返回错误。
func (al *ArrayList[T]) Set(idx int, val T) error {
	length := al.Len()
	if idx < 0 || idx >= length {
		return errs.NewErrIndexOutOfRange(length, idx)
	}
	al.vals[idx] = val
	return nil
}

// Get 返回对应下标的元素。
// 如果下标超出合法范围，返回错误。
func (al *ArrayList[T]) Get(idx int) (val T, err error) {
	length := al.Len()
	if idx < 0 || idx >= length {
		return val, errs.NewErrIndexOutOfRange(length, idx)
	}
	return al.vals[idx], err
}

// Len 返回 ArrayList 中元素的数量。
func (al *ArrayList[T]) Len() int {
	return len(al.vals)
}

// Cap 返回 ArrayList 的容量。
func (al *ArrayList[T]) Cap() int {
	return cap(al.vals)
}

// Range 遍历 ArrayList 的所有元素，并使用给定的函数访问每个元素。
func (al *ArrayList[T]) Range(onVal func(idx int, val T) error) error {
	for index, value := range al.vals {
		err := onVal(index, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// AsSlice 将 ArrayList 转化为一个新切片，即使 ArrayList 为 nil，也返回一个长度和容量都为0的切片。
// 由于返回的是新切片，对返回的切片所做的任何修改都不会影响原始的 ArrayList。
// 但是如果 al.vals 为空，即 []T{}，则返回的 res 新切片 []T{} ，地址与 al.vals 相同，此时共用底层数组
func (al *ArrayList[T]) AsSlice() []T {
	res := make([]T, al.Len())
	copy(res, al.vals)
	return res
}

// NewArrayList 创建并返回一个新的ArrayList实例。
func NewArrayList[T any](initCap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, initCap),
	}
}

// NewArrayListOf 创建一个新的ArrayList实例，使用提供的切片vals作为底层存储。
// 此函数不会复制vals切片，因此对返回的ArrayList实例的任何修改都会反映在原始切片中。
// 为避免意外修改，如果需要保持原始数据不变，应在调用之前创建vals的副本。
func NewArrayListOf[T any](vals []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: vals,
	}
}
