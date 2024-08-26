// Package linked_list
/**
* @Project : GenericGo
* @File    : linked_list.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/16 12:14
**/

package main

import (
	"fmt"

	"github.com/HJH0924/GenericGo/list"
)

func ExampleLinkedList() {
	// 创建一个新的 LinkedList 实例
	ll := list.NewLinkedList[int]()
	fmt.Println("Len of linked list:", ll.Len())

	// 添加元素到 LinkedList
	ll.Append(1, 2, 3)
	fmt.Println("After Append:", ll.AsSlice())

	// 在特定位置添加元素
	err := ll.Add(1, 4) // 在索引 1 的位置添加元素 4
	if err != nil {
		fmt.Println("Error adding element:", err)
	}
	fmt.Println("After Add:", ll.AsSlice())

	// 获取元素
	val, err := ll.Get(1) // 获取索引 1 的元素
	if err != nil {
		fmt.Println("Error getting element:", err)
	}
	fmt.Println("Element at index 1:", val)

	// 设置元素
	err = ll.Set(2, 5) // 将索引 2 的元素设置为 5
	if err != nil {
		fmt.Println("Error setting element:", err)
	}
	fmt.Println("After Set:", ll.AsSlice())

	// 删除元素
	deletedVal, err := ll.Delete(0) // 删除索引 0 的元素
	if err != nil {
		fmt.Println("Error deleting element:", err)
	}
	fmt.Println("Deleted value:", deletedVal)
	fmt.Println("After Delete:", ll.AsSlice())

	// 遍历 LinkedList 中的所有元素
	err = ll.Range(func(idx int, val int) error {
		fmt.Printf("Index: %d, Value: %d\n", idx, val)
		return nil
	})
	if err != nil {
		fmt.Println("Error during range:", err)
	}

	// 将 LinkedList 转换为切片
	slice := ll.AsSlice()
	fmt.Println("Converted slice:", slice)

	// 使用 NewLinkedListOf 创建一个新的 LinkedList 实例
	existingSlice := []int{7, 8, 9}
	ll2 := list.NewLinkedListOf(existingSlice)
	fmt.Println("LinkedList created from existing slice:", ll2.AsSlice())
}

func main() {
	ExampleLinkedList()
}
