// Package timer provides .NET System.Timers.Timer-like functionality.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.timers.timer?view=netframework-4.7.2
package timer

import (
	"sync"
	"time"
)

// Timer provides a mechanism for executing a method at specified intervals.
// Equivalent to System.Timers.Timer in .NET.
type Timer struct {
	timer    *time.Ticker
	callback func()
	interval time.Duration
	enabled  bool
	autoReset bool
	stop     chan struct{}
	mu       sync.Mutex
}

// NewTimer creates a new Timer with the specified interval.
func NewTimer(interval time.Duration) *Timer {
	return &Timer{
		interval:  interval,
		autoReset: true,
		stop:      make(chan struct{}),
	}
}

// Interval gets the interval at which to raise the Elapsed event.
func (t *Timer) Interval() time.Duration {
	return t.interval
}

// SetInterval sets the interval.
func (t *Timer) SetInterval(d time.Duration) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.interval = d
	if t.enabled {
		t.timer.Reset(d)
	}
}

// AutoReset indicates whether the Timer should raise the Elapsed event each time the interval elapses.
func (t *Timer) AutoReset() bool {
	return t.autoReset
}

// SetAutoReset sets whether the timer repeats.
func (t *Timer) SetAutoReset(v bool) {
	t.autoReset = v
}

// Enabled indicates whether the Timer is running.
func (t *Timer) Enabled() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.enabled
}

// Start begins running the timer.
func (t *Timer) Start() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.enabled {
		return
	}
	t.enabled = true
	t.timer = time.NewTicker(t.interval)
	go t.run()
}

// Stop stops the timer.
func (t *Timer) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.enabled {
		return
	}
	t.enabled = false
	t.timer.Stop()
}

// Close releases all resources.
func (t *Timer) Close() {
	t.Stop()
}

// Elapsed registers a callback to be called when the interval elapses.
func (t *Timer) Elapsed(callback func()) {
	t.callback = callback
}

func (t *Timer) run() {
	for {
		select {
		case <-t.timer.C:
			if t.callback != nil {
				t.callback()
			}
			if !t.autoReset {
				t.Stop()
				return
			}
		case <-t.stop:
			return
		}
	}
}

// StopTimer is a convenience function to stop a timer after a duration.
func StopTimer(timer *time.Timer) bool {
	return timer.Stop()
}
