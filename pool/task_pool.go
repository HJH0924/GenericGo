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
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/HJH0924/GenericGo/option"
	"github.com/HJH0924/GenericGo/set"
)

var (
	// _ 是一个特殊的变量名，用于忽略未使用的变量值。
	_ TaskPool = &OnDemandBlockTaskPool{}
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

// Submit 提交一个任务
// 如果此时队列已满，那么将会阻塞调用者。
// 如果因为 ctx 的原因返回，那么将会返回 ctx.Err()
// 在调用 Start 前后都可以调用 Submit
func (Self *OnDemandBlockTaskPool) Submit(ctx context.Context, task Task) error {
	if task == nil {
		return NewErrInvalidTask
	}

	for {
		// 持续尝试提交任务，直到成功或发生错误
		if atomic.LoadInt32(&Self.state) == stateClosing {
			return NewErrTaskPoolIsClosing
		}
		if atomic.LoadInt32(&Self.state) == stateStopped {
			return NewErrTaskPoolIsStopped
		}

		taskWrap := &taskWrapper{
			task: task,
		}

		ok, err := Self.trySubmit(ctx, taskWrap, stateCreated)
		if ok || err != nil {
			// 任务提交成功，返回true，nil
			// 任务提交失败，返回false，err
			// 如果返回false，nil，则需要继续循环
			return err
		}

		ok, err = Self.trySubmit(ctx, taskWrap, stateRunning)
		if ok || err != nil {
			return err
		}
	}
}

// trySubmit 尝试提交任务，如果成功则返回 true，否则返回 false。
// 如果因为上下文 ctx 的原因返回，将会返回 ctx.Err()。
func (Self *OnDemandBlockTaskPool) trySubmit(ctx context.Context, task Task, state int32) (bool, error) {
	// if atomic.CompareAndSwapInt32(&Self.state, state, stateLocked) 等价于
	//if Self.state != state {
	//	Self.state = stateLocked
	//}
	// 但是是原子操作
	if atomic.CompareAndSwapInt32(&Self.state, state, stateLocked) {
		// 进入临界区
		defer atomic.CompareAndSwapInt32(&Self.state, stateLocked, state)

		select {
		case <-ctx.Done():
			return false, ctx.Err()
		case Self.taskQueue <- task:
			if state == stateRunning && Self.allow2CreateNewGoRoutine() {
				Self.increaseTotalGo(1)
				id := int(atomic.AddInt32(&Self.id, 1))
				go Self.goroutine(id)
			}
			return true, nil
		default:
			// 不能阻塞在临界区，要给 Shutdown 和 ShutdownNow 机会
			return false, nil
		}
	}
	return false, nil
}

// allow2CreateNewGoRoutine 判断是否可以创建新的 goroutine。
// 如果总goroutine数小于最大允许数，并且队列非空且积压率大于等于设定的阈值，则允许创建新的goroutine。
func (Self *OnDemandBlockTaskPool) allow2CreateNewGoRoutine() bool {
	Self.rwLock.RLock()
	defer Self.rwLock.RUnlock()

	backlogRate := float64(len(Self.taskQueue)) / float64(cap(Self.taskQueue))

	return Self.totalGo < Self.maxGo && backlogRate != 0 && backlogRate >= Self.queueBacklogRate
}

// increaseTotalGo 原子地增加 totalGo 的值。
func (Self *OnDemandBlockTaskPool) increaseTotalGo(n int32) {
	Self.rwLock.Lock()
	Self.totalGo += n
	Self.rwLock.Unlock()
}

// decreaseTotalGo 原子地减少 totalGo 的值。
func (Self *OnDemandBlockTaskPool) decreaseTotalGo(n int32) {
	Self.rwLock.Lock()
	Self.totalGo -= n
	Self.rwLock.Unlock()
}

// goroutine 是任务池中每个 Goroutine 运行的函数，负责持续从任务队列中获取并执行任务。
func (Self *OnDemandBlockTaskPool) goroutine(id int) {
	// 创建一个定时器，用于在 Goroutine 空闲时触发超时，0表示定时器立即开始计时。
	idleTimer := time.NewTimer(0)
	// 如果定时器已经启动，确保停止它，避免立即触发。
	if !idleTimer.Stop() {
		<-idleTimer.C
	}

	for {
		select {
		case <-Self.interruptCtx.Done():
			Self.decreaseTotalGo(1)
			return
		case <-idleTimer.C:
			Self.rwLock.Lock()
			Self.totalGo--
			Self.timeoutGoGroup.remove(id)
			Self.rwLock.Unlock()
			return
		case task, ok := <-Self.taskQueue:
			if !ok {
				// 任务队列已经关闭，减少 Goroutine 数量并退出
				Self.decreaseTotalGo(1)
				if Self.getTotalGo() == 0 {
					// 如果所有 Goroutine 都已完成，或者因调用Shutdown方法导致的协程退出
					// 最后一个退出的协程负责状态迁移及显示通知外部调用者
					if atomic.CompareAndSwapInt32(&Self.state, stateClosing, stateStopped) {
						Self.interruptCtxCancel()
					}
				}
				return
			}

			if Self.timeoutGoGroup.isIn(id) {
				// 如果 Goroutine 在超时组中，移除它，因为即将执行任务。
				Self.timeoutGoGroup.remove(id)
				// timer 只保证至少在等待d时间后才发送信号，而不是在d时间内发送信号
				// timer的Stop方法不保证一定成功
				// 不加判断并将信号清除可能会导致协程下次在 case <-idleTimer.C 处退出
				// 重置定时器
				if !idleTimer.Stop() {
					<-idleTimer.C
				}
			}

			atomic.AddInt32(&Self.numGoRunningTasks, 1)
			// TODO: handle error
			_ = task.Run(Self.interruptCtx)
			atomic.AddInt32(&Self.numGoRunningTasks, -1)

			Self.rwLock.Lock()
			// 检查是否还有任务可以执行
			noTask2Run := len(Self.taskQueue) == 0 || int32(len(Self.taskQueue)) < Self.totalGo
			if noTask2Run && Self.coreGo < Self.totalGo && Self.totalGo <= Self.maxGo {
				// 如果没有任务可以执行，并且 totalGo 超过了 coreGo数量，则减少 Goroutine 并退出
				Self.totalGo--
				Self.rwLock.Unlock()
				return
			}

			// 如果 Goroutine 数量在初始和核心 Goroutine 数量之间，设置一个新的定时器。
			// 1. 如果当前协程属于(initGo，coreGo]区间，需要为其分配一个超时器。
			//    - 当前协程在超时退出前（最大空闲时间内）尝试拿任务，拿到则继续执行，没拿到则超时退出。
			// 2. 如果当前协程属于(coreGo, maxGo]区间，且有任务可执行，也需要为其分配一个超时器兜底。
			//    - 因为此时看队列中有任务，等真去拿的时候可能恰好没任务
			//    - 这会导致当前协程总数（totalGo）长时间大于始协程数（initGo)直到队列再次有任务时才可能将当前总协程数准确地降至初始协程数
			if Self.initGo < Self.totalGo-int32(Self.timeoutGoGroup.size()) {
				idleTimer = time.NewTimer(defaultMaxIdleTime)
				Self.timeoutGoGroup.add(id)
			}

			Self.rwLock.Unlock()
		} // end of select
	} // end of for
} // end of func

// getTotalGo 用于查看TaskPool中有多少工作协程
// 直接返回 Self.totalGo 可能会导致竞态条件（race condition），
// 因为 Self.totalGo 可能在读取过程中被其他 goroutine 修改。
// 通过使用局部变量 totalGo，可以确保在读取 Self.totalGo 后，
// 即使其他 goroutine 修改了 Self.totalGo，也不会影响已经读取的值。
func (Self *OnDemandBlockTaskPool) getTotalGo() int32 {
	var totalGo int32
	Self.rwLock.RLock()
	totalGo = Self.totalGo
	Self.rwLock.RUnlock()
	return totalGo
}

func (Self *OnDemandBlockTaskPool) Start() error {
	//TODO implement me
	panic("implement me")
}

func (Self *OnDemandBlockTaskPool) Shutdown() (<-chan ShutdownSignal, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *OnDemandBlockTaskPool) ShutdownNow() ([]Task, error) {
	//TODO implement me
	panic("implement me")
}

func (Self *OnDemandBlockTaskPool) States(ctx context.Context, interval time.Duration) (<-chan State, error) {
	//TODO implement me
	panic("implement me")
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

// taskWrapper Task 的装饰器
type taskWrapper struct {
	task Task
}

// Run 执行包装的任务，并捕获 panic 异常。
func (Self *taskWrapper) Run(ctx context.Context) (err error) {
	defer func() {
		// 处理 panic
		if r := recover(); r != nil {
			buf := make([]byte, panicBuffLen)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("%w: %s", NewErrTaskRunningPanic, fmt.Sprintf("[PANIC]:\t%+v\n%s\n", r, buf))
			// 这里可以将异常信息输出到日志，或者进行其他处理
		}
	}()

	return Self.task.Run(ctx)
}
