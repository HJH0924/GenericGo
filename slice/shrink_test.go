// Package slice
/**
* @Project : GenericGo
* @File    : shrink_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/28 22:31
**/

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShrinkSlice(t *testing.T) {
	testCases := []struct {
		name             string
		originalCapacity int
		elementsToAdd    int
		wantCapacity     int
	}{
		{
			name:             "Small capacity slice, current capacity less than 64",
			originalCapacity: 32,
			elementsToAdd:    10,
			wantCapacity:     32,
		},
		{
			name:             "Current Capacity less than 2048, length <= (current capacity / 4)",
			originalCapacity: 1000,
			elementsToAdd:    20,
			wantCapacity:     500,
		},
		{
			name:             "Current Capacity less than 2048, length > (current capacity / 4)",
			originalCapacity: 1000,
			elementsToAdd:    500,
			wantCapacity:     1000,
		},
		{
			name:             "Current Capacity large than 2048, length > (current capacity / 2)",
			originalCapacity: 3000,
			elementsToAdd:    2000,
			wantCapacity:     3000,
		},
		{
			name:             "Current Capacity large than 2048, length <= (current capacity / 2)",
			originalCapacity: 3000,
			elementsToAdd:    1000,
			wantCapacity:     1875,
		},
		{
			name:             "Zero length slice",
			originalCapacity: 0,
			elementsToAdd:    0,
			wantCapacity:     0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice := make([]int, 0, tc.originalCapacity)

			for i := 0; i < tc.elementsToAdd; i++ {
				slice = append(slice, i)
			}

			sliceAfterShink := ShrinkSlice[int](slice)

			assert.Equal(t, tc.wantCapacity, cap(sliceAfterShink))
		})
	}
}

func TestDeleteShrink(t *testing.T) {
	testCases := []struct {
		name             string
		originalCapacity int
		elementsToDelete int
		wantCapacity     int
	}{
		{
			name:             "Current Capacity less than 2048, length <= (current capacity / 4)",
			originalCapacity: 1000,
			elementsToDelete: 900,
			wantCapacity:     250,
		},
		{
			name:             "Current Capacity less than 2048, length > (current capacity / 4)",
			originalCapacity: 1000,
			elementsToDelete: 500,
			wantCapacity:     1000,
		},
		{
			name:             "Current Capacity large than 2048, length > (current capacity / 2)",
			originalCapacity: 3000,
			elementsToDelete: 1000,
			wantCapacity:     3000,
		},
		{
			name:             "Current Capacity large than 2048, length <= (current capacity / 2)",
			originalCapacity: 3000,
			elementsToDelete: 2000,
			wantCapacity:     1875,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice := make([]int, 0, tc.originalCapacity)
			elementsToAdd := tc.originalCapacity
			for i := 0; i < elementsToAdd; i++ {
				slice = append(slice, i)
			}
			for i := 0; i < tc.elementsToDelete; i++ {
				slice, _, _ = Delete(slice, 0)
			}
			assert.Equal(t, tc.wantCapacity, cap(slice))
		})
	}
}
