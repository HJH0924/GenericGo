// Package randx
/**
* @Project : GenericGo
* @File    : rand_code_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/24 13:52
**/

package randx

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStrByType(t *testing.T) {
	tests := []struct {
		name      string
		length    int
		typ       Type
		wantMatch string
		wantErr   error
	}{
		{
			name:      "数字验证码",
			length:    100,
			typ:       TypeDigit, // 1
			wantMatch: `^[0-9]+$`,
			wantErr:   nil,
		},
		{
			name:      "小写字母验证码",
			length:    100,
			typ:       TypeLowerCase, // 2
			wantMatch: `^[a-z]+$`,
			wantErr:   nil,
		},
		{
			name:      "数字+小写字母验证码",
			length:    100,
			typ:       TypeDigit | TypeLowerCase, // 3
			wantMatch: `^[a-z0-9]+$`,
			wantErr:   nil,
		},
		{
			name:      "大写字母验证码",
			length:    100,
			typ:       TypeUpperCase, // 4
			wantMatch: `^[A-Z]+$`,
			wantErr:   nil,
		},
		{
			name:      "数字+大写字母验证码",
			length:    100,
			typ:       TypeUpperCase, // 5
			wantMatch: `^[A-Z0-9]+$`,
			wantErr:   nil,
		},
		{
			name:      "大小写字母验证码",
			length:    100,
			typ:       TypeLowerCase | TypeUpperCase, // 6
			wantMatch: `^[a-zA-Z]+$`,
			wantErr:   nil,
		},
		{
			name:      "数字+大小写字母验证码",
			length:    100,
			typ:       TypeDigit | TypeLowerCase | TypeUpperCase, // 7
			wantMatch: `^[0-9a-zA-Z]+$`,
			wantErr:   nil,
		},
		{
			name:      "特殊字符验证码",
			length:    100,
			typ:       TypeSpecial, // 8
			wantMatch: `^[^0-9a-zA-Z]+$`,
			wantErr:   nil,
		},
		{
			name:      "数字+特殊字符+大小写字母验证码",
			length:    100,
			typ:       TypeAllMixed, // 15
			wantMatch: `^[\S\s]+$`,
			wantErr:   nil,
		},
		{
			name:      "未定义类型(超过范围)",
			length:    100,
			typ:       TypeAllMixed + 1,
			wantMatch: "",
			wantErr:   NewErrTypeNotSupported,
		},
		{
			name:      "未定义类型(0)",
			length:    100,
			typ:       0,
			wantMatch: "",
			wantErr:   NewErrTypeNotSupported,
		},
		{
			name:      "长度小于0",
			length:    -1,
			typ:       0,
			wantMatch: "",
			wantErr:   NewErrLengthLessThanZero,
		},
		{
			name:      "长度等于0",
			length:    0,
			typ:       TypeAllMixed,
			wantMatch: "",
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, err := RandStrByType(tt.length, tt.typ)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
				assert.Equal(t, "", str)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.length, len(str))
				matched, err := regexp.MatchString(tt.wantMatch, str)
				assert.NoError(t, err)
				assert.True(t, matched)
			}
		})
	}
}

func TestRandStrByCharset(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		charset string
		wantErr error
	}{
		{
			name:    "长度小于0",
			length:  -1,
			charset: CharsetDigit,
			wantErr: NewErrLengthLessThanZero,
		},
		{
			name:    "长度等于0",
			length:  0,
			charset: CharsetDigit,
			wantErr: nil,
		},
		{
			name:    "字符集为空",
			length:  100,
			charset: "",
			wantErr: NewErrEmptyCharset,
		},
		{
			name:    "数字验证码",
			length:  100,
			charset: CharsetDigit, // 1
			wantErr: nil,
		},
		{
			name:    "小写字母验证码",
			length:  100,
			charset: CharsetLowerCase, // 2
			wantErr: nil,
		},
		{
			name:    "数字+小写字母验证码",
			length:  100,
			charset: CharsetDigit + CharsetLowerCase, // 3
			wantErr: nil,
		},
		{
			name:    "大写字母验证码",
			length:  100,
			charset: CharsetUpperCase, // 4
			wantErr: nil,
		},
		{
			name:    "数字+大写字母验证码",
			length:  100,
			charset: CharsetDigit + CharsetUpperCase, // 5
			wantErr: nil,
		},
		{
			name:    "大小写字母验证码",
			length:  100,
			charset: CharsetLowerCase + CharsetUpperCase, // 6
			wantErr: nil,
		},
		{
			name:    "数字+大小写字母验证码",
			length:  100,
			charset: CharsetDigit + CharsetLowerCase + CharsetUpperCase, // 7
			wantErr: nil,
		},
		{
			name:    "特殊字符验证码",
			length:  100,
			charset: CharsetSpecial, // 8
			wantErr: nil,
		},
		{
			name:    "数字+特殊字符验证码",
			length:  100,
			charset: CharsetDigit + CharsetSpecial, // 9
			wantErr: nil,
		},
		{
			name:    "小写字母+特殊字符验证码",
			length:  100,
			charset: CharsetLowerCase + CharsetSpecial, // 10
			wantErr: nil,
		},
		{
			name:    "数字+小写字母+特殊字符验证码",
			length:  100,
			charset: CharsetDigit + CharsetLowerCase + CharsetSpecial, // 11
			wantErr: nil,
		},
		{
			name:    "大写字母+特殊字符验证码",
			length:  100,
			charset: CharsetUpperCase + CharsetSpecial, // 12
			wantErr: nil,
		},
		{
			name:    "数字+大写字母+特殊字符验证码",
			length:  100,
			charset: CharsetDigit + CharsetUpperCase + CharsetSpecial, // 13
			wantErr: nil,
		},
		{
			name:    "大小写字母+特殊字符验证码",
			length:  100,
			charset: CharsetLowerCase + CharsetUpperCase + CharsetSpecial, // 14
			wantErr: nil,
		},
		{
			name:    "数字+特殊字符+大小写字母验证码",
			length:  100,
			charset: CharsetDigit + CharsetSpecial + CharsetLowerCase + CharsetUpperCase, // 15
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, err := RandStrByCharset(tt.length, tt.charset)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
				assert.Equal(t, "", str)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.length, len(str))
				assert.NoError(t, err)
				assert.True(t, matchFunc(str, tt.charset))
			}
		})
	}
}

func matchFunc(randStr string, charset string) bool {
	for _, char := range randStr {
		if !strings.ContainsRune(charset, char) {
			return false
		}
	}
	return true
}
