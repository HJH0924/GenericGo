// Package slice
/**
* @Project : GenericGo
* @File    : slice.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/27 16:49
**/

package main

import (
	"fmt"

	"github.com/HJH0924/GenericGo/slice"
)

func ExampleSliceAdd() {
	fmt.Println("ExampleSliceAdd:")
	// 使用 GenericGo 的 Add 函数向切片添加元素
	res, err := slice.Add[int]([]int{3, 5, 7}, 0, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result slice:", res)
	}

	res, err = slice.Add[int]([]int{3, 5, 7}, -1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result slice:", res)
	}
}

func ExampleSliceDelete() {
	fmt.Println("ExampleSliceDelete:")
	// 使用 GenericGo 的 Delete 函数删除切片的某个元素
	res, elem, err := slice.Delete[string]([]string{"hello", "world"}, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result slice:", res)
		fmt.Println("Deleted element:", elem)
	}

	res, elem, err = slice.Delete[string]([]string{"hello", "world"}, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result slice:", res)
		fmt.Println("Deleted element:", elem)
	}
}

func ExampleSliceDeleteIf() {
	fmt.Println("ExampleSliceDeleteIf:")
	// 使用 GenericGo 的 DeleteIf 函数删除切片中满足条件的元素
	// 删除值大于4的元素
	res := slice.DeleteIf[float64]([]float64{3.14, 3.45, 4.6}, func(index int, value float64) bool {
		return value > 4
	})
	fmt.Println("Result slice:", res)

	// 删除第一个元素
	res = slice.DeleteIf[float64]([]float64{3.14, 3.45, 4.6}, func(index int, value float64) bool {
		return index == 0
	})
	fmt.Println("Result slice:", res)
}

func ExampleSliceShrink() {
	fmt.Println("ExampleSliceShrink:")

	originalSlice := make([]int, 0, 1000)
	for i := 0; i < 100; i++ {
		originalSlice = append(originalSlice, i)
	}
	fmt.Println("Original slice:", originalSlice)
	fmt.Printf("Original capacity: %d, length: %d\n", cap(originalSlice), len(originalSlice))

	// 使用 GenericGo 的 ShrinkSlice 函数对切片进行缩容
	sliceAfterShrink := slice.ShrinkSlice[int](originalSlice)
	fmt.Println("Slice after shrink:", sliceAfterShrink)
	fmt.Printf("New Capacity: %d, length: %d\n", cap(sliceAfterShrink), len(sliceAfterShrink))
}

func ExampleSliceDeleteShrink() {
	fmt.Println("ExampleSliceDeleteShrink:")
	originalCapacity := 3000
	src := make([]int, 0, originalCapacity)
	elementsToAdd := originalCapacity
	for i := 0; i < elementsToAdd; i++ {
		src = append(src, i)
	}
	fmt.Println("Original slice:", src)
	fmt.Printf("Original cap: %d, len: %d\n", cap(src), len(src))
	res := slice.DeleteIf[int](src, func(index int, value int) bool {
		return value > 600
	})
	fmt.Println("Result slice:", res)
	fmt.Printf("Result cap: %d, len: %d\n", cap(res), len(res))
}

func ExampleSliceAggregate() {
	fmt.Println("ExampleSliceAggregate:")
	intSlice := []int{1, 3, 2, 4, 5}
	maxInt, err := slice.Max(intSlice)
	if err == nil {
		fmt.Println("Max int:", maxInt)
	}

	minInt, err := slice.Min(intSlice)
	if err == nil {
		fmt.Println("Min int:", minInt)
	}

	sumInt, err := slice.Sum(intSlice)
	if err == nil {
		fmt.Println("Sum int:", sumInt)
	}

	floatSlice := []float64{1.5, 3.2, 2.8, 4.1}
	maxFloat, err := slice.Max(floatSlice)
	if err == nil {
		fmt.Println("Max float:", maxFloat)
	}

	minFloat, err := slice.Min(floatSlice)
	if err == nil {
		fmt.Println("Min float:", minFloat)
	}

	sumFloat, err := slice.Sum(floatSlice)
	if err == nil {
		fmt.Println("Sum float:", sumFloat)
	}
}

func ExampleSliceReverse() {
	fmt.Println("ExampleSliceReverse:")
	floatSlice := []float64{1.5, 3.2, 2.8, 4.1}
	fmt.Printf("Float slice address: %p\n", &floatSlice)
	fmt.Println("Original float slice:", floatSlice)
	res := slice.Reverse[float64](floatSlice)
	fmt.Println("Reversed float slice:", res)
	fmt.Printf("Reversed float slice address: %p\n", &res)

	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Int slice address: %p\n", &intSlice)
	fmt.Println("Original int slice:", intSlice)
	slice.ReverseInPlace[int](intSlice)
	fmt.Println("Reversed int slice:", intSlice)
	fmt.Printf("Reversed int slice address: %p\n", &intSlice)
}

func ExampleSliceFind() {
	fmt.Println("ExampleSliceFind:")
	intSlice := []int{4, 1, 3, 2, 5}
	match := func(x int) bool {
		return x > 2
	}
	result, found := slice.Find[int](intSlice, match)
	if found {
		fmt.Printf("Found the first element greater than 2: %d\n", result)
	} else {
		fmt.Println("No element greater than 2 found.")
	}

	intSlice2 := []int{4, 1, 3, 2, 5}
	match2 := func(x int) bool {
		return x > 2
	}
	result2, found2 := slice.FindAll[int](intSlice2, match2)
	if found2 {
		fmt.Println("Found all elements greater than 2:", result2)
	} else {
		fmt.Println("No element greater than 2 found.")
	}
}

func ExampleSliceIndex() {
	fmt.Println("ExampleSliceIndex:")
	intSlice := []int{1, 2, 3, 4, 4, 5, 4}
	target := 4

	// 使用Index查找目标值的第一个出现位置
	firstIndex := slice.Index[int](intSlice, target)
	fmt.Printf("First occurrence of %d is at index: %d\n", target, firstIndex)

	// 使用LastIndex查找目标值的最后一个出现位置
	lastIndex := slice.LastIndex[int](intSlice, target)
	fmt.Printf("Last occurrence of %d is at index: %d\n", target, lastIndex)

	// 使用IndexAll查找目标值的所有出现位置
	allIndexes := slice.IndexAll[int](intSlice, target)
	fmt.Printf("All occurrences of %d are at indexes: %v\n", target, allIndexes)
}

func ExampleSliceMap() {
	fmt.Println("ExampleSliceMap:")

	// 计算每个数字出现的次数
	nums := []int{1, 2, 2, 3, 3, 3, 4}
	resMap := slice.ToMap(nums, func(num int) int {
		return num
	})
	fmt.Println("nums to map:", resMap)
	fmt.Println("Unique nums: ", len(resMap))

	// 结构体数组转map
	type Person struct {
		Name string
		Age  int
	}

	persons := []Person{
		{Name: "Tom", Age: 18},
		{Name: "Jerry", Age: 20},
	}
	personsMap := slice.ToMapV(persons, func(person Person) (string, int) {
		return person.Name, person.Age
	})
	fmt.Println("persons to map:", personsMap)
}

func ExampleSliceContains() {
	fmt.Println("ExampleSliceContains:")

	// Contains
	intSlice := []int{1, 2, 3, 4, 5}
	target := 3

	// 使用 Contains 检查 target 是否存在于 intSlice 中
	if slice.Contains(intSlice, target) {
		fmt.Printf("%d is in the slice.\n", target)
	} else {
		fmt.Printf("%d is not in the slice.\n", target)
	}

	// ContainsAny
	intSlice = []int{1, 2, 3, 4, 5}
	targets := []int{7, 8, 9}

	// 使用 ContainsAny 检查 intSlice 中是否存在 targets 中的任何一个元素
	if slice.ContainsAny(intSlice, targets) {
		fmt.Println("The slice contains any of the target elements.")
	} else {
		fmt.Println("The slice does not contain any of the target elements.")
	}

	// ContainsAll
	intSlice = []int{1, 2, 3, 4, 5}
	targets = []int{2, 4}

	// 使用 ContainsAll 检查 intSlice 是否包含 targets 中的所有元素
	if slice.ContainsAll(intSlice, targets) {
		fmt.Println("The slice contains all of the target elements.")
	} else {
		fmt.Println("The slice does not contain all of the target elements.")
	}
}

func ExampleSliceIntersection() {
	fmt.Println("ExampleSliceIntersection:")

	// Intersection
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	intersection := slice.Intersection[int](slice1, slice2)

	fmt.Println("Common elements in slice1 and slice2: ", intersection)

	// IntersectionFunc
	type Person struct {
		Name string
		Age  int
	}

	persons1 := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Eve", Age: 22},
	}
	persons2 := []Person{
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Eve", Age: 21},
	}

	commonPersons := slice.IntersectionFunc(persons1, persons2, func(left, right Person) bool {
		return left.Name == right.Name && left.Age == right.Age
	})

	for _, person := range commonPersons {
		fmt.Println("Common person: ", person)
	}
}

func ExampleSliceUnion() {
	fmt.Println("ExampleSliceIntersection:")

	// Union
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	unionSet := slice.Union(slice1, slice2)

	fmt.Println("The union of slice1 and slice2 is: ", unionSet)

	// UnionFunc
	type Person struct {
		Name string
		Age  int
	}

	persons1 := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Eve", Age: 22},
	}
	persons2 := []Person{
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Eve", Age: 21},
	}

	unionPersons := slice.UnionFunc(persons1, persons2, func(left, right Person) bool {
		return left.Name == right.Name && left.Age == right.Age
	})

	fmt.Println(unionPersons)
}

func ExampleSliceDifference() {
	fmt.Println("ExampleSliceDifference:")

	// Difference
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	differSet := slice.Difference(slice1, slice2)

	fmt.Println("The difference of slice1 and slice2 is: ", differSet)

	// DifferenceFunc
	type Person struct {
		Name string
		Age  int
	}

	persons1 := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Eve", Age: 22},
	}
	persons2 := []Person{
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Eve", Age: 21},
	}

	differPersons := slice.DifferenceFunc(persons1, persons2, func(left, right Person) bool {
		return left.Name == right.Name && left.Age == right.Age
	})

	fmt.Println(differPersons)
}

func ExampleSymmetricDifference() {
	fmt.Println("ExampleSymmetricDifference:")

	// SymmetricDifference
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	symDifferSet := slice.SymmetricDifference(slice1, slice2)

	fmt.Println("The symmetric difference of slice1 and slice2 is: ", symDifferSet)

	// SymmetricDifferenceFunc
	type Person struct {
		Name string
		Age  int
	}

	persons1 := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Eve", Age: 22},
	}
	persons2 := []Person{
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Eve", Age: 21},
	}

	symDifferPersons := slice.SymmetricDifferenceFunc(persons1, persons2, func(left, right Person) bool {
		return left.Name == right.Name && left.Age == right.Age
	})

	fmt.Println(symDifferPersons)
}

func main() {
	//ExampleSliceAdd()
	//ExampleSliceDelete()
	//ExampleSliceDeleteIf()
	//ExampleSliceShrink()
	//ExampleSliceDeleteShrink()
	//ExampleSliceAggregate()
	//ExampleSliceReverse()
	//ExampleSliceFind()
	//ExampleSliceIndex()
	//ExampleSliceMap()
	//ExampleSliceContains()
	//ExampleSliceIntersection()
	//ExampleSliceUnion()
	//ExampleSliceDifference()
	ExampleSymmetricDifference()
}
