// Package slice
/**
* @Project : GenericGo
* @File    : intersect_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/10 15:15
**/

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		src1     []int
		src2     []int
		expected []int
	}{
		{
			name:     "Common elements in both slices",
			src1:     []int{1, 2, 3, 4},
			src2:     []int{3, 4, 5, 6},
			expected: []int{3, 4},
		},
		{
			name:     "No common elements",
			src1:     []int{1, 2, 3},
			src2:     []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Duplicates in both slices",
			src1:     []int{1, 1, 2, 2},
			src2:     []int{2, 2, 3, 3},
			expected: []int{2},
		},
		{
			name:     "Empty first slice",
			src1:     []int{},
			src2:     []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Empty second slice",
			src1:     []int{1, 2, 3},
			src2:     []int{},
			expected: []int{},
		},
		{
			name:     "First slice is a subset of second",
			src1:     []int{1, 2},
			src2:     []int{2, 3, 4, 5, 1},
			expected: []int{1, 2},
		},
		{
			name:     "Second slice is a subset of first",
			src1:     []int{1, 2, 3, 4, 5},
			src2:     []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Both slices contain the same elements in different order",
			src1:     []int{3, 1, 4, 2},
			src2:     []int{2, 4, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.ElementsMatch(t, test.expected, Intersection(test.src1, test.src2))
		})
	}
}

func TestIntersectionFunc(t *testing.T) {
	tests := []struct {
		name     string
		src1     []int
		src2     []int
		expected []int
	}{
		{
			name:     "Common elements in both slices",
			src1:     []int{1, 2, 3, 4},
			src2:     []int{3, 4, 5, 6},
			expected: []int{3, 4},
		},
		{
			name:     "No common elements",
			src1:     []int{1, 2, 3},
			src2:     []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Duplicates in both slices",
			src1:     []int{1, 1, 2, 2},
			src2:     []int{2, 2, 3, 3},
			expected: []int{2},
		},
		{
			name:     "Empty first slice",
			src1:     []int{},
			src2:     []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Empty second slice",
			src1:     []int{1, 2, 3},
			src2:     []int{},
			expected: []int{},
		},
		{
			name:     "First slice is a subset of second",
			src1:     []int{1, 2},
			src2:     []int{2, 3, 4, 5, 1},
			expected: []int{1, 2},
		},
		{
			name:     "Second slice is a subset of first",
			src1:     []int{1, 2, 3, 4, 5},
			src2:     []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Both slices contain the same elements in different order",
			src1:     []int{3, 1, 4, 2},
			src2:     []int{2, 4, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.ElementsMatch(t, test.expected, IntersectionFunc(test.src1, test.src2, func(left, right int) bool {
				return left == right
			}))
		})
	}
}