// Package slice
/**
* @Project : GenericGo
* @File    : add_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/27 15:09
**/

package slice

import (
	"testing"

	"github.com/HJH0924/GenericGo/errs"
	"github.com/stretchr/testify/assert"
)

func TestAddInt(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		index     int
		element   int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "Add at index 0",
			slice:     []int{235, 346},
			index:     0,
			element:   478,
			wantSlice: []int{478, 235, 346},
		},
		{
			name:      "Add at index middle",
			slice:     []int{478, 235, 346},
			index:     len([]int{478, 235, 346}) / 2,
			element:   867,
			wantSlice: []int{478, 867, 235, 346},
		},
		{
			name:      "Add at last index",
			slice:     []int{478, 867, 235, 346},
			index:     len([]int{478, 867, 235, 346}),
			element:   345,
			wantSlice: []int{478, 867, 235, 346, 345},
		},
		{
			name:    "Add at out of range index",
			slice:   []int{235, 346, 345},
			index:   len([]int{235, 346, 345}) + 1,
			element: 678,
			wantErr: errs.NewErrIndexOutOfRange(len([]int{235, 346, 345}), len([]int{235, 346, 345})+1),
		},
		{
			name:    "Add at negative index",
			slice:   []int{235, 346, 345},
			index:   -1,
			element: 678,
			wantErr: errs.NewErrIndexOutOfRange(len([]int{235, 346, 345}), -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotSlice, gotErr := Add(tc.slice, tc.index, tc.element)
			if gotErr != nil {
				assert.Equal(t, tc.wantErr, gotErr)
			} else {
				assert.Equal(t, tc.wantSlice, gotSlice)
			}
		})
	}
}
