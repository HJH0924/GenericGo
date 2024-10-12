// Package sqlx
/**
* @Project : GenericGo
* @File    : newnull_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/12 15:44
**/

package sqlx

import (
	"database/sql"
	"reflect"
	"testing"
	"time"
)

func TestNewNullString(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "non-empty string",
			args: args{val: "name"},
			want: sql.NullString{String: "name", Valid: true},
		},
		{
			name: "empty string",
			args: args{val: ""},
			want: sql.NullString{String: "", Valid: false},
		},
		{
			name: "zero value",
			want: sql.NullString{String: "", Valid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullString(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullInt64(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		args args
		want sql.NullInt64
	}{
		{
			name: "non-zero value",
			args: args{val: 123},
			want: sql.NullInt64{Int64: 123, Valid: true},
		},
		{
			name: "zero value",
			args: args{val: 0},
			want: sql.NullInt64{Int64: 0, Valid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullInt64(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullFloat64(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		{
			name: "non-zero value",
			args: args{val: 123.4},
			want: sql.NullFloat64{Float64: 123.4, Valid: true},
		},
		{
			name: "zero value",
			args: args{val: 0},
			want: sql.NullFloat64{Float64: 0, Valid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullFloat64(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullBool(t *testing.T) {
	type args struct {
		val bool
	}
	tests := []struct {
		name string
		args args
		want sql.NullBool
	}{
		{
			name: "true value",
			args: args{val: true},
			want: sql.NullBool{Bool: true, Valid: true},
		},
		{
			name: "false value",
			args: args{val: false},
			want: sql.NullBool{Bool: false, Valid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullBool(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullTime(t *testing.T) {
	now := time.Now()
	type args struct {
		val time.Time
	}
	tests := []struct {
		name string
		args args
		want sql.NullTime
	}{
		{
			name: "valid time",
			args: args{val: now},
			want: sql.NullTime{Time: now, Valid: true},
		},
		{
			name: "zero time",
			args: args{val: time.Time{}},
			want: sql.NullTime{Time: time.Time{}, Valid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullTime(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNullBytes(t *testing.T) {
	type args struct {
		val []byte
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "non-empty bytes",
			args: args{val: []byte("name")},
			want: sql.NullString{String: "name", Valid: true},
		},
		{
			name: "empty bytes",
			args: args{val: []byte{}},
			want: sql.NullString{String: "", Valid: false},
		},
		{
			name: "nil bytes",
			args: args{val: nil},
			want: sql.NullString{String: "", Valid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNullBytes(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNullBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
