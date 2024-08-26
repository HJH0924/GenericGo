// Package list
/**
* @Project : GenericGo
* @File    : concurrent_list.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/16 14:14
**/

package list

import "sync"

var (
	_ List[any] = &ConcurrentList[any]{}
)

// ConcurrentList 是一个线程安全的 List 接口包装器，它嵌入了泛型 List 接口，
// 并提供了读写锁来确保并发访问时的数据安全。
type ConcurrentList[T any] struct {
	List[T]
	rwLock sync.RWMutex
}

// Append 在 ConcurrentList 末尾追加一个或多个元素。
func (cl *ConcurrentList[T]) Append(vals ...T) {
	cl.rwLock.Lock() // 写锁，确保独占访问
	defer cl.rwLock.Unlock()
	cl.List.Append(vals...)
}

// Add 在特定下标处增加一个新元素。
// 如果下标超出合法范围，返回错误。
// 如果 idx 等于 ConcurrentList 长度，则表示往 ConcurrentList 末端增加元素。
func (cl *ConcurrentList[T]) Add(idx int, val T) error {
	cl.rwLock.Lock() // 写锁，确保独占访问
	defer cl.rwLock.Unlock()
	return cl.List.Add(idx, val)
}

// Delete 删除指定下标的元素，并返回被删除的元素。
// 如果下标超出合法范围，返回错误。
func (cl *ConcurrentList[T]) Delete(idx int) (T, error) {
	cl.rwLock.Lock() // 写锁，确保独占访问
	defer cl.rwLock.Unlock()
	return cl.List.Delete(idx)
}

// Set 重置指定下标位置的元素为 val。
// 如果下标超出合法范围，返回错误。
func (cl *ConcurrentList[T]) Set(idx int, val T) error {
	cl.rwLock.Lock() // 写锁，确保独占访问
	defer cl.rwLock.Unlock()
	return cl.List.Set(idx, val)
}

// Get 返回对应下标的元素。
// 如果下标超出合法范围，返回错误。
func (cl *ConcurrentList[T]) Get(idx int) (T, error) {
	cl.rwLock.RLock() // 读锁，允许多个读操作
	defer cl.rwLock.RUnlock()
	return cl.List.Get(idx)
}

// Len 返回 ConcurrentList 中元素的数量。
func (cl *ConcurrentList[T]) Len() int {
	cl.rwLock.RLock() // 读锁，允许多个读操作
	defer cl.rwLock.RUnlock()
	return cl.List.Len()
}

// Cap 方法返回 ConcurrentList 的容量。
func (cl *ConcurrentList[T]) Cap() int {
	cl.rwLock.RLock() // 读锁，允许多个读操作
	defer cl.rwLock.RUnlock()
	return cl.List.Cap()
}

// Range 遍历 ConcurrentList 的所有元素，并使用给定的函数访问每个元素。
func (cl *ConcurrentList[T]) Range(onVal func(idx int, val T) error) error {
	cl.rwLock.RLock() // 读锁，允许多个读操作
	defer cl.rwLock.RUnlock()
	return cl.List.Range(onVal)
}

// AsSlice 将 ConcurrentList 转化为一个新切片，即使 ConcurrentList 为空，也返回一个长度和容量都为0的切片。
func (cl *ConcurrentList[T]) AsSlice() []T {
	cl.rwLock.RLock() // 读锁，允许多个读操作
	defer cl.rwLock.RUnlock()
	return cl.List.AsSlice()
}

// NewConcurrentListOf 创建并返回一个新的线程安全的 ConcurrentList 实例。
func NewConcurrentListOf[T any](list List[T]) *ConcurrentList[T] {
	return &ConcurrentList[T]{
		List: list,
	}
}
