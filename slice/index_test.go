// Package slice
/**
* @Project : GenericGo
* @File    : index_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/9 15:19
**/

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target int
		want   int
	}{
		{
			name:   "First occurrence of target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "First element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "Last element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 5,
			want:   5,
		},
		{
			name:   "Empty slice",
			src:    []int{},
			target: 3,
			want:   -1,
		},
		{
			name:   "Nil slice",
			target: 3,
			want:   -1,
		},
		{
			name:   "Target not in slice",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 6,
			want:   -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Index[int](test.src, test.target))
		})
	}
}

func TestIndexFunc(t *testing.T) {
	tests := []struct {
		name  string
		src   []int
		match matchFunc[int]
		want  int
	}{
		{
			name: "First occurrence of target",
			src:  []int{1, 2, 3, 4, 4, 5},
			match: func(i int) bool {
				return i > 3
			},
			want: 3,
		},
		{
			name: "First element is target",
			src:  []int{1, 2, 3, 4, 4, 5},
			match: func(i int) bool {
				return i < 2
			},
			want: 0,
		},
		{
			name: "Last element is target",
			src:  []int{1, 2, 3, 4, 4, 5},
			match: func(i int) bool {
				return i > 4
			},
			want: 5,
		},
		{
			name: "Empty slice",
			src:  []int{},
			match: func(i int) bool {
				return i > 4
			},
			want: -1,
		},
		{
			name: "Nil slice",
			match: func(i int) bool {
				return i > 4
			},
			want: -1,
		},
		{
			name: "Target not in slice",
			src:  []int{1, 2, 3, 3, 4, 5},
			match: func(i int) bool {
				return i > 5
			},
			want: -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, IndexFunc[int](test.src, test.match))
		})
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target int
		want   int
	}{
		{
			name:   "Last occurrence of target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 3,
			want:   3,
		},
		{
			name:   "First element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "Last element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 5,
			want:   5,
		},
		{
			name:   "Empty slice",
			src:    []int{},
			target: 3,
			want:   -1,
		},
		{
			name:   "Nil slice",
			target: 3,
			want:   -1,
		},
		{
			name:   "Target not in slice",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 6,
			want:   -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, LastIndex[int](test.src, test.target))
		})
	}
}

func TestLastIndexFunc(t *testing.T) {
	tests := []struct {
		name  string
		src   []int
		match matchFunc[int]
		want  int
	}{
		{
			name: "Last occurrence of target",
			src:  []int{1, 2, 3, 4, 4, 5},
			match: func(i int) bool {
				return i > 3
			},
			want: 5,
		},
		{
			name: "First element is target",
			src:  []int{1, 2, 3, 4, 4, 5},
			match: func(i int) bool {
				return i < 2
			},
			want: 0,
		},
		{
			name: "Last element is target",
			src:  []int{1, 2, 3, 4, 4, 5},
			match: func(i int) bool {
				return i > 4
			},
			want: 5,
		},
		{
			name: "Empty slice",
			src:  []int{},
			match: func(i int) bool {
				return i > 4
			},
			want: -1,
		},
		{
			name: "Nil slice",
			match: func(i int) bool {
				return i > 4
			},
			want: -1,
		},
		{
			name: "Target not in slice",
			src:  []int{1, 2, 3, 3, 4, 5},
			match: func(i int) bool {
				return i > 5
			},
			want: -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, LastIndexFunc[int](test.src, test.match))
		})
	}
}

func TestIndexAll(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target int
		want   []int
	}{
		{
			name:   "All occurrence of target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 3,
			want:   []int{2, 3},
		},
		{
			name:   "First element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 1,
			want:   []int{0},
		},
		{
			name:   "Last element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 5,
			want:   []int{5},
		},
		{
			name:   "Empty slice",
			src:    []int{},
			target: 3,
			want:   []int{},
		},
		{
			name:   "Nil slice",
			target: 3,
			want:   []int{},
		},
		{
			name:   "Target not in slice",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 6,
			want:   []int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, IndexAll[int](test.src, test.target))
		})
	}
}

func TestIndexAllFunc(t *testing.T) {
	tests := []struct {
		name  string
		src   []int
		match matchFunc[int]
		want  []int
	}{
		{
			name: "All occurrence of target",
			src:  []int{1, 2, 3, 3, 4, 5},
			match: func(i int) bool {
				return i > 3
			},
			want: []int{4, 5},
		},
		{
			name: "First element is target",
			src:  []int{1, 2, 3, 3, 4, 5},
			match: func(i int) bool {
				return i < 2
			},
			want: []int{0},
		},
		{
			name: "Last element is target",
			src:  []int{1, 2, 3, 3, 4, 5},
			match: func(i int) bool {
				return i > 4
			},
			want: []int{5},
		},
		{
			name: "Empty slice",
			src:  []int{},
			match: func(i int) bool {
				return i > 3
			},
			want: []int{},
		},
		{
			name: "Nil slice",
			match: func(i int) bool {
				return i > 3
			},
			want: []int{},
		},
		{
			name: "Target not in slice",
			src:  []int{1, 2, 3, 3, 4, 5},
			match: func(i int) bool {
				return i > 5
			},
			want: []int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, IndexAllFunc[int](test.src, test.match))
		})
	}
}
