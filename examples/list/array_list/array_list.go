// Package array_list
/**
* @Project : GenericGo
* @File    : array_list.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/16 9:51
**/

package main

import (
	"fmt"

	"github.com/HJH0924/GenericGo/list"
)

func ExampleArrayList() {
	// 创建一个新的 ArrayList 实例，初始容量为 10
	al := list.NewArrayList[int](10)
	fmt.Println("Len of array list:", al.Len())
	fmt.Println("Cap of array list:", al.Cap())

	// 添加元素到 ArrayList
	al.Append(1, 2, 3)
	fmt.Println("After Append:", al.AsSlice())

	// 在特定位置添加元素
	err := al.Add(1, 4) // 在索引 1 的位置添加元素 4
	if err != nil {
		fmt.Println("Error adding element:", err)
	}
	fmt.Println("After Add:", al.AsSlice())

	// 获取元素
	val, err := al.Get(1) // 获取索引 1 的元素
	if err != nil {
		fmt.Println("Error getting element:", err)
	}
	fmt.Println("Element at index 1:", val)

	// 设置元素
	err = al.Set(2, 5) // 将索引 2 的元素设置为 5
	if err != nil {
		fmt.Println("Error setting element:", err)
	}
	fmt.Println("After Set:", al.AsSlice())

	// 删除元素
	deletedVal, err := al.Delete(0) // 删除索引 0 的元素
	if err != nil {
		fmt.Println("Error deleting element:", err)
	}
	fmt.Println("Deleted value:", deletedVal)
	fmt.Println("After Delete:", al.AsSlice())

	// 遍历 ArrayList 中的所有元素
	err = al.Range(func(idx int, val int) error {
		fmt.Printf("Index: %d, Value: %d\n", idx, val)
		return nil
	})
	if err != nil {
		fmt.Println("Error during range:", err)
	}

	// 将 ArrayList 转换为切片
	slice := al.AsSlice()
	fmt.Println("Converted slice:", slice)

	// 使用 NewArrayListOf 创建一个新的 ArrayList 实例，使用提供的切片作为底层存储
	existingSlice := []int{7, 8, 9}
	al2 := list.NewArrayListOf(existingSlice)
	fmt.Println("ArrayList created from existing slice:", al2.AsSlice())
}

func main() {
	ExampleArrayList()
}
