// Package slice
/**
* @Project : GenericGo
* @File    : aggregate_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/7/31 21:03
**/

package slice

import (
	"testing"

	genericgo "github.com/HJH0924/GenericGo"
	"github.com/HJH0924/GenericGo/errs"
	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		wantRes int
		wantErr error
	}{
		{
			name:    "Empty slice",
			slice:   []int{},
			wantRes: 0,
			wantErr: errs.NewErrEmptySlice(),
		},
		{
			name:    "Single value",
			slice:   []int{1},
			wantRes: 1,
		},
		{
			name:    "Multiple values",
			slice:   []int{1, 2, 3},
			wantRes: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRes, gotErr := Max(tc.slice)
			if gotErr != nil {
				assert.Equal(t, tc.wantErr, gotErr)
			} else {
				assert.Equal(t, tc.wantRes, gotRes)
			}
		})
	}

	testMaxTypes[uint](t)
	testMaxTypes[uint8](t)
	testMaxTypes[uint16](t)
	testMaxTypes[uint32](t)
	testMaxTypes[uint64](t)
	testMaxTypes[int](t)
	testMaxTypes[int8](t)
	testMaxTypes[int16](t)
	testMaxTypes[int32](t)
	testMaxTypes[int64](t)
	testMaxTypes[float32](t)
	testMaxTypes[float64](t)
}

func testMaxTypes[T genericgo.RealNumber](t *testing.T) {
	res, _ := Max[T]([]T{1, 2, 3})
	assert.Equal(t, T(3), res)
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		wantRes int
		wantErr error
	}{
		{
			name:    "Empty slice",
			slice:   []int{},
			wantRes: 0,
			wantErr: errs.NewErrEmptySlice(),
		},
		{
			name:    "Single value",
			slice:   []int{1},
			wantRes: 1,
		},
		{
			name:    "Multiple values",
			slice:   []int{1, 2, 3},
			wantRes: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRes, gotErr := Min(tc.slice)
			if gotErr != nil {
				assert.Equal(t, tc.wantErr, gotErr)
			} else {
				assert.Equal(t, tc.wantRes, gotRes)
			}
		})
	}

	testMinTypes[uint](t)
	testMinTypes[uint8](t)
	testMinTypes[uint16](t)
	testMinTypes[uint32](t)
	testMinTypes[uint64](t)
	testMinTypes[int](t)
	testMinTypes[int8](t)
	testMinTypes[int16](t)
	testMinTypes[int32](t)
	testMinTypes[int64](t)
	testMinTypes[float32](t)
	testMinTypes[float64](t)
}

func testMinTypes[T genericgo.RealNumber](t *testing.T) {
	res, _ := Min[T]([]T{1, 2, 3})
	assert.Equal(t, T(1), res)
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name    string
		slice   []int
		wantRes int
		wantErr error
	}{
		{
			name:    "Empty slice",
			slice:   []int{},
			wantRes: 0,
			wantErr: errs.NewErrEmptySlice(),
		},
		{
			name:    "Single value",
			slice:   []int{1},
			wantRes: 1,
		},
		{
			name:    "Multiple values",
			slice:   []int{1, 2, 3},
			wantRes: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRes, gotErr := Sum(tc.slice)
			if gotErr != nil {
				assert.Equal(t, tc.wantErr, gotErr)
			} else {
				assert.Equal(t, tc.wantRes, gotRes)
			}
		})
	}

	testSumTypes[uint](t)
	testSumTypes[uint8](t)
	testSumTypes[uint16](t)
	testSumTypes[uint32](t)
	testSumTypes[uint64](t)
	testSumTypes[int](t)
	testSumTypes[int8](t)
	testSumTypes[int16](t)
	testSumTypes[int32](t)
	testSumTypes[int64](t)
	testSumTypes[float32](t)
	testSumTypes[float64](t)
}

func testSumTypes[T genericgo.RealNumber](t *testing.T) {
	res, _ := Sum[T]([]T{1, 2, 3})
	assert.Equal(t, T(6), res)
}
