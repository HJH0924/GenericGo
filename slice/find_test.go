// Package slice
/**
* @Project : GenericGo
* @File    : find_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/5 22:08
**/

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		match   matchFunc[int]
		wantRes int
		found   bool
	}{
		{
			name:  "Find first match",
			slice: []int{1, 2, 3},
			match: func(val int) bool {
				return val == 3
			},
			wantRes: 3,
			found:   true,
		},
		{
			name:  "No match found",
			slice: []int{1, 2, 3},
			match: func(val int) bool {
				return val == 4
			},
			wantRes: 0,
			found:   false,
		},
		{
			name:  "Empty slice",
			slice: []int{},
			match: func(val int) bool {
				return val == 4
			},
			wantRes: 0,
			found:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, found := Find[int](test.slice, test.match)
			assert.Equal(t, test.wantRes, res)
			assert.Equal(t, test.found, found)
		})
	}
}

func TestFindAll(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		match   matchFunc[int]
		wantAll []int
		found   bool
	}{
		{
			name:    "Find all matches",
			slice:   []int{1, 2, 3, 4, 2},
			match:   func(val int) bool { return val%2 == 0 },
			wantAll: []int{2, 4, 2},
			found:   true,
		},
		{
			name:    "No matches found",
			slice:   []int{1, 3, 5},
			match:   func(val int) bool { return val%2 == 0 },
			wantAll: []int{},
			found:   false,
		},
		{
			name:    "Empty slice",
			slice:   []int{},
			match:   func(val int) bool { return val%2 == 0 },
			wantAll: []int{},
			found:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resAll, found := FindAll[int](test.slice, test.match)
			assert.ElementsMatch(t, test.wantAll, resAll)
			assert.Equal(t, test.found, found)
		})
	}
}
