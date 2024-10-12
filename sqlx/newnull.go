// Package sqlx
/**
* @Project : GenericGo
* @File    : newnull.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/12 15:33
**/

package sqlx

import (
	"database/sql"
	"time"
)

// NewNullString 创建一个新的 sql.NullString，如果 val 不为空字符串，则 Valid 设置为 true。
func NewNullString(val string) sql.NullString {
	return sql.NullString{String: val, Valid: val != ""}
}

// NewNullInt64 创建一个新的 sql.NullInt64，如果 val 不为零值，则 Valid 设置为 true。
func NewNullInt64(val int64) sql.NullInt64 {
	return sql.NullInt64{Int64: val, Valid: val != 0}
}

// NewNullFloat64 创建一个新的 sql.NullFloat64，如果 val 不为零值，则 Valid 设置为 true。
func NewNullFloat64(val float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: val, Valid: val != 0}
}

// NewNullBool 创建一个新的 sql.NullBool，如果 val 为 true，则 Valid 设置为 true。
func NewNullBool(val bool) sql.NullBool {
	return sql.NullBool{Bool: val, Valid: val}
}

// NewNullTime 创建一个新的 sql.NullTime，如果 val 不是零时间，则 Valid 设置为 true。
func NewNullTime(val time.Time) sql.NullTime {
	return sql.NullTime{Time: val, Valid: !val.IsZero()}
}

// NewNullBytes 根据传入的字节切片，创建一个新的 sql.NullString，如果 val 不为空字节切片，则 Valid 设置为 true。
func NewNullBytes(val []byte) sql.NullString {
	return sql.NullString{String: string(val), Valid: len(val) > 0}
}
