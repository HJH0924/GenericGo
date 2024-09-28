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
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/HJH0924/GenericGo/option"
	"github.com/HJH0924/GenericGo/set"
)

// OnDemandBlockTaskPool 按需创建 Goroutine 的并发阻塞的任务池
type OnDemandBlockTaskPool struct {
	state  int32        // 任务池当前的状态
	rwLock sync.RWMutex // 读写锁

	taskQueue chan Task // 任务队列，存放等待执行的任务

	numGoRunningTasks int32 // 当前正在执行任务的 Goroutine 数量
	totalGo           int32 // 当前任务池中的 Goroutine 总数

	initGo int32 // 初始 Goroutine 数量
	coreGo int32 // 核心 Goroutine 数量，通常在低负载时保持
	maxGo  int32 // 允许的最大 Goroutine 数量

	timeoutGoGroup *timeoutGroup // 超时 Goroutine 组，用于管理需要超时控制的 Goroutine
	maxIdleTime    time.Duration // Goroutine 空闲时的最长时间

	queueBacklogRate float64 // 队列积压率，用于决定何时创建新的 Goroutine

	id int32 // 用于生成 Goroutine 的唯一标识符

	// 上下文管理，用于优雅地中断 Goroutine
	interruptCtx       context.Context    // 用于中断 Goroutine 的上下文
	interruptCtxCancel context.CancelFunc // 用于取消 interruptCtx 的函数
}

// NewOnDemandBlockTaskPool 创建一个新的 OnDemandBlockTaskPool
// initGo 是初始协程数
// taskQueueSize 是队列大小，即最多有多少个任务在等待调度
// 使用相应的Option选项可以动态扩展协程数
func NewOnDemandBlockTaskPool(initGo int, taskQueueSize int, opts ...option.Option[OnDemandBlockTaskPool]) (*OnDemandBlockTaskPool, error) {
	// check pool params
	if initGo <= 0 {
		return nil, NewErrInitGoInvalid
	}
	if taskQueueSize < 0 {
		return nil, NewErrTaskQueueSizeInvalid
	}

	pool := &OnDemandBlockTaskPool{
		taskQueue:      make(chan Task, taskQueueSize),
		initGo:         int32(initGo),
		coreGo:         int32(initGo),
		maxGo:          int32(initGo),
		maxIdleTime:    defaultMaxIdleTime,
		timeoutGoGroup: newTimeOutGroup(),
	}

	ctx := context.Background()
	pool.interruptCtx, pool.interruptCtxCancel = context.WithCancel(ctx)

	atomic.StoreInt32(&pool.state, stateCreated)

	option.Apply(pool, opts...)

	// check pool params
	if pool.queueBacklogRate < float64(0) || pool.queueBacklogRate > float64(1) {
		return nil, NewErrQueueBacklogRateInvalid
	}
	if pool.coreGo != pool.initGo && pool.maxGo == pool.initGo {
		// 确保在高负载时任务池不会超过 coreGo 的数量
		pool.maxGo = pool.coreGo
	} else if pool.coreGo == pool.initGo && pool.maxGo != pool.initGo {
		// 确保任务池在一般情况下不会低于 maxGo 的数量
		pool.coreGo = pool.maxGo
	}
	if !(pool.initGo <= pool.coreGo && pool.coreGo <= pool.maxGo) {
		return nil, NewErrGoroutineConditionNotMet
	}

	return pool, nil
}

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

	NewErrInvalidArgument          = errors.New("invalid argument")
	NewErrInitGoInvalid            = fmt.Errorf("%w: initGo should be greater than 0", NewErrInvalidArgument)
	NewErrTaskQueueSizeInvalid     = fmt.Errorf("%w: taskQueueSize should be greater than or equal to 0", NewErrInvalidArgument)
	NewErrQueueBacklogRateInvalid  = fmt.Errorf("%w: queueBacklogRate should be within the range [0.0, 1.0]", NewErrInvalidArgument)
	NewErrGoroutineConditionNotMet = fmt.Errorf("%w: initGo, coreGo, and maxGo must satisfy the condition: initGo <= coreGo <= maxGo", NewErrInvalidArgument)
)

// WithQueueBacklogRate 设置任务池的队列积压率选项。
func WithQueueBacklogRate(queueBacklogRate float64) option.Option[OnDemandBlockTaskPool] {
	return func(pool *OnDemandBlockTaskPool) {
		pool.queueBacklogRate = queueBacklogRate
	}
}

// WithCoreGo 设置任务池的核心协程数选项。
func WithCoreGo(coreGo int) option.Option[OnDemandBlockTaskPool] {
	return func(pool *OnDemandBlockTaskPool) {
		pool.coreGo = int32(coreGo)
	}
}

// WithMaxGo 设置任务池的最大协程数选项。
func WithMaxGo(maxGo int) option.Option[OnDemandBlockTaskPool] {
	return func(pool *OnDemandBlockTaskPool) {
		pool.maxGo = int32(maxGo)
	}
}

// WithMaxIdleTime 设置任务池的最大空闲时间选项。
func WithMaxIdleTime(maxIdleTime time.Duration) option.Option[OnDemandBlockTaskPool] {
	return func(pool *OnDemandBlockTaskPool) {
		pool.maxIdleTime = maxIdleTime
	}
}

type timeoutGroup struct {
	group  set.Set[int]
	rwLock sync.RWMutex
}

func newTimeOutGroup() *timeoutGroup {
	return &timeoutGroup{
		group: set.NewHashSet[int](),
	}
}

func (Self *timeoutGroup) isIn(id int) bool {
	Self.rwLock.RLock()
	defer Self.rwLock.RUnlock()
	return Self.group.Contains(id)
}

func (Self *timeoutGroup) add(id int) {
	Self.rwLock.Lock()
	defer Self.rwLock.Unlock()
	Self.group.Add(id)
}

func (Self *timeoutGroup) remove(id int) {
	Self.rwLock.Lock()
	defer Self.rwLock.Unlock()
	Self.group.Remove(id)
}

func (Self *timeoutGroup) size() int {
	Self.rwLock.RLock()
	defer Self.rwLock.RUnlock()
	return Self.group.Size()
}
