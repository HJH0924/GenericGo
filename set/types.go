// Package set
/**
* @Project : GenericGo
* @File    : types.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/27 16:54
**/

package set

// Set 接口定义了一个通用的集合类型，其中元素必须是可比较的
type Set[T comparable] interface {
	// Add 向集合中添加一个元素
	Add(key T)
	// AddKeys 向集合中添加一组元素
	AddKeys(keys []T)
	// Remove 从集合中删除一个元素
	Remove(key T)
	// RemoveKeys 从集合中删除一组元素
	RemoveKeys(keys []T)
	// Contains 检查集合中是否包含某个元素
	Contains(key T) bool
	// ContainsAny 检查集合中是否包含给定切片中的某个元素
	ContainsAny(keys []T) bool
	// ContainsAll 检查集合中是否包含给定切片中的所有元素
	ContainsAll(keys []T) bool
	// Size 返回集合中的元素数量
	Size() int
	// Keys 返回集合中所有的元素
	Keys() []T
}
