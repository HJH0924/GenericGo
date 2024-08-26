// Package slice
/**
* @Project : GenericGo
* @File    : map_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/9 16:54
**/

package slice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		mapping mapFunc[int, string]
		want    []string
	}{
		{
			name:  "Map integers to their string representation",
			slice: []int{1, 2, 3},
			mapping: func(index int, source int) string {
				return strconv.Itoa(source)
			},
			want: []string{"1", "2", "3"},
		},
		{
			name:    "Map an empty slice",
			slice:   []int{},
			mapping: func(index int, source int) string { return strconv.Itoa(source) },
			want:    []string{},
		},
		{
			name:    "Map an nil slice",
			mapping: func(index int, source int) string { return strconv.Itoa(source) },
			want:    []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Map[int, string](test.slice, test.mapping))
		})
	}
}

func TestFilterAndMap(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		match   matchFunc[int]
		mapping mapFunc[int, string]
		want    []string
	}{
		{
			name:  "",
			slice: []int{1, 2, -2, 3},
			match: func(i int) bool {
				return i > 0
			},
			mapping: func(index int, source int) string {
				return strconv.Itoa(source)
			},
			want: []string{"1", "2", "3"},
		},
		{
			name:  "Map an empty slice",
			slice: []int{},
			match: func(i int) bool {
				return i > 0
			},
			mapping: func(index int, source int) string { return strconv.Itoa(source) },
			want:    []string{},
		},
		{
			name: "Map an nil slice",
			match: func(i int) bool {
				return i > 0
			},
			mapping: func(index int, source int) string { return strconv.Itoa(source) },
			want:    []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, FilterAndMap[int, string](test.slice, test.match, test.mapping))
		})
	}
}

func TestToMap(t *testing.T) {
	t.Run("integer strings to map[int]string", func(t *testing.T) {
		elements := []string{"1", "2", "3", "4", "5"}
		resMap := ToMap(elements, func(element string) int {
			elementInt, _ := strconv.Atoi(element)
			return elementInt
		})
		expectedMap := map[int]string{
			1: "1",
			2: "2",
			3: "3",
			4: "4",
			5: "5",
		}
		assert.Equal(t, expectedMap, resMap)
	})

	t.Run("struct to map with string key", func(t *testing.T) {
		type User struct {
			name  string
			phone string
			id    int
		}
		users := []User{
			{name: "Tom", phone: "189xxxxxxxx", id: 1},
			{name: "Jerry", phone: "136xxxxxxxx", id: 2},
		}
		resMap := ToMap(users, func(user User) string {
			return user.name
		})
		expectedMap := map[string]User{
			"Tom":   {name: "Tom", phone: "189xxxxxxxx", id: 1},
			"Jerry": {name: "Jerry", phone: "136xxxxxxxx", id: 2},
		}
		assert.Equal(t, expectedMap, resMap)
	})

	t.Run("struct to map with duplicate keys", func(t *testing.T) {
		type User struct {
			name  string
			phone string
			id    int
		}
		users := []User{
			{name: "Tom", phone: "189xxxxxxxx", id: 1},
			{name: "Jerry", phone: "136xxxxxxxx", id: 2},
			{name: "Tom", phone: "136xxxxxxxx", id: 3}, // 故意添加重复的键
		}
		resMap := ToMap(users, func(user User) string {
			return user.name
		})
		// 预期结果应只包含最后出现的元素，因为映射的键是唯一的
		expectedMap := map[string]User{
			"Tom":   {name: "Tom", phone: "136xxxxxxxx", id: 3},
			"Jerry": {name: "Jerry", phone: "136xxxxxxxx", id: 2},
		}
		assert.Equal(t, expectedMap, resMap)
	})

	t.Run("passing nil slice, expect empty map", func(t *testing.T) {
		var elements []string = nil
		resMap := ToMap(elements, func(element string) int {
			elementInt, _ := strconv.Atoi(element)
			return elementInt
		})
		expectedMap := make(map[int]string)
		assert.Equal(t, expectedMap, resMap)
	})
}

func TestToMapV(t *testing.T) {
	t.Run("integer strings to map[int]int", func(t *testing.T) {
		elements := []string{"1", "2", "3", "4", "5"}
		resMap := ToMapV(elements, func(element string) (int, int) {
			elementInt, _ := strconv.Atoi(element)
			return elementInt, elementInt
		})
		expectedMap := map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 4,
			5: 5,
		}
		assert.Equal(t, expectedMap, resMap)
	})

	t.Run("struct to map with string key", func(t *testing.T) {
		type User struct {
			name  string
			phone string
			id    int
		}
		users := []User{
			{name: "Tom", phone: "189xxxxxxxx", id: 1},
			{name: "Jerry", phone: "136xxxxxxxx", id: 2},
		}
		resMap := ToMapV(users, func(user User) (int, string) {
			return user.id, user.name
		})
		expectedMap := map[int]string{
			1: "Tom",
			2: "Jerry",
		}
		assert.Equal(t, expectedMap, resMap)
	})

	t.Run("struct to map with duplicate keys", func(t *testing.T) {
		type User struct {
			name  string
			phone string
			id    int
		}
		users := []User{
			{name: "Tom", phone: "189xxxxxxxx", id: 1},
			{name: "Jerry", phone: "136xxxxxxxx", id: 2},
			{name: "Tom", phone: "136xxxxxxxx", id: 3}, // 故意添加重复的键
		}
		resMap := ToMapV(users, func(user User) (string, string) {
			return user.name, user.phone
		})
		// 预期结果应只包含最后出现的元素，因为映射的键是唯一的
		expectedMap := map[string]string{
			"Tom":   "136xxxxxxxx",
			"Jerry": "136xxxxxxxx",
		}
		assert.Equal(t, expectedMap, resMap)
	})

	t.Run("passing nil slice, expect empty map", func(t *testing.T) {
		var elements []string = nil
		resMap := ToMap(elements, func(element string) int {
			elementInt, _ := strconv.Atoi(element)
			return elementInt
		})
		expectedMap := make(map[int]string)
		assert.Equal(t, expectedMap, resMap)
	})
}
