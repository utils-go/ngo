package diagnostics

import (
	"os"
	"os/exec"
	"time"
)

// Stopwatch provides a set of methods and properties that you can use to accurately measure elapsed time
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.diagnostics.stopwatch?view=netframework-4.7.2
type Stopwatch struct {
	startTime time.Time
	elapsed   time.Duration
	isRunning bool
}

// Frequency gets the frequency of the timer as the number of ticks per second
// This field is read-only
var Frequency int64 = int64(time.Second / time.Nanosecond)

// IsHighResolution indicates whether the timer is based on a high-resolution performance counter
// This field is read-only
var IsHighResolution bool = true

// NewStopwatch initializes a new Stopwatch instance
func NewStopwatch() *Stopwatch {
	return &Stopwatch{
		elapsed:   0,
		isRunning: false,
	}
}

// StartNew initializes a new Stopwatch instance, sets the elapsed time property to zero, and starts measuring elapsed time
func StartNew() *Stopwatch {
	sw := NewStopwatch()
	sw.Start()
	return sw
}

// Start starts, or resumes, measuring elapsed time for an interval
func (sw *Stopwatch) Start() {
	if !sw.isRunning {
		sw.startTime = time.Now()
		sw.isRunning = true
	}
}

// Stop stops measuring elapsed time for an interval
func (sw *Stopwatch) Stop() {
	if sw.isRunning {
		sw.elapsed += time.Since(sw.startTime)
		sw.isRunning = false
	}
}

// Reset stops time interval measurement and resets the elapsed time to zero
func (sw *Stopwatch) Reset() {
	sw.elapsed = 0
	sw.isRunning = false
}

// Restart stops time interval measurement, resets the elapsed time to zero, and starts measuring elapsed time
func (sw *Stopwatch) Restart() {
	sw.Reset()
	sw.Start()
}

// IsRunning gets a value indicating whether the Stopwatch timer is running
func (sw *Stopwatch) IsRunning() bool {
	return sw.isRunning
}

// Elapsed gets the total elapsed time measured by the current instance
func (sw *Stopwatch) Elapsed() time.Duration {
	if sw.isRunning {
		return sw.elapsed + time.Since(sw.startTime)
	}
	return sw.elapsed
}

// ElapsedMilliseconds gets the total elapsed time measured by the current instance, in milliseconds
func (sw *Stopwatch) ElapsedMilliseconds() int64 {
	return sw.Elapsed().Nanoseconds() / int64(time.Millisecond)
}

// ElapsedTicks gets the total elapsed time measured by the current instance, in timer ticks
func (sw *Stopwatch) ElapsedTicks() int64 {
	return sw.Elapsed().Nanoseconds()
}

// GetTimestamp gets the current number of ticks in the timer mechanism
func GetTimestamp() int64 {
	return time.Now().UnixNano()
}

// Process 表示一个进程
type Process struct {
	cmd     *exec.Cmd
	process *os.Process
}

// ProcessStartInfo 表示进程启动信息
type ProcessStartInfo struct {
	FileName               string
	Arguments              string
	WorkingDirectory       string
	UseShellExecute        bool
	RedirectStandardInput  bool
	RedirectStandardOutput bool
	RedirectStandardError  bool
}

// Start 启动进程
func (p *Process) Start() error {
	err := p.cmd.Start()
	if err != nil {
		return err
	}
	p.process = p.cmd.Process
	return nil
}

// Kill 终止进程
func (p *Process) Kill() error {
	if p.process != nil {
		return p.process.Kill()
	}
	return nil
}

// WaitForExit 等待进程退出
func (p *Process) WaitForExit() error {
	if p.cmd != nil {
		return p.cmd.Wait()
	}
	return nil
}

// Id 获取进程 ID
func (p *Process) Id() int {
	if p.process != nil {
		return p.process.Pid
	}
	return 0
}

// ProcessName 获取进程名称
func (p *Process) ProcessName() string {
	if p.cmd != nil && p.cmd.Path != "" {
		return p.cmd.Path
	}
	return ""
}

// GetProcessById 根据进程 ID 获取进程
func GetProcessById(id int) (*Process, error) {
	process, err := os.FindProcess(id)
	if err != nil {
		return nil, err
	}

	return &Process{
		process: process,
	}, nil
}

// StartProcess 启动一个新进程
func StartProcess(startInfo *ProcessStartInfo) (*Process, error) {
	cmd := exec.Command(startInfo.FileName)
	if startInfo.Arguments != "" {
		cmd.Args = append(cmd.Args, startInfo.Arguments)
	}
	if startInfo.WorkingDirectory != "" {
		cmd.Dir = startInfo.WorkingDirectory
	}

	if startInfo.RedirectStandardInput {
		cmd.Stdin = os.Stdin
	}
	if startInfo.RedirectStandardOutput {
		cmd.Stdout = os.Stdout
	}
	if startInfo.RedirectStandardError {
		cmd.Stderr = os.Stderr
	}

	process := &Process{
		cmd: cmd,
	}

	err := process.Start()
	if err != nil {
		return nil, err
	}

	return process, nil
}

// EventLog 表示事件日志
type EventLog struct {
	// 简化实现
}

// WriteEntry 写入事件日志条目
func (e *EventLog) WriteEntry(message string, eventType string, eventID int) error {
	// 简化实现，实际应写入系统日志
	// 在 Go 中，可以使用 log 包或其他日志库
	return nil
}

// NewEventLog 创建一个新的 EventLog
func NewEventLog(source string) *EventLog {
	return &EventLog{}
}
