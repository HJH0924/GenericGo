// Package slice
/**
* @Project : GenericGo
* @File    : contains.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/10 10:37
**/

package slice

// Contains 检查目标元素是否包含在源切片中。
func Contains[T comparable](src []T, target T) bool {
	return ContainsFunc[T](src, func(t T) bool {
		return t == target
	})
}

// ContainsFunc 检查源切片中是否包含满足匹配条件的元素。
// match 函数定义了元素的匹配条件。
// 优先使用 Contains
func ContainsFunc[T any](src []T, match matchFunc[T]) bool {
	for _, value := range src {
		if match(value) {
			return true
		}
	}
	return false
}

// ContainsAny 检查源切片中是否包含目标切片中的任何一个元素。
func ContainsAny[T comparable](src, targets []T) bool {
	srcMap := toMap(src)
	for _, target := range targets {
		if _, exists := srcMap[target]; exists {
			return true
		}
	}
	return false
}

// ContainsAnyFunc 检查源切片中是否包含目标切片中的任何一个元素，使用提供的 isEqual 函数来比较元素。
// 优先使用 ContainsAny
func ContainsAnyFunc[T any](src, targets []T, isEqual equalFunc[T]) bool {
	for _, value := range src {
		for _, target := range targets {
			if isEqual(value, target) {
				return true
			}
		}
	}
	return false
}

// ContainsAll 检查源切片是否包含目标切片中的所有元素。
func ContainsAll[T comparable](src, targets []T) bool {
	srcMap := toMap(src)
	for _, target := range targets {
		if _, exists := srcMap[target]; !exists {
			return false
		}
	}
	return true
}

// ContainsAllFunc 检查源切片是否包含目标切片中的所有元素，使用提供的 isEqual 函数来比较元素。
// 优先使用 ContainsAll
func ContainsAllFunc[T comparable](src, targets []T, isEqual equalFunc[T]) bool {
	for _, target := range targets {
		isContains := ContainsFunc(src, func(t T) bool {
			return isEqual(t, target)
		})
		if !isContains {
			return false
		}
	}
	return true
}
