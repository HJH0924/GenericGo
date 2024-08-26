// Package list
/**
* @Project : GenericGo
* @File    : concurrent_list.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/16 17:19
**/

package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/HJH0924/GenericGo/list"
	"github.com/HJH0924/GenericGo/slice"
)

func ExampleConcurrentListAppend() {
	const (
		numElements          = 1000                        // 要添加的元素数量
		numGoroutines        = 10                          // 用于并发追加的goroutine数量
		elementsPerGoroutine = numElements / numGoroutines // 每个goroutine要处理的元素数量
	)

	// 准备测试数据
	testData := make([]int, numElements)
	for i := range testData {
		testData[i] = rand.Int() % 100 // 0-99
	}

	cl := list.NewConcurrentListOf(list.List[int](list.NewArrayList[int](0)))
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		start := i * elementsPerGoroutine       // 0 100 200 ...
		end := start + elementsPerGoroutine - 1 // 99 199 299 ...

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j <= end; j++ {
				cl.Append(testData[j])
			}
		}(start, end)
	}

	wg.Wait() // 等待所有goroutine完成

	// 验证 List 长度是否正确
	if cl.Len() == numElements {
		fmt.Println("List length is correct.")
	} else {
		fmt.Printf("Expected list length to be %d, but got %d\n", numElements, cl.Len())
	}

	// 验证所有元素是否被追加到了 List 中
	if slice.ContainsAll(testData, cl.AsSlice()) {
		fmt.Println("All elements are correctly appended to the list.")
	} else {
		fmt.Println("Some elements are missing.")
	}
}

func main() {
	ExampleConcurrentListAppend()
}
