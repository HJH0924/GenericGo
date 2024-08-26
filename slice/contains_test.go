// Package slice
/**
* @Project : GenericGo
* @File    : contains_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/10 11:07
**/

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target int
		want   bool
	}{
		{
			name:   "First occurrence of target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 3,
			want:   true,
		},
		{
			name:   "First element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 1,
			want:   true,
		},
		{
			name:   "Last element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 5,
			want:   true,
		},
		{
			name:   "Empty slice",
			src:    []int{},
			target: 3,
			want:   false,
		},
		{
			name:   "Nil slice",
			target: 3,
			want:   false,
		},
		{
			name:   "Target not in slice",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: 6,
			want:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Contains[int](test.src, test.target))
		})
	}
}

func TestContainsFunc(t *testing.T) {
	tests := []struct {
		name  string
		src   []int
		match matchFunc[int]
		want  bool
	}{
		{
			name: "First occurrence of target",
			src:  []int{1, 2, 3, 4, 4, 5},
			match: func(i int) bool {
				return i > 3
			},
			want: true,
		},
		{
			name: "First element is target",
			src:  []int{1, 2, 3, 4, 4, 5},
			match: func(i int) bool {
				return i < 2
			},
			want: true,
		},
		{
			name: "Last element is target",
			src:  []int{1, 2, 3, 4, 4, 5},
			match: func(i int) bool {
				return i > 4
			},
			want: true,
		},
		{
			name: "Empty slice",
			src:  []int{},
			match: func(i int) bool {
				return i > 4
			},
			want: false,
		},
		{
			name: "Nil slice",
			match: func(i int) bool {
				return i > 4
			},
			want: false,
		},
		{
			name: "Target not in slice",
			src:  []int{1, 2, 3, 3, 4, 5},
			match: func(i int) bool {
				return i > 5
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsFunc[int](test.src, test.match))
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target []int
		want   bool
	}{
		{
			name:   "All occurrence of target",
			src:    []int{1, 2},
			target: []int{1, 2, 3, 3, 4},
			want:   true,
		},
		{
			name:   "First element is target",
			src:    []int{1},
			target: []int{1, 2, 3, 3, 4, 5},
			want:   true,
		},
		{
			name:   "Last element is target",
			src:    []int{5},
			target: []int{1, 2, 3, 3, 4, 5},
			want:   true,
		},
		{
			name:   "Empty slice",
			src:    []int{},
			target: []int{},
			want:   false,
		},
		{
			name:   "Nil slice",
			target: []int{},
			want:   false,
		},
		{
			name:   "Target not in slice",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: []int{6, 7},
			want:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAny[int](test.src, test.target))
		})
	}
}

func TestContainsAnyFunc(t *testing.T) {
	tests := []struct {
		name    string
		src     []int
		targets []int
		want    bool
	}{
		{
			name:    "All occurrence of target",
			src:     []int{1, 2, 3, 3, 4, 5},
			targets: []int{2, 4, 5},
			want:    true,
		},
		{
			name:    "First element is target",
			src:     []int{1, 2, 3, 3, 4, 5},
			targets: []int{1},
			want:    true,
		},
		{
			name:    "Last element is target",
			src:     []int{1, 2, 3, 3, 4, 5},
			targets: []int{5},
			want:    true,
		},
		{
			name:    "Empty slice",
			src:     []int{},
			targets: []int{},
			want:    false,
		},
		{
			name:    "Nil slice",
			targets: []int{},
			want:    false,
		},
		{
			name:    "Target not in slice",
			src:     []int{1, 2, 3, 3, 4, 5},
			targets: []int{6, 7},
			want:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAnyFunc[int](test.src, test.targets, func(left, right int) bool {
				return left == right
			}))
		})
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target []int
		want   bool
	}{
		{
			name:   "All occurrence of target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: []int{1, 2, 3, 3, 4},
			want:   true,
		},
		{
			name:   "First element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: []int{1, 2},
			want:   true,
		},
		{
			name:   "Last element is target",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: []int{5, 4},
			want:   true,
		},
		{
			name:   "Empty slice",
			src:    []int{},
			target: []int{},
			want:   true,
		},
		{
			name:   "Nil slice",
			target: []int{},
			want:   true,
		},
		{
			name:   "Target not in slice",
			src:    []int{1, 2, 3, 3, 4, 5},
			target: []int{6, 7},
			want:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAll[int](test.src, test.target))
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	tests := []struct {
		name    string
		src     []int
		targets []int
		want    bool
	}{
		{
			name:    "All occurrence of target",
			src:     []int{1, 2, 3, 3, 4, 5},
			targets: []int{2, 4, 5},
			want:    true,
		},
		{
			name:    "First element is target",
			src:     []int{1, 2, 3, 3, 4, 5},
			targets: []int{1},
			want:    true,
		},
		{
			name:    "Last element is target",
			src:     []int{1, 2, 3, 3, 4, 5},
			targets: []int{5},
			want:    true,
		},
		{
			name:    "Empty slice",
			src:     []int{},
			targets: []int{},
			want:    true,
		},
		{
			name:    "Nil slice",
			targets: []int{},
			want:    true,
		},
		{
			name:    "Target not in slice",
			src:     []int{1, 2, 3, 3, 4, 5},
			targets: []int{6, 7},
			want:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAllFunc[int](test.src, test.targets, func(left, right int) bool {
				return left == right
			}))
		})
	}
}
