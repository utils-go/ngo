package diagnostics

import (
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