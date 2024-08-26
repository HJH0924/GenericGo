// Package slice
/**
* @Project : GenericGo
* @File    : reverse_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/5 18:55
**/

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseInt(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		wantSlice []int
	}{
		{
			name:      "Reverse of positive integers",
			slice:     []int{1, 2, 3, 4, 5},
			wantSlice: []int{5, 4, 3, 2, 1},
		},
		{
			name:      "Reverse of empty slice",
			slice:     []int{},
			wantSlice: []int{},
		},
		{
			name:      "Reverse of nil slice",
			slice:     nil,
			wantSlice: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := Reverse[int](test.slice)
			for i := 0; i < len(res); i++ {
				assert.Equal(t, test.wantSlice[i], res[i])
			}
			assert.NotSame(t, test.wantSlice, res)
		})
	}
}

func TestReverseStruct(t *testing.T) {
	type ExampleStruct struct {
		Key   int
		Value string
	}
	tests := []struct {
		name      string
		slice     []ExampleStruct
		wantSlice []ExampleStruct
	}{
		{
			name:      "Reverse of structs",
			slice:     []ExampleStruct{{1, "hello"}, {2, "world"}},
			wantSlice: []ExampleStruct{{2, "world"}, {1, "hello"}},
		},
		{
			name:      "Reverse of empty slice",
			slice:     []ExampleStruct{},
			wantSlice: []ExampleStruct{},
		},
		{
			name:      "Reverse of nil slice",
			slice:     nil,
			wantSlice: []ExampleStruct{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := Reverse[ExampleStruct](test.slice)
			assert.ElementsMatch(t, test.wantSlice, res)
			assert.NotSame(t, test.wantSlice, res)
		})
	}
}

func TestReverseInPlaceInt(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		wantSlice []int
	}{
		{
			name:      "Reverse of positive integers",
			slice:     []int{1, 2, 3, 4, 5},
			wantSlice: []int{5, 4, 3, 2, 1},
		},
		{
			name:      "Reverse of empty slice",
			slice:     []int{},
			wantSlice: []int{},
		},
		{
			name:      "Reverse of nil slice",
			slice:     nil,
			wantSlice: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ReverseInPlace[int](test.slice)
			for i := 0; i < len(test.slice); i++ {
				assert.Equal(t, test.wantSlice[i], test.slice[i])
			}
		})
	}
}

func TestReverseInPlaceStruct(t *testing.T) {
	type ExampleStruct struct {
		Key   int
		Value string
	}
	tests := []struct {
		name      string
		slice     []ExampleStruct
		wantSlice []ExampleStruct
	}{
		{
			name:      "Reverse of structs",
			slice:     []ExampleStruct{{1, "hello"}, {2, "world"}},
			wantSlice: []ExampleStruct{{2, "world"}, {1, "hello"}},
		},
		{
			name:      "Reverse of empty slice",
			slice:     []ExampleStruct{},
			wantSlice: []ExampleStruct{},
		},
		{
			name:      "Reverse of nil slice",
			slice:     nil,
			wantSlice: []ExampleStruct{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ReverseInPlace[ExampleStruct](test.slice)
			assert.ElementsMatch(t, test.wantSlice, test.slice)
		})
	}
}
