// Package set
/**
* @Project : GenericGo
* @File    : hashset_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/27 16:44
**/

package set

import (
	"testing"

	"github.com/HJH0924/GenericGo/slice"
	"github.com/stretchr/testify/assert"
)

func TestHashSet_AddKeys(t *testing.T) {
	tests := []struct {
		name    string
		addVals []int
		wantMap map[int]struct{}
	}{
		{
			name:    "Test with unique values",
			addVals: []int{1, 2, 3},
			wantMap: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
		{
			name:    "Test with duplicate values",
			addVals: []int{1, 1, 2, 2, 3, 3},
			wantMap: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
		{
			name:    "Test with empty slice",
			addVals: []int{},
			wantMap: map[int]struct{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashSet := NewHashSetWithCap[int](len(tt.addVals))
			hashSet.AddKeys(tt.addVals)
			assert.Equal(t, tt.wantMap, hashSet.m)
		})
	}
}

func TestHashSet_RemoveKeys(t *testing.T) {
	tests := []struct {
		name       string
		removeKeys []int
		wantMap    map[int]struct{}
	}{
		{
			name:       "Test with unique values",
			removeKeys: []int{1, 2, 3},
			wantMap: map[int]struct{}{
				4: {},
				5: {},
			},
		},
		{
			name:       "Test with duplicate values",
			removeKeys: []int{1, 1, 2, 2, 3, 3},
			wantMap: map[int]struct{}{
				4: {},
				5: {},
			},
		},
		{
			name:       "Test with empty slice",
			removeKeys: []int{},
			wantMap: map[int]struct{}{
				1: {},
				2: {},
				3: {},
				4: {},
				5: {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addVals := []int{1, 2, 3, 4, 5}
			hashSet := NewHashSetWithCap[int](len(addVals))
			hashSet.AddKeys(addVals)
			hashSet.RemoveKeys(tt.removeKeys)
			assert.Equal(t, tt.wantMap, hashSet.m)
		})
	}
}

func TestHashSet_ContainsAny(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target []int
		want   bool
	}{
		{
			name:   "Source contains target element",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{1},
			want:   true,
		},
		{
			name:   "Source does not contain target element",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{6},
			want:   false,
		},
		{
			name:   "Source contains multiple target elements",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{1, 2, 3},
			want:   true,
		},
		{
			name:   "Source contains some target elements",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{1, 2, 6},
			want:   true,
		},
		{
			name:   "Target elements not in source",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{6, 7},
			want:   false,
		},
		{
			name:   "Source is empty",
			src:    []int{},
			target: []int{6, 7},
			want:   false,
		},
		{
			name:   "Target is empty",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashSet := NewHashSetWithCap[int](len(tt.src))
			hashSet.AddKeys(tt.src)
			got := hashSet.ContainsAny(tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHashSet_ContainsAll(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target []int
		want   bool
	}{
		{
			name:   "Source contains target element",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{1},
			want:   true,
		},
		{
			name:   "Source does not contain target element",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{6},
			want:   false,
		},
		{
			name:   "Source contains multiple target elements",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{1, 2, 3},
			want:   true,
		},
		{
			name:   "Source contains some target elements",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{1, 2, 6},
			want:   false,
		},
		{
			name:   "Target elements not in source",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{6, 7},
			want:   false,
		},
		{
			name:   "Source is empty",
			src:    []int{},
			target: []int{6, 7},
			want:   false,
		},
		{
			name:   "Target is empty",
			src:    []int{1, 2, 3, 4, 5},
			target: []int{},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashSet := NewHashSetWithCap[int](len(tt.src))
			hashSet.AddKeys(tt.src)
			got := hashSet.ContainsAll(tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHashSet_Size(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want int
	}{
		{
			name: "Test with unique elements",
			src:  []int{1, 2, 3, 4},
			want: 4,
		},
		{
			name: "Test with duplicate elements",
			src:  []int{1, 1, 2, 2, 3, 3},
			want: 3,
		},
		{
			name: "Test with empty slice",
			src:  []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashSet := NewHashSetWithCap[int](len(tt.src))
			hashSet.AddKeys(tt.src)
			got := hashSet.Size()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHashSet_Keys(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []int
	}{
		{
			name: "Test with unique elements",
			src:  []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Test with duplicate elements",
			src:  []int{1, 1, 2, 2, 3, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Test with empty slice",
			src:  []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashSet := NewHashSetWithCap[int](len(tt.src))
			hashSet.AddKeys(tt.src)
			gotKeys := hashSet.Keys()
			isContainsAll := slice.ContainsAll(tt.src, gotKeys)
			assert.True(t, true, isContainsAll)
		})
	}
}
