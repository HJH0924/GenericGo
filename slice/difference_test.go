// Package slice
/**
* @Project : GenericGo
* @File    : difference_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/10 16:22
**/

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
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
			expected: []int{1, 2},
		},
		{
			name:     "No common elements",
			src1:     []int{1, 2, 3},
			src2:     []int{4, 5, 6},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Duplicates in both slices",
			src1:     []int{1, 1, 2, 2},
			src2:     []int{2, 2, 3, 3},
			expected: []int{1},
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
			expected: []int{1, 2, 3},
		},
		{
			name:     "First slice is a subset of second",
			src1:     []int{1, 2},
			src2:     []int{2, 3, 4, 5, 1},
			expected: []int{},
		},
		{
			name:     "Second slice is a subset of first",
			src1:     []int{1, 2, 3, 4, 5},
			src2:     []int{1, 2},
			expected: []int{3, 4, 5},
		},
		{
			name:     "Both slices contain the same elements in different order",
			src1:     []int{3, 1, 4, 2},
			src2:     []int{2, 4, 1, 3},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.ElementsMatch(t, test.expected, Difference(test.src1, test.src2))
		})
	}
}

func TestDifferenceFunc(t *testing.T) {
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
			expected: []int{1, 2},
		},
		{
			name:     "No common elements",
			src1:     []int{1, 2, 3},
			src2:     []int{4, 5, 6},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Duplicates in both slices",
			src1:     []int{1, 1, 2, 2},
			src2:     []int{2, 2, 3, 3},
			expected: []int{1},
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
			expected: []int{1, 2, 3},
		},
		{
			name:     "First slice is a subset of second",
			src1:     []int{1, 2},
			src2:     []int{2, 3, 4, 5, 1},
			expected: []int{},
		},
		{
			name:     "Second slice is a subset of first",
			src1:     []int{1, 2, 3, 4, 5},
			src2:     []int{1, 2},
			expected: []int{3, 4, 5},
		},
		{
			name:     "Both slices contain the same elements in different order",
			src1:     []int{3, 1, 4, 2},
			src2:     []int{2, 4, 1, 3},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.ElementsMatch(t, test.expected, DifferenceFunc(test.src1, test.src2, func(left, right int) bool {
				return left == right
			}))
		})
	}
}
