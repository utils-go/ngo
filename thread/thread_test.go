package thread

import (
	"testing"
	"time"
)

func TestThread_StartAndJoin(t *testing.T) {
	var done bool
	
	// 创建一个线程，执行后设置done为true
	th := NewThread(func() {
		time.Sleep(100 * time.Millisecond)
		done = true
	})
	
	// 启动线程
	th.Start()
	
	// 检查线程是否活跃
	if !th.IsAlive() {
		t.Error("Thread should be alive after Start()")
	}
	
	// 等待线程完成
	th.Join()
	
	// 检查线程是否已停止
	if th.IsAlive() {
		t.Error("Thread should not be alive after Join()")
	}
	
	// 检查任务是否完成
	if !done {
		t.Error("Thread function should have executed")
	}
}

func TestThread_NameAndPriority(t *testing.T) {
	th := NewThread(func() {})
	
	// 测试名称设置和获取
	testName := "TestThread"
	th.SetName(testName)
	if th.Name() != testName {
		t.Errorf("Expected name %s, got %s", testName, th.Name())
	}
	
	// 测试优先级设置和获取
	testPriority := 5
	th.SetPriority(testPriority)
	if th.Priority() != testPriority {
		t.Errorf("Expected priority %d, got %d", testPriority, th.Priority())
	}
}

func TestThread_State(t *testing.T) {
	th := NewThread(func() {
		time.Sleep(50 * time.Millisecond)
	})
	
	// 初始状态应该是Stopped
	if th.State() != ThreadStateStopped {
		t.Errorf("Expected initial state ThreadStateStopped, got %v", th.State())
	}
	
	// 启动后状态应该是Running
	th.Start()
	if th.State() != ThreadStateRunning {
		t.Errorf("Expected state ThreadStateRunning after Start(), got %v", th.State())
	}
	
	// 等待完成后状态应该是Stopped
	th.Join()
	if th.State() != ThreadStateStopped {
		t.Errorf("Expected state ThreadStateStopped after Join(), got %v", th.State())
	}
}

func TestThread_Sleep(t *testing.T) {
	start := time.Now()
	
	// 休眠100毫秒
	Sleep(100)
	
	elapsed := time.Since(start)
	// 检查是否至少休眠了90毫秒
	if elapsed < 90*time.Millisecond {
		t.Errorf("Expected sleep of at least 90ms, got %v", elapsed)
	}
}

func TestThread_CurrentThread(t *testing.T) {
	th := CurrentThread()
	
	// 检查当前线程是否非空
	if th == nil {
		t.Error("CurrentThread() should return a non-nil Thread")
	}
	
	// 检查当前线程是否标记为活跃
	if !th.IsAlive() {
		t.Error("CurrentThread() should return an alive Thread")
	}
	
	// 检查当前线程状态是否为Running
	if th.State() != ThreadStateRunning {
		t.Errorf("Expected current thread state ThreadStateRunning, got %v", th.State())
	}
}
