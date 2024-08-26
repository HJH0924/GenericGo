// Package slice
/**
* @Project : GenericGo
* @File    : shrink.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/28 22:18
**/

package slice

// calculateNewCapacity 计算基于当前容量和长度的新容量。
// 如果新容量与当前容量不同，将返回计算出的新容量和 true 表示容量已调整；
// 否则返回当前容量和 false 表示容量未调整。
func calculateNewCapacity(currentCapacity int, length int) (int, bool) {
	// 如果当前容量小于或等于 64，通常认为这是一个小切片，
	// 因此不需要调整容量，直接返回当前容量和 false。
	if currentCapacity <= 64 {
		return currentCapacity, false
	}

	// 如果当前容量大于 2048，并且长度小于等于当前容量的 1/2，
	// 则将容量减少到原来的 62.5%。这有助于减少内存占用，
	// 同时仍然保留一定比例的空间以供未来可能的扩展。
	if currentCapacity > 2048 && (length <= (currentCapacity / 2)) {
		factor := 0.625 // 减少到当前容量的 62.5%
		return int(float64(currentCapacity) * factor), true
	}

	// 如果当前容量在 2048 以内，但长度小于等于当前容量的 1/4，
	// 则将容量减半。这种调整较为保守，适用于容量较大但未使用空间比例很高的情况。
	if currentCapacity <= 2048 && (length <= (currentCapacity / 4)) {
		return currentCapacity / 2, true
	}

	// 如果上述条件都不满足，返回原始容量和 false，表示不需要调整容量。
	return currentCapacity, false
}

// ShrinkSlice 尝试缩减切片的容量以匹配其长度。
// 如果原始切片的容量可以被减少，则返回一个新的，容量更小的切片；
// 否则返回原始切片。
func ShrinkSlice[T any](src []T) []T {
	currentCapacity, length := cap(src), len(src)
	newCapacity, isAdjusted := calculateNewCapacity(currentCapacity, length)
	if !isAdjusted {
		return src
	}

	// 创建新切片，指定长度和缩容后的容量
	newSlice := make([]T, length, newCapacity)
	// 复制原始切片的元素到新切片
	copy(newSlice, src)
	return newSlice
}
