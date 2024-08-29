// Package set
/**
* @Project : GenericGo
* @File    : hashset.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/27 16:44
**/

package set

var (
	// 参考：
	// https://github.com/xxjwxc/uber_go_guide_cn?tab=readme-ov-file#interface-%E5%90%88%E7%90%86%E6%80%A7%E9%AA%8C%E8%AF%81
	_ Set[any] = (*HashSet[any])(nil)
)

// HashSet 基于 map 实现的哈希集合
type HashSet[T comparable] struct {
	m map[T]struct{}
}

// Add 向 HashSet 中添加一个元素
func (Self *HashSet[T]) Add(key T) {
	Self.m[key] = struct{}{}
}

// AddKeys 向 HashSet 中添加一组元素
func (Self *HashSet[T]) AddKeys(keys []T) {
	for _, key := range keys {
		Self.Add(key)
	}
}

// Remove 从 HashSet 中删除一个元素
func (Self *HashSet[T]) Remove(key T) {
	delete(Self.m, key)
}

// RemoveKeys 从 HashSet 中删除一组元素
func (Self *HashSet[T]) RemoveKeys(keys []T) {
	for _, key := range keys {
		Self.Remove(key)
	}
}

// Contains 检查 HashSet 中是否包含某个元素
func (Self *HashSet[T]) Contains(key T) bool {
	_, exists := Self.m[key]
	return exists
}

// ContainsAny 检查 HashSet 中是否包含给定切片中的某个元素
func (Self *HashSet[T]) ContainsAny(keys []T) bool {
	for _, key := range keys {
		if Self.Contains(key) {
			return true
		}
	}
	return false
}

// ContainsAll 检查 HashSet 中是否包含给定切片中的所有元素
func (Self *HashSet[T]) ContainsAll(keys []T) bool {
	for _, key := range keys {
		if !Self.Contains(key) {
			return false
		}
	}
	return true
}

// Size 返回 HashSet 中的元素数量
func (Self *HashSet[T]) Size() int {
	return len(Self.m)
}

// Keys 返回集合中所有的元素
// 返回的顺序不固定
func (Self *HashSet[T]) Keys() []T {
	res := make([]T, 0, len(Self.m))
	for key := range Self.m {
		res = append(res, key)
	}
	return res
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		m: make(map[T]struct{}),
	}
}

// NewHashSetWithCap 创建并返回一个新的HashSet实例，接受一个指定的初始容量
// 优先使用指定容量
func NewHashSetWithCap[T comparable](cap int) *HashSet[T] {
	return &HashSet[T]{
		m: make(map[T]struct{}, cap),
	}
}
