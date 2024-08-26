// Package slice
/**
* @Project : GenericGo
* @File    : map.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/9 16:49
**/

package slice

// Map 函数通过提供的映射函数对源切片的每个元素执行操作，并返回结果切片。
// mapping 函数定义了如何将 Source 类型的元素转换为 Result 类型的结果。
// 该函数返回一个新的切片，包含所有转换后的结果。
func Map[Source any, Result any](src []Source, mapping mapFunc[Source, Result]) []Result {
	res := make([]Result, len(src))
	for index, value := range src {
		res[index] = mapping(index, value)
	}
	return res
}

// FilterAndMap 函数对源切片进行过滤和映射操作。
// 仅包含通过 match 函数的元素，并使用 mapping 函数将它们转换为 Result 类型。
// 返回一个包含所有映射结果的新切片。
func FilterAndMap[Source any, Result any](src []Source, match matchFunc[Source], mapping mapFunc[Source, Result]) []Result {
	res := make([]Result, 0, len(src))
	for index, value := range src {
		if match(value) {
			res = append(res, mapping(index, value))
		}
	}
	return res
}

// ToMap 将元素切片转换为映射，其中元素作为值，由KeyExtractor函数生成的键作为映射的键。
// KeyExtractor 函数接收一个元素并返回一个用作映射键的值。
// 返回结果是一个映射，其中包含由KeyExtractor函数生成的键和对应的元素。
// 注意：如果有重复的键，则只包含最后出现的元素，因为映射的键是唯一的。
func ToMap[Element any, Key comparable](elements []Element, keyExtractor KeyExtractor[Element, Key]) map[Key]Element {
	return ToMapV(elements, func(element Element) (Key, Element) {
		return keyExtractor(element), element
	})
}

// ToMapV 将元素切片转换为映射，其中由KeyValueMapper函数生成的键值对存储在映射中。
// KeyValueMapper 函数接收一个元素并返回一个键值对。
// 返回结果是一个映射，其中包含由KeyValueMapper函数生成的键值对。
// 注意：如果有重复的键，则只包含最后出现的元素键值对，因为映射的键是唯一的。
func ToMapV[Element any, Key comparable, Value any](elements []Element, keyValueMapper KeyValueMapper[Element, Key, Value]) map[Key]Value {
	resMap := make(map[Key]Value, len(elements))
	for _, element := range elements {
		key, value := keyValueMapper(element)
		resMap[key] = value
	}
	return resMap
}
