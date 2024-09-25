// Package list
/**
* @Project : GenericGo
* @File    : linked_list_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/16 10:38
**/

package list

import (
	"fmt"
	"testing"

	"errors"
	"github.com/HJH0924/GenericGo/errs"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Append(t *testing.T) {
	tests := []struct {
		name      string
		list      *LinkedList[int]
		inputVals []int
		wantSlice []int
	}{
		{
			name:      "Append non-empty values to non-empty list",
			list:      NewLinkedListOf([]int{1, 2, 3}),
			inputVals: []int{4, 5, 6},
			wantSlice: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:      "Append empty values to non-empty list",
			list:      NewLinkedListOf([]int{1, 2, 3}),
			inputVals: []int{},
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "Append nil to non-empty list",
			list:      NewLinkedListOf([]int{1, 2, 3}),
			inputVals: nil,
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "Append non-empty values to empty list",
			list:      NewLinkedListOf([]int{}),
			inputVals: []int{1, 2, 3},
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "Append empty values to empty list",
			list:      NewLinkedListOf([]int{}),
			inputVals: []int{},
			wantSlice: []int{},
		},
		{
			name:      "Append nil to empty list",
			list:      NewLinkedListOf([]int{}),
			inputVals: nil,
			wantSlice: []int{},
		},
		{
			name:      "Append non-empty values to nil list",
			list:      NewLinkedListOf[int](nil),
			inputVals: []int{1, 2, 3},
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "Append empty values to nil list",
			list:      NewLinkedListOf[int](nil),
			inputVals: []int{},
			wantSlice: []int{},
		},
		{
			name:      "Append nil to nil list",
			list:      NewLinkedListOf[int](nil),
			inputVals: nil,
			wantSlice: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.list.Append(test.inputVals...)
			assert.Equal(t, test.wantSlice, test.list.AsSlice())
		})
	}
}

func TestLinkedList_Add(t *testing.T) {
	tests := []struct {
		name      string
		list      *LinkedList[int]
		index     int
		value     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "Add at index 0",
			list:      NewLinkedListOf([]int{235, 346}),
			index:     0,
			value:     478,
			wantSlice: []int{478, 235, 346},
		},
		{
			name:      "Add at index middle",
			list:      NewLinkedListOf([]int{478, 235, 346}),
			index:     len([]int{478, 235, 346}) / 2,
			value:     867,
			wantSlice: []int{478, 867, 235, 346},
		},
		{
			name:      "Add at last index",
			list:      NewLinkedListOf([]int{478, 867, 235, 346}),
			index:     len([]int{478, 867, 235, 346}),
			value:     345,
			wantSlice: []int{478, 867, 235, 346, 345},
		},
		{
			name:    "Add at out of range index",
			list:    NewLinkedListOf([]int{235, 346, 345}),
			index:   len([]int{235, 346, 345}) + 1,
			value:   678,
			wantErr: errs.NewErrIndexOutOfRange(len([]int{235, 346, 345}), len([]int{235, 346, 345})+1),
		},
		{
			name:    "Add at negative index",
			list:    NewLinkedListOf([]int{235, 346, 345}),
			index:   -1,
			value:   678,
			wantErr: errs.NewErrIndexOutOfRange(len([]int{235, 346, 345}), -1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotErr := test.list.Add(test.index, test.value)
			if gotErr != nil {
				assert.Equal(t, test.wantErr, gotErr)
			} else {
				assert.Equal(t, test.wantSlice, test.list.AsSlice())
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	tests := []struct {
		name        string
		list        *LinkedList[int]
		index       int
		wantSlice   []int
		wantElement int
		wantErr     error
	}{

		{
			name:        "Delete at index 0",
			list:        NewLinkedListOf([]int{235, 346}),
			index:       0,
			wantSlice:   []int{346},
			wantElement: 235,
		},
		{
			name:        "Delete at index middle",
			list:        NewLinkedListOf([]int{478, 235, 346}),
			index:       len([]int{478, 235, 346}) / 2,
			wantSlice:   []int{478, 346},
			wantElement: 235,
		},
		{
			name:        "Delete at last index",
			list:        NewLinkedListOf([]int{478, 867, 235, 346}),
			index:       len([]int{478, 867, 235, 346}),
			wantElement: 0,
			wantErr:     errs.NewErrIndexOutOfRange(len([]int{478, 867, 235, 346}), len([]int{478, 867, 235, 346})),
		},
		{
			name:        "Delete at negative index",
			list:        NewLinkedListOf([]int{235, 346, 345}),
			index:       -1,
			wantElement: 0,
			wantErr:     errs.NewErrIndexOutOfRange(len([]int{235, 346, 345}), -1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotElement, gotErr := test.list.Delete(test.index)
			assert.Equal(t, test.wantElement, gotElement)
			if gotErr != nil {
				assert.Equal(t, test.wantErr, gotErr)
			} else {
				assert.Equal(t, test.wantElement, gotElement)
			}
		})
	}
}

func TestLinkedList_Set(t *testing.T) {
	tests := []struct {
		name     string
		list     *LinkedList[int]
		index    int
		value    int
		expected []int
		wantErr  error
	}{
		{
			name:     "Set element at the beginning of the list",
			list:     NewLinkedListOf([]int{1, 2, 3, 4}),
			index:    0,
			value:    10,
			expected: []int{10, 2, 3, 4},
		},
		{
			name:     "Set element at the middle of the list",
			list:     NewLinkedListOf([]int{1, 2, 3, 4}),
			index:    2,
			value:    30,
			expected: []int{1, 2, 30, 4},
		},
		{
			name:     "Set element at the end of the list",
			list:     NewLinkedListOf([]int{1, 2, 3}),
			index:    2,
			value:    3,
			expected: []int{1, 2, 3},
		},
		{
			name:    "Set element beyond the end of the list",
			list:    NewLinkedListOf([]int{1, 2}),
			index:   10,
			value:   3,
			wantErr: errs.NewErrIndexOutOfRange(len([]int{1, 2}), 10),
		},
		{
			name:    "Set element beyond the end of the list (index == length)",
			list:    NewLinkedListOf([]int{1, 2}),
			index:   2,
			value:   3,
			wantErr: errs.NewErrIndexOutOfRange(len([]int{1, 2}), 2),
		},
		{
			name:    "Set element in an empty list",
			list:    NewLinkedListOf([]int{}),
			index:   0,
			value:   1,
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
		},
		{
			name:     "Set element with negative index",
			list:     NewLinkedListOf([]int{1, 2}),
			index:    -1,
			value:    0,
			expected: []int{1, 2},
			wantErr:  errs.NewErrIndexOutOfRange(len([]int{1, 2}), -1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.list.Set(test.index, test.value)
			if err != nil {
				assert.Equal(t, test.wantErr, err)
			} else {
				assert.Equal(t, test.expected, test.list.AsSlice())
			}
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	tests := []struct {
		name     string
		list     *LinkedList[int]
		index    int
		expected int
		wantErr  error
	}{
		{
			name:     "Get element at the beginning of the list",
			list:     NewLinkedListOf([]int{1, 2, 3, 4}),
			index:    0,
			expected: 1,
		},
		{
			name:     "Get element at the middle of the list",
			list:     NewLinkedListOf([]int{1, 2, 3, 4}),
			index:    2,
			expected: 3,
		},
		{
			name:     "Get element at the end of the list",
			list:     NewLinkedListOf([]int{1, 2, 3}),
			index:    2,
			expected: 3,
		},
		{
			name:     "Get element beyond the end of the list",
			list:     NewLinkedListOf([]int{1, 2}),
			index:    10,
			expected: 3,
			wantErr:  errs.NewErrIndexOutOfRange(len([]int{1, 2}), 10),
		},
		{
			name:     "Get element beyond the end of the list (index == length)",
			list:     NewLinkedListOf([]int{1, 2}),
			index:    2,
			expected: 3,
			wantErr:  errs.NewErrIndexOutOfRange(len([]int{1, 2}), 2),
		},
		{
			name:     "Get element in an empty list",
			list:     NewLinkedListOf([]int{}),
			index:    0,
			expected: 1,
			wantErr:  errs.NewErrIndexOutOfRange(0, 0),
		},
		{
			name:    "Get element with negative index",
			list:    NewLinkedListOf([]int{1, 2}),
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(len([]int{1, 2}), -1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotVal, err := test.list.Get(test.index)
			if err != nil {
				assert.Equal(t, test.wantErr, err)
			} else {
				assert.Equal(t, test.expected, gotVal)
			}
		})
	}
}

func TestLinkedList_Len(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		wantLen int
	}{
		{
			name:    "empty list",
			slice:   []int{},
			wantLen: 0,
		},
		{
			name:    "list with one element",
			slice:   []int{1},
			wantLen: 1,
		},
		{
			name:    "list with multiple elements",
			slice:   []int{1, 2, 3, 4},
			wantLen: 4,
		},
		{
			name:    "nil list",
			slice:   nil,
			wantLen: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := NewLinkedListOf(test.slice)
			assert.Equal(t, test.wantLen, list.Len())
		})
	}
}

func TestLinkedList_Cap(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		wantCap int
	}{
		{
			name:    "empty list",
			slice:   []int{},
			wantCap: 0,
		},
		{
			name:    "list with one element",
			slice:   []int{1},
			wantCap: 1,
		},
		{
			name:    "list with multiple elements",
			slice:   []int{1, 2, 3, 4},
			wantCap: 4,
		},
		{
			name:    "nil list",
			slice:   nil,
			wantCap: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := NewLinkedListOf(test.slice)
			assert.Equal(t, test.wantCap, list.Cap())
		})
	}
}

func TestLinkedList_Range(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected int
		wantErr  error
	}{
		{
			name:     "Calculate the sum of all elements",
			slice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 55,
		},
		{
			name:    "Test interruption",
			slice:   []int{1, 2, 3, 4, -5, 6, 7, -8, 9, 10},
			wantErr: errors.New("index 4 is error"),
		},
		{
			name:    "Test array is nil",
			slice:   nil,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			called := false // 用于检查回调函数是否被调用
			sum := 0
			al := NewArrayListOf(test.slice)
			err := al.Range(func(idx int, val int) error {
				called = true
				if val < 0 {
					return fmt.Errorf("index %d is error", idx)
				}
				sum += val
				return nil
			})

			if test.slice != nil {
				// nil 切片不会调用回调函数，并且返回nil
				// 检查回调是否被调用
				assert.True(t, called, "callback function was not called")
			}

			if err != nil {
				assert.Equal(t, test.wantErr, err)
			} else {
				assert.Equal(t, test.expected, sum)
			}
		})
	}
}

func TestLinkedList_AsSlice(t *testing.T) {
	tests := []struct {
		name      string
		vals      []int
		wantSlice []int
	}{
		{
			name:      "convert non-empty list to slice",
			vals:      []int{1, 2, 3, 4, 5},
			wantSlice: []int{1, 2, 3, 4, 5},
		},
		{
			name:      "convert empty list to slice",
			vals:      []int{},
			wantSlice: []int{},
		},
		{
			name:      "convert nil list to slice",
			vals:      nil,
			wantSlice: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			al := NewArrayListOf(test.vals)
			slice := al.AsSlice()
			// 判断内容是否相同
			assert.Equal(t, test.wantSlice, slice)
			// 因为返回的是一个新切片，所以两个切片的地址必定不同
			// 但是如果 vals 为空，则返回的切片与 vals 共享底层数组
			if len(test.vals) > 0 {
				valsAddr := fmt.Sprintf("%p", test.vals)
				sliceAddr := fmt.Sprintf("%p", slice)
				assert.NotEqual(t, valsAddr, sliceAddr)
			}
		})
	}
}
