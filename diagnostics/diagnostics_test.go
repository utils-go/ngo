package diagnostics

import (
	"os"
	"testing"
)

func TestProcess_GetProcessById(t *testing.T) {
	// 获取当前进程ID
	currentPID := os.Getpid()

	// 尝试获取当前进程
	process, err := GetProcessById(currentPID)
	if err != nil {
		t.Errorf("GetProcessById() failed: %v", err)
	}

	// 检查返回的进程是否非空
	if process == nil {
		t.Error("GetProcessById() should return a non-nil Process")
	}

	// 检查进程ID是否正确
	if process.Id() != currentPID {
		t.Errorf("Expected process ID %d, got %d", currentPID, process.Id())
	}
}

func TestProcess_StartProcess(t *testing.T) {
	// 准备启动信息 - 使用Windows系统的cmd.exe
	startInfo := &ProcessStartInfo{
		FileName:  "cmd.exe",
		Arguments: "/c echo Hello, World!",
	}

	// 启动进程
	process, err := StartProcess(startInfo)
	if err != nil {
		t.Errorf("StartProcess() failed: %v", err)
		return
	}

	// 检查返回的进程是否非空
	if process == nil {
		t.Error("StartProcess() should return a non-nil Process")
		return
	}

	// 等待进程完成
	err = process.WaitForExit()
	if err != nil {
		t.Errorf("WaitForExit() failed: %v", err)
	}
}

func TestEventLog_NewEventLog(t *testing.T) {
	// 创建EventLog
	eventLog := NewEventLog("TestSource")

	// 检查返回的EventLog是否非空
	if eventLog == nil {
		t.Error("NewEventLog() should return a non-nil EventLog")
	}
}

func TestEventLog_WriteEntry(t *testing.T) {
	// 创建EventLog
	eventLog := NewEventLog("TestSource")

	// 尝试写入日志条目
	err := eventLog.WriteEntry("Test message", "Information", 1000)
	if err != nil {
		t.Errorf("WriteEntry() failed: %v", err)
	}

	// 测试不同类型的日志条目
	err = eventLog.WriteEntry("Error message", "Error", 2000)
	if err != nil {
		t.Errorf("WriteEntry() with error type failed: %v", err)
	}
}
