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
	"math/rand"
	"sync"
	"testing"

	"github.com/HJH0924/GenericGo/errs"
	"github.com/stretchr/testify/assert"
)

// 多个 goroutine 执行入队操作，完成后，主协程把元素逐一出队，只要有序，可以认为并发入队没有问题
func TestConcurrentPriorityQueue_EnQueue(t *testing.T) {
	tests := []struct {
		name          string
		capacity      int
		numElements   int // 要入队的元素数量
		numGoroutines int // 用于并发的 goroutine 数量
		wantErrCount  int
	}{
		{
			name:          "Test with vals less than capacity",
			capacity:      10010,
			numElements:   10000,
			numGoroutines: 100,
		},
		{
			name:          "Test with vals more than capacity",
			capacity:      10010,
			numElements:   10100,
			numGoroutines: 100,
			wantErrCount:  90,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elementsPerGoroutine := tt.numElements / tt.numGoroutines // 每个 goroutine 要处理的元素数量

			// 准备测试数据
			testData := make([]int, tt.numElements)
			for i := range testData {
				testData[i] = rand.Int() % 100 // 0-99
			}

			pq := NewConcurrentPriorityQueue(tt.capacity, getIntComparator())

			errChan := make(chan error, tt.numGoroutines)

			var wg sync.WaitGroup
			for i := 0; i < tt.numGoroutines; i++ {
				start := i * elementsPerGoroutine       // 0 100 200 ...
				end := start + elementsPerGoroutine - 1 // 99 199 299 ...

				wg.Add(1)
				go func(start, end int) {
					defer wg.Done()
					for j := start; j <= end; j++ {
						err := pq.EnQueue(testData[j])
						if err != nil {
							errChan <- err
						}
					}
				}(start, end)
			}

			wg.Wait() // 等待所有 goroutine 完成

			close(errChan)

			// 验证错误数量是否符合预期
			assert.Equal(t, tt.wantErrCount, len(errChan))

			// 验证元素是否按预期顺序出队
			prev := 100
			for !pq.IsEmpty() {
				getDeVal, _ := pq.DeQueue()
				assert.GreaterOrEqual(t, prev, getDeVal)
				prev = getDeVal
			}
		})
	}
}

// 预先入队一组数据，通过测试多个协程并发出队时，每个协程内出队元素有序，间接确认并发安全
func TestConcurrentPriorityQueue_DeQueue(t *testing.T) {
	tests := []struct {
		name                 string
		numElements          int // 要出队的元素数量
		numGoroutines        int // 用于并发的 goroutine 数量
		elementsPerGoroutine int // 每个 goroutine 要处理的元素数量
		remain               int // 队列中剩余的元素数量
		wantErrCount         int
	}{
		{
			name:                 "More Elements Enqueued Than Dequeued",
			numElements:          10000,
			numGoroutines:        99,
			elementsPerGoroutine: 100,
			remain:               10000 - 99*100,
			wantErrCount:         0,
		},
		{
			name:                 "Fewer Elements Enqueued Than Dequeued",
			numElements:          10000,
			numGoroutines:        101,
			elementsPerGoroutine: 100,
			remain:               0,
			wantErrCount:         100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 准备测试数据
			testData := make([]int, tt.numElements)
			for i := range testData {
				testData[i] = rand.Int() % 1000 // 0-999
			}

			// 预先入队一组数据
			pq := NewConcurrentPriorityQueueOf(tt.numElements, testData, getIntComparator())

			errChan := make(chan error, tt.numGoroutines*tt.elementsPerGoroutine)
			disOrderChan := make(chan bool, tt.numGoroutines*tt.elementsPerGoroutine)

			var wg sync.WaitGroup
			for i := 0; i < tt.numGoroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					prev := 1000
					for i := 0; i < tt.elementsPerGoroutine; i++ {
						deVal, err := pq.DeQueue()
						if err != nil {
							errChan <- err
						} else {
							if prev < deVal {
								disOrderChan <- true
							}
							prev = deVal
						}
					}
				}()
			}

			wg.Wait() // 等待所有 goroutine 完成

			close(errChan)
			close(disOrderChan)

			// 验证错误数量是否符合预期
			assert.Equal(t, tt.wantErrCount, len(errChan))
			for err := range errChan {
				assert.Equal(t, err, errs.NewErrEmptyQueue())
			}

			// 验证是否所有元素都按预期顺序出队
			assert.Equal(t, 0, len(disOrderChan))

			// 验证队列中剩余的元素数量是否符合预期
			assert.Equal(t, tt.remain, pq.Len())
		})
	}
}
