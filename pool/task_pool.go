// Package pool
/**
* @Project : GenericGo
* @File    : task_pool.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/24 15:49
**/

package pool

import (
	"errors"
	"time"
)

// 状态常量定义
const (
	stateCreated int32 = iota + 1 // 1
	stateRunning                  // 2
	stateClosing                  // 3
	stateStopped                  // 4
	stateLocked                   // 5
)

const (
	panicBuffLen       = 2048
	defaultMaxIdleTime = 10 * time.Second
)

// 错误定义
var (
	NewErrTaskPoolIsNotRunning = errors.New("taskpool is not running")
	NewErrTaskPoolIsClosing    = errors.New("taskpool is closing")
	NewErrTaskPoolIsStopped    = errors.New("taskpool is stopped")
	NewErrTaskPoolIsStarted    = errors.New("taskpool is started")
	NewErrTaskRunningPanic     = errors.New("task running panic")
	NewErrInvalidTask          = errors.New("invalid task")
	NewErrInvalidArgument      = errors.New("invalid argument")
)
