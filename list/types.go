// Package list
/**
* @Project : GenericGo
* @File    : types.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/12 9:39
**/

package list

type List[T any] interface {
	// Append 在 List 末尾追加一个或多个元素。
	Append(vals ...T)

	// Add 在特定下标处增加一个新元素。
	// 如果下标超出合法范围，返回错误。
	// 如果 idx 等于 List 长度，则表示往 List 末端增加元素。
	Add(idx int, val T) error

	// Delete 删除指定下标的元素，并返回被删除的元素。
	// 如果下标超出合法范围，返回错误。
	// 可能会发生缩容。
	Delete(idx int) (T, error)

	// Set 重置指定下标位置的元素为 val。
	// 如果下标超出合法范围，返回错误。
	Set(idx int, val T) error

	// Get 返回对应下标的元素。
	// 如果下标超出合法范围，返回错误。
	Get(idx int) (T, error)

	// Len 返回 List 中元素的数量。
	Len() int

	// Cap 返回 List 的容量。
	Cap() int

	// Range 遍历 List 的所有元素，并使用给定的函数访问每个元素。
	Range(onVal func(idx int, val T) error) error

	// AsSlice 将 List 转化为一个新切片，即使 List 为空，也返回一个长度和容量都为0的切片。
	AsSlice() []T
}
