// Package tuple
/**
* @Project : GenericGo
* @File    : pair_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/14 11:59
**/

package tuple

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewPair(t *testing.T) {
	tests := []struct {
		name string
		key  any
		val  any
	}{
		{
			name: "key:string, val:int",
			key:  "Tvux",
			val:  23,
		},
		{
			name: "key:int, val:struct",
			key:  1,
			val: user{
				name: "Tvux",
				age:  23,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pair := NewPair(tt.key, tt.val)
			assert.Equal(t, tt.key, pair.Key)
			assert.Equal(t, tt.val, pair.Val)

			assert.Equal(t, reflect.TypeOf(tt.key), reflect.TypeOf(pair.Key))
			assert.Equal(t, reflect.TypeOf(tt.val), reflect.TypeOf(pair.Val))
		})
	}
}

func TestPair_ToString(t *testing.T) {
	tests := []struct {
		name string
		key  any
		val  any
	}{
		{
			name: "key:string, val:int",
			key:  "Tvux",
			val:  23,
		},
		{
			name: "key:int, val:struct",
			key:  1,
			val: user{
				name: "Tvux",
				age:  23,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pair := NewPair(tt.key, tt.val)
			pairStr := pair.ToString()
			fmt.Println(pairStr)
			assert.Equal(t, fmt.Sprintf("<%#v, %#v>", tt.key, tt.val), pairStr)
		})
	}
}

func TestPair_Split(t *testing.T) {
	tests := []struct {
		name string
		key  any
		val  any
	}{
		{
			name: "key:string, val:int",
			key:  "Tvux",
			val:  23,
		},
		{
			name: "key:int, val:struct",
			key:  1,
			val: user{
				name: "Tvux",
				age:  23,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pair := NewPair(tt.key, tt.val)
			pairKey, pairVal := pair.Split()

			assert.Equal(t, tt.key, pairKey)
			assert.Equal(t, tt.val, pairVal)

			assert.Equal(t, reflect.TypeOf(tt.key), reflect.TypeOf(pairKey))
			assert.Equal(t, reflect.TypeOf(tt.val), reflect.TypeOf(pairVal))
		})
	}
}

func TestNewPairs(t *testing.T) {
	tests := []struct {
		name    string
		keys    []any
		vals    []any
		wantErr error
	}{
		{
			name: "ValidPairs",
			keys: []any{"Tom", "Jerry"},
			vals: []any{18, 20},
		},
		{
			name:    "NilKeys",
			keys:    nil,
			vals:    []any{18, 20},
			wantErr: NewErrNilKeysVals(),
		},
		{
			name:    "NilVals",
			keys:    []any{"Tom", "Jerry"},
			vals:    nil,
			wantErr: NewErrNilKeysVals(),
		},
		{
			name:    "NilKeysAndVals",
			keys:    nil,
			vals:    nil,
			wantErr: NewErrNilKeysVals(),
		},
		{
			name: "EmptyKeysAndVals",
			keys: []any{},
			vals: []any{},
		},
		{
			name:    "UnequalLengthsKeysLonger",
			keys:    []any{"Tom", "Jerry", "Tvux"},
			vals:    []any{18, 20},
			wantErr: NewErrUnequalLengthsOfKeysVals(3, 2),
		},
		{
			name:    "UnequalLengthsValsLonger",
			keys:    []any{"Tom", "Jerry"},
			vals:    []any{18, 20, 30},
			wantErr: NewErrUnequalLengthsOfKeysVals(2, 3),
		},
		{
			name: "MixedTypes",
			keys: []any{1, 2.5, "three"},
			vals: []any{"one", "two", "three"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs, err := NewPairs(tt.keys, tt.vals)
			if err != nil {
				assert.Nil(t, pairs)
				assert.Equal(t, tt.wantErr, err)
			} else {
				expectedLen := len(tt.keys)
				assert.Equal(t, expectedLen, len(pairs))
				for i := 0; i < expectedLen; i++ {
					assert.Equal(t, tt.keys[i], pairs[i].Key)
					assert.Equal(t, tt.vals[i], pairs[i].Val)

					assert.Equal(t, reflect.TypeOf(tt.keys[i]), reflect.TypeOf(pairs[i].Key))
					assert.Equal(t, reflect.TypeOf(tt.vals[i]), reflect.TypeOf(pairs[i].Val))

					fmt.Println(pairs[i].ToString())
				}
			}
		})
	}
}

func TestSplitPairs(t *testing.T) {
	tests := []struct {
		name    string
		keys    []any
		vals    []any
		wantErr error
	}{
		{
			name: "ValidPairs",
			keys: []any{"Tom", "Jerry"},
			vals: []any{18, 20},
		},
		{
			name: "StructPairs",
			keys: []any{
				user{name: "Tvux", age: 23},
				user{name: "Tom", age: 18},
				user{name: "Jerry", age: 20},
			},
			vals: []any{"Boss", "Developer", "Tester"},
		},
		{
			name: "EmptyKeysAndVals",
			keys: []any{},
			vals: []any{},
		},
		{
			name: "MixedTypes",
			keys: []any{1, 2.5, "three"},
			vals: []any{"one", "two", "three"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs, err := NewPairs(tt.keys, tt.vals)
			assert.NoError(t, err)

			keys, vals := SplitPairs(pairs)

			expectedLen := len(tt.keys)
			assert.Equal(t, expectedLen, len(keys))
			assert.Equal(t, expectedLen, len(vals))

			for i := 0; i < expectedLen; i++ {
				assert.Equal(t, tt.keys[i], keys[i])
				assert.Equal(t, tt.vals[i], vals[i])

				assert.Equal(t, reflect.TypeOf(tt.keys[i]), reflect.TypeOf(keys[i]))
				assert.Equal(t, reflect.TypeOf(tt.vals[i]), reflect.TypeOf(vals[i]))
			}
		})
	}
}

func TestFlattenPairs(t *testing.T) {
	tests := []struct {
		name    string
		keys    []any
		vals    []any
		wantErr error
	}{
		{
			name: "ValidPairs",
			keys: []any{"Tom", "Jerry"},
			vals: []any{18, 20},
		},
		{
			name: "StructPairs",
			keys: []any{
				user{name: "Tvux", age: 23},
				user{name: "Tom", age: 18},
				user{name: "Jerry", age: 20},
			},
			vals: []any{"Boss", "Developer", "Tester"},
		},
		{
			name: "EmptyKeysAndVals",
			keys: []any{},
			vals: []any{},
		},
		{
			name: "MixedTypes",
			keys: []any{1, 2.5, "three"},
			vals: []any{"one", "two", "three"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs, err := NewPairs(tt.keys, tt.vals)
			assert.NoError(t, err)

			flatPairs := FlattenPairs(pairs)
			n := len(flatPairs)
			assert.Equal(t, len(tt.keys)*2, n)

			for i := 0; i < n; i++ {
				if i%2 == 0 {
					// 0 2 4 ...
					assert.Equal(t, tt.keys[i/2], flatPairs[i])
					assert.Equal(t, reflect.TypeOf(tt.keys[i/2]), reflect.TypeOf(flatPairs[i]))
				} else {
					// 1 3 5 ...
					assert.Equal(t, tt.vals[i/2], flatPairs[i])
					assert.Equal(t, reflect.TypeOf(tt.vals[i/2]), reflect.TypeOf(flatPairs[i]))
				}
			}
		})
	}
}

// 在 Go 语言中，泛型类型参数在运行时并不是动态的，它们需要在编译时就确定。
// 因此，不能在测试用例中直接使用 reflect.TypeOf 来传递类型参数。
// 所以，PackPairs 函数的类型参数也不能在运行时动态决定。
func TestPackPairs_StringInt(t *testing.T) {
	tests := []struct {
		name      string
		flatPairs []any
		wantErr   error
	}{
		{
			name:      "",
			flatPairs: []any{"Tom", 18, "Jerry", 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs, err := PackPairs[string, int](tt.flatPairs)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Equal(t, len(tt.flatPairs)/2, len(pairs))

				for i := 0; i < len(tt.flatPairs); i += 2 {
					assert.Equal(t, tt.flatPairs[i], pairs[i/2].Key)
					assert.Equal(t, tt.flatPairs[i+1], pairs[i/2].Val)
				}
			}
		})
	}
}

func TestPackPairs_UserString(t *testing.T) {
	tests := []struct {
		name      string
		flatPairs []any
		wantErr   error
	}{
		{
			name: "",
			flatPairs: []any{
				user{name: "Tvux", age: 23}, "Boss",
				user{name: "Tom", age: 18}, "Developer",
				user{name: "Jerry", age: 20}, "Tester",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs, err := PackPairs[user, string](tt.flatPairs)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Equal(t, len(tt.flatPairs)/2, len(pairs))

				for i := 0; i < len(tt.flatPairs); i += 2 {
					assert.Equal(t, tt.flatPairs[i], pairs[i/2].Key)
					assert.Equal(t, tt.flatPairs[i+1], pairs[i/2].Val)
				}
			}
		})
	}
}

func TestPackPairs_Empty(t *testing.T) {
	flatPairs := []any{}
	pairs, err := PackPairs[string, int](flatPairs)
	assert.NoError(t, err)
	assert.Empty(t, pairs)
}

func TestPackPairs_Nil(t *testing.T) {
	pairs, err := PackPairs[string, int](nil)
	assert.NoError(t, err)
	assert.Nil(t, pairs)
}

func TestPackPairs_TypeAssertionError(t *testing.T) {
	tests := []struct {
		name      string
		flatPairs []any
		wantErr   error
	}{
		{
			name:      "",
			flatPairs: []any{1, "one", 2, "two", "three", "three"},
			wantErr:   NewErrTypeAssertionForKey(4),
		},
		{
			name:      "",
			flatPairs: []any{1, 1, 2, "two", "three", "three"},
			wantErr:   NewErrTypeAssertionForVal(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := PackPairs[int, string](tt.flatPairs)
			assert.Error(t, err)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func TestPackPairs_InvalidLength(t *testing.T) {
	flatPairs := []any{1, "one", 2.5, "two", "three"}
	_, err := PackPairs[int, string](flatPairs)
	assert.Error(t, err)
	assert.Equal(t, NewErrInvalidFlatPairsLength(), err)
}

type user struct {
	name string
	age  int
}
