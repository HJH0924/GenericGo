// Package option
/**
* @Project : GenericGo
* @File    : option.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/14 09:40
**/

package option

// Option 定义了一个泛型选项函数，它接受一个指向 T 类型的指针，并可以修改它的属性。
// T 通常是一个结构体，而 Option 用于在创建时对其进行配置。
type Option[T any] func(t *T)

// Apply 将一系列的 Option 应用于给定的 T 类型对象。
// 这个函数会按照提供的顺序依次执行每个选项函数。
func Apply[T any](t *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt(t)
	}
}

// OptionErr 类似于 Option，但它允许选项函数返回一个 error。
// 这可以在配置对象时进行额外的验证。
// 你应该优先使用 Option，除非你在设计 option 模式的时候需要进行一些校验
type OptionErr[T any] func(t *T) error

// ApplyErr 将一系列的 OptionErr 应用于给定的 T 类型对象。
// 如果任何一个选项函数返回 error，ApplyErr 将停止执行并返回该 error。
func ApplyErr[T any](t *T, opts ...OptionErr[T]) error {
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return err
		}
	}
	return nil
}
