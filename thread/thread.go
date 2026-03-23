package thread

import (
	"fmt"
	"sync"
	"time"
)

// ThreadState 表示线程的状态
type ThreadState int

const (
	ThreadStateStopped ThreadState = iota
	ThreadStateRunning
	ThreadStateWaitSleepJoin
)

// Thread 表示一个线程
type Thread struct {
	name       string
	priority   int
	isAlive    bool
	state      ThreadState
	mu         sync.Mutex
	waiter     sync.WaitGroup
	threadFunc func()
}

// NewThread 创建一个新的线程
func NewThread(f func()) *Thread {
	return &Thread{
		threadFunc: f,
	}
}

// Start 启动线程
func (t *Thread) Start() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.isAlive {
		return
	}

	t.isAlive = true
	t.state = ThreadStateRunning
	t.waiter.Add(1)

	go func() {
		defer func() {
			t.mu.Lock()
			t.isAlive = false
			t.state = ThreadStateStopped
			t.mu.Unlock()
			t.waiter.Done()
		}()

		t.threadFunc()
	}()
}

// Join 等待线程完成
func (t *Thread) Join() {
	t.waiter.Wait()
}

// Sleep 使当前线程休眠指定的毫秒数
func Sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

// CurrentThread 获取当前线程
func CurrentThread() *Thread {
	// 简化实现，返回一个表示当前 goroutine 的 Thread 对象
	return &Thread{
		name:    fmt.Sprintf("goroutine-%d", time.Now().UnixNano()),
		isAlive: true,
		state:   ThreadStateRunning,
	}
}

// IsAlive 检查线程是否活跃
func (t *Thread) IsAlive() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.isAlive
}

// Name 获取或设置线程名称
func (t *Thread) Name() string {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.name
}

func (t *Thread) SetName(name string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.name = name
}

// Priority 获取或设置线程优先级
func (t *Thread) Priority() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.priority
}

func (t *Thread) SetPriority(priority int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.priority = priority
}

// State 获取线程状态
func (t *Thread) State() ThreadState {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.state
}
