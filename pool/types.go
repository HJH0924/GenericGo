// Package pool
/**
* @Project : GenericGo
* @File    : types.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/14 09:13
**/

package pool

import (
	"context"
	"time"
)

// TaskPool 定义了一个任务池的接口，它能够提交任务、启动任务调度、关闭任务池，并提供任务池状态信息。
type TaskPool interface {
	// Submit 提交一个任务到任务池中执行。
	// ctx 提供了一个可以控制任务提交超时的上下文。
	// task 是要提交的任务。
	// 如果在 ctx 过期前任务没有被成功提交，或者任务池已经关闭，将返回错误。
	Submit(ctx context.Context, task Task) error

	// Start 启动任务池，开始调度任务执行。
	// 在调用此方法之前，任务池不会执行任何任务。
	// 某些实现可能允许在调用 Start 之后继续提交任务。
	Start() error

	// Shutdown 安全关闭任务池，停止接手新任务，等待已提交的任务执行完成。
	// 如果任务池尚未启动，则立即返回。
	// 返回一个 chan struct{}，当所有任务执行完毕时，会向该通道发送一个信号（空结构体）作为通知。
	Shutdown() (<-chan struct{}, error)

	// ShutdownNow 立即关闭任务池，尝试中断正在执行的任务，并返回所有未执行的任务列表。
	// 能否中断任务取决于 TaskPool 和 Task 的具体实现。
	ShutdownNow() ([]Task, error)

	// States 提供一个通道，定期发送任务池的运行状态。
	// ctx 用于控制何时停止发送状态信息。
	// interval 指定发送状态信息的时间间隔。
	States(ctx context.Context, interval time.Duration) (<-chan State, error)
}

// Task 定义了一个任务的接口，包含 Run 方法用于执行任务。
type Task interface {
	// Run 执行任务的方法。
	// ctx 提供了任务执行的上下文，可以实现超时控制。
	Run(ctx context.Context) error
}

// State 定义了任务池的运行状态结构体
type State struct {
	PoolState      int32 // 任务池的状态码
	GoCnt          int32 // 当前任务池中 Goroutine 的数量
	QueueSize      int32 // 任务队列的大小
	WaitingTaskCnt int32 // 等待执行的任务数量
	RunningTaskCnt int32 // 当前正在执行的任务数量
	Timestamp      int64 // 状态记录的时间戳
}
