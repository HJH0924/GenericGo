// Package queue
/**
* @Project : GenericGo
* @File    : concurrent_priority_queue.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/26 17:03
**/

package queue

import (
	"sync"

	genericgo "github.com/HJH0924/GenericGo"
)

// ConcurrentPriorityQueue 并发安全的优先级队列
type ConcurrentPriorityQueue[T any] struct {
	pq     PriorityQueue[T]
	rwLock sync.RWMutex
}

// Len 返回 ConcurrentPriorityQueue 中元素的数量（不包括索引0的占位元素）。
func (Self *ConcurrentPriorityQueue[T]) Len() int {
	Self.rwLock.RLock()
	defer Self.rwLock.RUnlock()
	return Self.pq.Len()
}

// Cap 返回 ConcurrentPriorityQueue 的容量。
func (Self *ConcurrentPriorityQueue[T]) Cap() int {
	Self.rwLock.RLock()
	defer Self.rwLock.RUnlock()
	return Self.pq.Cap()
}

// IsBoundLess 返回队列是否无界，即容量是否小于或等于0。
func (Self *ConcurrentPriorityQueue[T]) IsBoundLess() bool {
	return Self.pq.Cap() <= 0
}

// IsFull 返回队列是否已满。
func (Self *ConcurrentPriorityQueue[T]) IsFull() bool {
	return Self.pq.Cap() > 0 && Self.pq.Len() == Self.pq.Cap()
}

// IsEmpty 返回 ConcurrentPriorityQueue 是否为空。
func (Self *ConcurrentPriorityQueue[T]) IsEmpty() bool {
	return len(Self.pq.vals) < 2
}

// Peek 返回队列顶部的元素，但不移除它。如果队列为空，返回错误。
func (Self *ConcurrentPriorityQueue[T]) Peek() (T, error) {
	Self.rwLock.RLock()
	defer Self.rwLock.RUnlock()
	return Self.pq.Peek()
}

// EnQueue 向 ConcurrentPriorityQueue 添加一个新元素。如果队列已满，返回错误。
func (Self *ConcurrentPriorityQueue[T]) EnQueue(val T) error {
	Self.rwLock.Lock()
	defer Self.rwLock.Unlock()
	return Self.pq.EnQueue(val)
}

// DeQueue 从 ConcurrentPriorityQueue 中移除并返回顶部的元素，即优先级最高的元素。
// 如果队列为空，返回错误。
func (Self *ConcurrentPriorityQueue[T]) DeQueue() (T, error) {
	Self.rwLock.Lock()
	defer Self.rwLock.Unlock()
	return Self.pq.DeQueue()
}

// AsSlice 返回 ConcurrentPriorityQueue 中元素的切片副本。
// 该方法创建一个新的切片，包含队列中的所有元素，
// 从队列的顶部（即最大元素）开始到队列的末尾。
// 注意：返回的切片是 vals 的副本，对副本的修改不会影响到原始的 vals。
func (Self *ConcurrentPriorityQueue[T]) AsSlice() []T {
	Self.rwLock.RLock()
	defer Self.rwLock.RUnlock()
	return Self.pq.AsSlice()
}

// NewConcurrentPriorityQueue 创建一个新的 ConcurrentPriorityQueue。
// 接受容量参数和比较函数，用于确定元素的优先级顺序。
// 当 capacity <= 0 时，视为无界队列，初始大小使用默认值64。
func NewConcurrentPriorityQueue[T any](capacity int, compare genericgo.Comparator[T]) *ConcurrentPriorityQueue[T] {
	return &ConcurrentPriorityQueue[T]{
		pq: *NewPriorityQueue[T](capacity, compare),
	}
}

// NewConcurrentPriorityQueueOf 创建一个具有初始值的 ConcurrentPriorityQueue。
// 它接受一个容量、一个包含初始元素的切片，以及一个比较函数。
// 如果添加元素时出现错误，将返回nil。
func NewConcurrentPriorityQueueOf[T any](capacity int, vals []T, compare genericgo.Comparator[T]) *ConcurrentPriorityQueue[T] {
	return &ConcurrentPriorityQueue[T]{
		pq: *NewPriorityQueueOf(capacity, vals, compare),
	}
}
