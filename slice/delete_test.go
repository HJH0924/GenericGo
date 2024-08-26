// Package slice
/**
* @Project : GenericGo
* @File    : delete_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/27 17:18
**/

package slice

import (
	"testing"

	"github.com/HJH0924/GenericGo/errs"
	"github.com/stretchr/testify/assert"
)

func TestDeleteInt(t *testing.T) {
	testCases := []struct {
		name        string
		slice       []int
		index       int
		wantSlice   []int
		wantElement int
		wantErr     error
	}{
		{
			name:        "Delete at index 0",
			slice:       []int{235, 346},
			index:       0,
			wantSlice:   []int{346},
			wantElement: 235,
		},
		{
			name:        "Delete at index middle",
			slice:       []int{478, 235, 346},
			index:       len([]int{478, 235, 346}) / 2,
			wantSlice:   []int{478, 346},
			wantElement: 235,
		},
		{
			name:        "Delete at last index",
			slice:       []int{478, 867, 235, 346},
			index:       len([]int{478, 867, 235, 346}),
			wantElement: 0,
			wantErr:     errs.NewErrIndexOutOfRange(len([]int{478, 867, 235, 346}), len([]int{478, 867, 235, 346})),
		},
		{
			name:        "Delete at negative index",
			slice:       []int{235, 346, 345},
			index:       -1,
			wantElement: 0,
			wantErr:     errs.NewErrIndexOutOfRange(len([]int{235, 346, 345}), -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotSlice, gotElement, gotErr := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.wantElement, gotElement)
			if gotErr != nil {
				assert.Equal(t, tc.wantErr, gotErr)
			} else {
				assert.Equal(t, tc.wantSlice, gotSlice)
			}
		})
	}
}

func TestDeleteIf(t *testing.T) {
	testCases := []struct {
		name            string
		slice           []int
		deleteCondition func(index int, value int) bool
		wantSlice       []int
	}{
		{
			name:  "Empty Slice",
			slice: []int{},
			deleteCondition: func(index int, value int) bool {
				return false
			},
			wantSlice: []int{},
		},
		{
			name:  "No elements deleted",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return false
			},
			wantSlice: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:  "Delete first element",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return index == 0
			},
			wantSlice: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:  "Delete first two elements",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return index == 0 || index == 1
			},
			wantSlice: []int{2, 3, 4, 5, 6, 7},
		},
		{
			name:  "Delete single middle element",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return index == 3
			},
			wantSlice: []int{0, 1, 2, 4, 5, 6, 7},
		},
		{
			name:  "Delete multiple non-consecutive middle elements",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return index == 2 || index == 4
			},
			wantSlice: []int{0, 1, 3, 5, 6, 7},
		},
		{
			name:  "Delete multiple consecutive middle elements",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return index == 3 || index == 4
			},
			wantSlice: []int{0, 1, 2, 5, 6, 7},
		},
		{
			name:  "Delete multiple middle elements with one at the beginning and consecutive ones after",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return index == 2 || index == 4 || index == 5
			},
			wantSlice: []int{0, 1, 3, 6, 7},
		},
		{
			name:  "Delete multiple middle elements with consecutive ones at the beginning and one at the end",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return index == 2 || index == 3 || index == 5
			},
			wantSlice: []int{0, 1, 4, 6, 7},
		},
		{
			name:  "Delete last two elements",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return index == 6 || index == 7
			},
			wantSlice: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:  "Delete last element",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return index == 7
			},
			wantSlice: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name:  "Delete all elements",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return true
			},
			wantSlice: []int{},
		},
		{
			name:  "Delete even numbers",
			slice: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(index int, value int) bool {
				return value%2 == 0
			},
			wantSlice: []int{1, 3, 5, 7},
		},
		{
			name:  "Delete negative numbers",
			slice: []int{-3, -1, 0, 1, 2, 3, 4},
			deleteCondition: func(index int, value int) bool {
				return value < 0
			},
			wantSlice: []int{0, 1, 2, 3, 4},
		},
		{
			name:  "Delete numbers greater than five",
			slice: []int{0, 5, 10, 3, 8, 5, 6},
			deleteCondition: func(index int, value int) bool {
				return value > 5
			},
			wantSlice: []int{0, 5, 3, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotSlice := DeleteIf(tc.slice, tc.deleteCondition)
			assert.Equal(t, gotSlice, tc.wantSlice)
		})
	}
}
