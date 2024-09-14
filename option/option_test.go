// Package option
/**
* @Project : GenericGo
* @File    : option_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/14 09:49
**/

package option

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	tests := []struct {
		name string
		u    *user
	}{
		{
			name: "valid user",
			u: &user{
				name: "Tvux",
				age:  23,
			},
		},
		{
			name: "empty name",
			u: &user{
				name: "",
				age:  23,
			},
		},
		{
			name: "negative age",
			u: &user{
				name: "Tvux",
				age:  -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newUser := &user{}
			Apply(newUser, withName(tt.u.name), withAge(tt.u.age))
			assert.Equal(t, tt.u, newUser)
		})
	}
}

func TestApplyErr(t *testing.T) {
	tests := []struct {
		name    string
		u       *user
		wantErr error
	}{
		{
			name: "valid user",
			u: &user{
				name: "Tvux",
				age:  23,
			},
		},
		{
			name: "empty name",
			u: &user{
				name: "",
				age:  23,
			},
			wantErr: errors.New("name cannot be empty"),
		},
		{
			name: "negative age",
			u: &user{
				name: "Tvux",
				age:  -1,
			},
			wantErr: errors.New("age must be greater than 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newUser := &user{}
			err := ApplyErr(newUser, withNameErr(tt.u.name), withAgeErr(tt.u.age))
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Equal(t, tt.u, newUser)
			}
		})
	}
}

type user struct {
	name string
	age  int
}

func withName(name string) Option[user] {
	return func(u *user) {
		u.name = name
	}
}

func withNameErr(name string) OptionErr[user] {
	return func(u *user) error {
		if name == "" {
			return errors.New("name cannot be empty")
		}
		u.name = name
		return nil
	}
}

func withAge(age int) Option[user] {
	return func(u *user) {
		u.age = age
	}
}

func withAgeErr(age int) OptionErr[user] {
	return func(u *user) error {
		if age <= 0 {
			return errors.New("age must be greater than 0")
		}
		u.age = age
		return nil
	}
}
