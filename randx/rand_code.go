// Package randx
/**
* @Project : GenericGo
* @File    : rand_code.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/14 17:16
**/

package randx

import (
	"math/rand"

	"github.com/HJH0924/GenericGo/errs"
	"github.com/HJH0924/GenericGo/tuple"
)

type Type int

const (
	TypeDigit     Type = 1 << iota // 1 - 0001
	TypeLowerCase                  // 2 - 0010
	TypeUpperCase                  // 4 - 0100
	TypeSpecial                    // 8 - 1000

	TypeAllMixed = TypeDigit | TypeLowerCase | TypeUpperCase | TypeSpecial // 15 - 1111
)

const (
	CharsetDigit     = `0123456789`
	CharsetLowerCase = `abcdefghijklmnopqrstuvwxyz`
	CharsetUpperCase = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	CharsetSpecial   = ` ~!@#$%^&*()_+-=[]{};'\\:\"|,./<>?`
)

var (
	typeCharsetPairs = []tuple.Pair[Type, string]{
		tuple.NewPair(TypeDigit, CharsetDigit),
		tuple.NewPair(TypeLowerCase, CharsetLowerCase),
		tuple.NewPair(TypeUpperCase, CharsetUpperCase),
		tuple.NewPair(TypeSpecial, CharsetSpecial),
	}
)

// RandStrByType 根据传入的长度和类型生成随机字符串。
// 请保证输入的 length >= 0，否则会返回 NewErrLengthLessThanZero
// 请保证输入的 typ 的取值范围在 (0, TypeAllMixed] 内，否则会返回 NewErrTypeNotSupported
func RandStrByType(length int, typ Type) (string, error) {
	if length < 0 {
		return "", NewErrLengthLessThanZero()
	}
	if length == 0 {
		return "", nil
	}
	if typ <= 0 || typ > TypeAllMixed {
		return "", NewErrTypeNotSupported()
	}

	charset := ""
	for _, pair := range typeCharsetPairs {
		if (typ & pair.Key) == pair.Key {
			charset += pair.Val
		}
	}

	if charset == "" {
		return "", NewErrEmptyCharset()
	}

	return randStr(length, charset), nil
}

// RandStrByCharset 根据传入的长度和字符集生成随机字符串。
// 请保证输入的 length >= 0，否则会返回 NewErrLengthLessThanZero。
// 请保证输入的字符集不为空字符串，否则会返回 NewErrEmptyCharset。
// 字符集内部字符可以无序或重复。
func RandStrByCharset(length int, charset string) (string, error) {
	if length < 0 {
		return "", NewErrLengthLessThanZero()
	}
	if length == 0 {
		return "", nil
	}
	if charset == "" {
		return "", NewErrEmptyCharset()
	}

	return randStr(length, charset), nil
}

func randStr(length int, charset string) string {
	idxBits := 0
	charsetSize := len(charset)
	for charsetSize > (1<<idxBits)-1 {
		idxBits++
	}

	idxMask := (1 << idxBits) - 1
	remain := 63 / idxBits
	cache := rand.Int63()
	res := make([]byte, length)

	for i := 0; i < length; {
		if remain == 0 {
			cache, remain = rand.Int63(), 63/idxBits
		}

		if randIdx := int(cache & int64(idxMask)); randIdx < charsetSize {
			res[i] = charset[randIdx]
			i++
		}

		cache >>= idxBits
		remain--
	}

	return string(res)
}

// NewErrTypeNotSupported 创建一个表示不支持的类型的新错误。
func NewErrTypeNotSupported() error {
	return errs.WrapError("Unsupported type")
}

// NewErrLengthLessThanZero 创建一个表示长度必须大于等于0的新错误。
func NewErrLengthLessThanZero() error {
	return errs.WrapError("length must be greater than or equal to zero")
}

// NewErrEmptyCharset 创建一个表示字符集不能为空的新错误。
func NewErrEmptyCharset() error {
	return errs.WrapError("charset cannot be empty")
}
