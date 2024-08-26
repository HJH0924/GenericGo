// Package queue
/**
* @Project : GenericGo
* @File    : priority_queue.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/17 10:44
**/

package queue

import (
	genericgo "github.com/HJH0924/GenericGo"
	"github.com/HJH0924/GenericGo/errs"
	"github.com/HJH0924/GenericGo/slice"
)

// PriorityQueue 是一个基于大根堆的优先级队列
// 当 capacity <= 0 时，为无界队列，切片容量会动态扩缩容
// 当 capacity > 0 时，为有界队列，初始化后就固定容量，不会扩缩容
// 为了方便计算节点的索引，底层切片的第一个元素（索引0）留空，实际元素从索引1开始。
type PriorityQueue[T any] struct {
	compare  genericgo.Comparator[T] // 用于比较两个元素的优先级
	capacity int                     // 优先级队列的容量
	vals     []T                     // 存储优先级队列中元素的切片，索引1为根节点（堆顶）
}

// Len 返回优先级队列中元素的数量（不包括索引0的占位元素）。
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.vals) - 1
}

// Cap 返回优先级队列的容量。
func (pq *PriorityQueue[T]) Cap() int {
	return pq.capacity
}

// IsBoundLess 返回队列是否无界，即容量是否小于或等于0。
func (pq *PriorityQueue[T]) IsBoundLess() bool {
	return pq.Cap() <= 0
}

// IsFull 返回队列是否已满。
func (pq *PriorityQueue[T]) IsFull() bool {
	return pq.Cap() > 0 && pq.Len() == pq.Cap()
}

// IsEmpty 返回队列是否为空。
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.vals) < 2
}

// Peek 返回队列顶部的元素，但不移除它。如果队列为空，返回错误。
func (pq *PriorityQueue[T]) Peek() (T, error) {
	if pq.IsEmpty() {
		return genericgo.Zero[T](), errs.NewErrEmptyQueue()
	}
	return pq.vals[1], nil
}

// EnQueue 向优先级队列添加一个新元素。如果队列已满，返回错误。
func (pq *PriorityQueue[T]) EnQueue(val T) error {
	if pq.IsFull() {
		return errs.NewErrOutOfCapacity()
	}

	pq.vals = append(pq.vals, val)
	pq.upHeapify()
	return nil
}

// upHeapify 重新调整堆以维持大根堆的性质，从堆的末尾元素开始向上调整。
func (pq *PriorityQueue[T]) upHeapify() {
	nodeIndex := pq.Len()

	// 上浮过程：从新元素开始，逐层向上比较并交换，直到满足堆的性质或到达堆顶
	for nodeIndex > 1 {
		parentIndex := nodeIndex / 2
		if pq.compare(pq.vals[nodeIndex], pq.vals[parentIndex]) > 0 {
			// 当前节点的优先级高于父节点
			pq.vals[nodeIndex], pq.vals[parentIndex] = pq.vals[parentIndex], pq.vals[nodeIndex]
			nodeIndex = parentIndex
		} else {
			// 当前节点值的优先级不高于父节点，上浮结束
			break
		}
	}
}

// DeQueue 从优先级队列中移除并返回顶部的元素，即优先级最高的元素。
// 如果队列为空，返回错误。
func (pq *PriorityQueue[T]) DeQueue() (T, error) {
	if pq.IsEmpty() {
		return genericgo.Zero[T](), errs.NewErrEmptyQueue()
	}

	val := pq.vals[1]              // 获取出队元素
	pq.vals[1] = pq.vals[pq.Len()] // 将最后一个元素移动到根节点位置
	pq.vals = pq.vals[:pq.Len()]   // 缩小切片移除最后一个元素

	// 无界队列可能需要缩容
	if pq.IsBoundLess() {
		slice.ShrinkSlice(pq.vals)
	}

	pq.downHeapify()
	return val, nil
}

// downHeapify 重新调整堆，从根节点开始向下调整，以维持大根堆的性质。
func (pq *PriorityQueue[T]) downHeapify() {
	// maxPos 用于记录当前节点及其子节点中最大元素的位置。
	// i 是当前节点的索引，从1开始，即根节点。
	// n 是优先级队列中元素的数量。
	maxPos, i, n := 1, 1, pq.Len()

	for {
		if left := 2 * i; left <= n && pq.compare(pq.vals[left], pq.vals[maxPos]) > 0 {
			// 左孩子存在，且左孩子的优先级高于当前最大节点
			maxPos = left
		}
		if right := 2*i + 1; right <= n && pq.compare(pq.vals[right], pq.vals[maxPos]) > 0 {
			// 右孩子存在，且右孩子的优先级高于当前最大节点
			maxPos = right
		}
		// 如果最小位置没有变化，或者没有子节点，退出循环
		if maxPos == i {
			break
		}
		// 交换当前节点与最大子节点的值
		pq.vals[maxPos], pq.vals[i] = pq.vals[i], pq.vals[maxPos]
		// 更新当前索引为最大位置索引，继续向下调整
		i = maxPos
	}
}

// AsSlice 返回优先级队列中元素的切片副本。
// 该方法创建一个新的切片，包含队列中的所有元素，
// 从队列的顶部（即最大元素）开始到队列的末尾。
// 注意：返回的切片是 vals 的副本，对副本的修改不会影响到原始的 vals。
func (pq *PriorityQueue[T]) AsSlice() []T {
	res := make([]T, pq.Len())
	copy(res, pq.vals[1:])
	return res
}

// NewPriorityQueue 创建一个新的优先级队列。
// 接受容量参数和比较函数，用于确定元素的优先级顺序。
// 当 capacity <= 0 时，视为无界队列，初始大小使用默认值64。
func NewPriorityQueue[T any](capacity int, compare genericgo.Comparator[T]) *PriorityQueue[T] {
	const (
		DefaultInitialCap = 64
	)
	initialCap := capacity + 1
	if capacity <= 0 {
		capacity = 0
		initialCap = DefaultInitialCap
	}
	return &PriorityQueue[T]{
		compare:  compare,
		capacity: capacity,
		vals:     make([]T, 1, initialCap), // +1 为了适应传统的堆实现中的虚拟头节点。
	}
}

// NewPriorityQueueOf 创建一个具有初始值的优先级队列。
// 它接受一个容量、一个包含初始元素的切片，以及一个比较函数。
// 如果添加元素时出现错误，将返回nil。
func NewPriorityQueueOf[T any](capacity int, vals []T, compare genericgo.Comparator[T]) *PriorityQueue[T] {
	pq := NewPriorityQueue[T](capacity, compare)
	for _, val := range vals {
		err := pq.EnQueue(val)
		if err != nil {
			return nil
		}
	}
	return pq
}
