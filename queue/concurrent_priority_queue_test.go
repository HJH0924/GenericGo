// Package queue
/**
* @Project : GenericGo
* @File    : concurrent_priority_queue_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/26 17:37
**/

package queue

import (
	"testing"

	"github.com/HJH0924/GenericGo/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConcurrentPriorityQueue(t *testing.T) {
	initialVals := []int{1, 2, 3, 4, 5, 6}
	tests := []struct {
		name     string
		capacity int
		vals     []int
		expected []int
	}{
		{
			name:     "Test with default capacity",
			capacity: 0,
			vals:     initialVals,
			expected: []int{6, 5, 4, 3, 2, 1},
		},
		{
			name:     "Test with specific capacity",
			capacity: len(initialVals),
			vals:     initialVals,
			expected: []int{6, 5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用提供的容量创建新的优先级队列
			pq := NewConcurrentPriorityQueue[int](tt.capacity, getIntComparator())
			// 验证队列初始长度为0
			assert.Equal(t, 0, pq.Len())
			// 向队列添加元素
			for _, val := range tt.vals {
				err := pq.EnQueue(val)
				assert.NoError(t, err)
			}
			// 验证队列容量和长度
			assert.Equal(t, tt.capacity, pq.Cap())
			assert.Equal(t, len(tt.vals), pq.Len())
			// 验证队列中元素顺序是否正确
			res := make([]int, 0, len(tt.vals))
			for !pq.IsEmpty() {
				val, err := pq.DeQueue()
				assert.NoError(t, err)
				res = append(res, val)
			}
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestConcurrentPriorityQueue_Peek(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		vals     []int
		wantErr  error
	}{
		{
			name:     "Test with vals",
			capacity: 0,
			vals:     []int{6, 5, 4, 3, 2, 1},
			wantErr:  errs.NewErrEmptyQueue(),
		},
		{
			name:     "Test empty queue",
			capacity: 0,
			vals:     []int{},
			wantErr:  errs.NewErrEmptyQueue(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewConcurrentPriorityQueueOf(tt.capacity, tt.vals, getIntComparator())
			for !pq.IsEmpty() {
				peek, err := pq.Peek()
				assert.NoError(t, err)
				pop, err := pq.DeQueue()
				assert.NoError(t, err)
				assert.Equal(t, pop, peek)
			}
			_, err := pq.Peek()
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestConcurrentPriorityQueue_EnQueue(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		vals     []int
		enVal    int
		wantErr  error
	}{
		{
			name:     "Bounded empty queue",
			capacity: 10,
			vals:     []int{},
			enVal:    10,
		},
		{
			name:     "Bounded full queue",
			capacity: 10,
			vals:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			enVal:    11,
			wantErr:  errs.NewErrOutOfCapacity(),
		},
		{
			name:     "Bounded queue with available space",
			capacity: 10,
			vals:     []int{1, 2, 3, 4, 5, 6},
			enVal:    10,
		},
		{
			name:     "Boundless empty queue",
			capacity: 0,
			vals:     []int{},
			enVal:    10,
		},
		{
			name:     "Boundless queue with available space",
			capacity: 0,
			vals:     []int{1, 2, 3, 4, 5, 6},
			enVal:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewConcurrentPriorityQueueOf(tt.capacity, tt.vals, getIntComparator())
			require.NotNil(t, pq)
			err := pq.EnQueue(tt.enVal)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.capacity, pq.Cap())
		})
	}
}

func TestConcurrentPriorityQueue_EnQueue2(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		enVal    int
		expected []int
	}{
		{
			name:     "New val is the largest",
			vals:     []int{10, 8, 7, 6, 2},
			enVal:    20,
			expected: []int{20, 8, 10, 6, 2, 7},
		},
		{
			name:     "New element is the smallest",
			vals:     []int{10, 8, 7, 6, 2},
			enVal:    1,
			expected: []int{10, 8, 7, 6, 2, 1},
		},
		{
			name:     "New element fits in between",
			vals:     []int{10, 6, 7, 5, 2},
			enVal:    8,
			expected: []int{10, 6, 8, 5, 2, 7},
		},
		{
			name:     "New element is a duplicate",
			vals:     []int{10, 6, 7, 5, 2},
			enVal:    10,
			expected: []int{10, 6, 10, 5, 2, 7},
		},
	}

	for _, tt := range tests {
		pq := NewConcurrentPriorityQueueOf(0, tt.vals, getIntComparator())
		require.NotNil(t, pq)
		err := pq.EnQueue(tt.enVal)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, pq.AsSlice())
	}
}
