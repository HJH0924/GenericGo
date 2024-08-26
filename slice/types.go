// Package slice
/**
* @Project : GenericGo
* @File    : types.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/5 21:59
**/

package slice

// matchFunc 定义了一个函数类型，用于匹配元素。
// 它接受一个类型为 T 的参数并返回一个布尔值，表示是否匹配成功。
type matchFunc[T any] func(T) bool

// equalFunc 定义了一个函数类型，用于比较两个元素是否相等。
type equalFunc[T any] func(left, right T) bool

// mapFunc 定义了一个函数类型，用于执行映射操作。
// 它接受两个参数：第一个是整数类型的索引，表示元素在源集合中的位置；
// 第二个是类型为 Source 的参数，表示元素本身。
// mapFunc 返回一个类型为 Result 的值，这个值是源元素经过某些处理或转换后的结果。
type mapFunc[Source any, Result any] func(int, Source) Result

// KeyExtractor 定义了一个函数类型，它接收一个元素并返回一个用作映射键的值
type KeyExtractor[Element any, Key comparable] func(Element) Key

// KeyValueMapper 定义了一个函数类型，它接收一个元素并返回一个键值对
type KeyValueMapper[Element any, Key comparable, Value any] func(Element) (Key, Value)
