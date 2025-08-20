package diagnostics

import (
	"testing"
	"time"
)

func TestNewStopwatch(t *testing.T) {
	sw := NewStopwatch()
	
	if sw.IsRunning() {
		t.Error("Expected new stopwatch to not be running")
	}
	
	if sw.ElapsedMilliseconds() != 0 {
		t.Errorf("Expected elapsed milliseconds to be 0, got %d", sw.ElapsedMilliseconds())
	}
}

func TestStartNew(t *testing.T) {
	sw := StartNew()
	
	if !sw.IsRunning() {
		t.Error("Expected StartNew stopwatch to be running")
	}
	
	// Let it run for a bit
	time.Sleep(10 * time.Millisecond)
	
	if sw.ElapsedMilliseconds() < 5 {
		t.Error("Expected some elapsed time")
	}
	
	sw.Stop()
}

func TestStartStop(t *testing.T) {
	sw := NewStopwatch()
	
	sw.Start()
	if !sw.IsRunning() {
		t.Error("Expected stopwatch to be running after Start()")
	}
	
	time.Sleep(10 * time.Millisecond)
	
	sw.Stop()
	if sw.IsRunning() {
		t.Error("Expected stopwatch to not be running after Stop()")
	}
	
	elapsed1 := sw.ElapsedMilliseconds()
	if elapsed1 < 5 {
		t.Error("Expected some elapsed time after running")
	}
	
	// Sleep more and ensure elapsed time doesn't change when stopped
	time.Sleep(10 * time.Millisecond)
	elapsed2 := sw.ElapsedMilliseconds()
	
	if elapsed2 != elapsed1 {
		t.Error("Expected elapsed time to not change when stopped")
	}
}

func TestReset(t *testing.T) {
	sw := NewStopwatch()
	sw.Start()
	
	time.Sleep(10 * time.Millisecond)
	sw.Stop()
	
	if sw.ElapsedMilliseconds() < 5 {
		t.Error("Expected some elapsed time before reset")
	}
	
	sw.Reset()
	
	if sw.IsRunning() {
		t.Error("Expected stopwatch to not be running after Reset()")
	}
	
	if sw.ElapsedMilliseconds() != 0 {
		t.Errorf("Expected elapsed time to be 0 after reset, got %d", sw.ElapsedMilliseconds())
	}
}

func TestRestart(t *testing.T) {
	sw := NewStopwatch()
	sw.Start()
	
	time.Sleep(10 * time.Millisecond)
	
	sw.Restart()
	
	if !sw.IsRunning() {
		t.Error("Expected stopwatch to be running after Restart()")
	}
	
	// Should have reset the elapsed time and started fresh
	if sw.ElapsedMilliseconds() > 5 {
		t.Error("Expected elapsed time to be near 0 after restart")
	}
	
	sw.Stop()
}

func TestElapsedWhileRunning(t *testing.T) {
	sw := NewStopwatch()
	sw.Start()
	
	time.Sleep(10 * time.Millisecond)
	
	elapsed1 := sw.ElapsedMilliseconds()
	
	time.Sleep(10 * time.Millisecond)
	
	elapsed2 := sw.ElapsedMilliseconds()
	
	if elapsed2 <= elapsed1 {
		t.Error("Expected elapsed time to increase while running")
	}
	
	sw.Stop()
}

func TestElapsedTicks(t *testing.T) {
	sw := NewStopwatch()
	sw.Start()
	
	time.Sleep(10 * time.Millisecond)
	sw.Stop()
	
	ticks := sw.ElapsedTicks()
	milliseconds := sw.ElapsedMilliseconds()
	
	// Ticks should be nanoseconds, milliseconds should be nanoseconds / 1,000,000
	expectedTicks := milliseconds * int64(time.Millisecond)
	
	// Allow some tolerance due to timing precision
	if ticks < expectedTicks/2 || ticks > expectedTicks*2 {
		t.Errorf("Expected ticks to be approximately %d, got %d", expectedTicks, ticks)
	}
}

func TestMultipleStartStop(t *testing.T) {
	sw := NewStopwatch()
	
	// First run
	sw.Start()
	time.Sleep(5 * time.Millisecond)
	sw.Stop()
	
	elapsed1 := sw.ElapsedMilliseconds()
	
	// Second run (should accumulate)
	sw.Start()
	time.Sleep(5 * time.Millisecond)
	sw.Stop()
	
	elapsed2 := sw.ElapsedMilliseconds()
	
	if elapsed2 <= elapsed1 {
		t.Error("Expected elapsed time to accumulate across multiple start/stop cycles")
	}
}

func TestGetTimestamp(t *testing.T) {
	timestamp1 := GetTimestamp()
	time.Sleep(1 * time.Millisecond)
	timestamp2 := GetTimestamp()
	
	if timestamp2 <= timestamp1 {
		t.Error("Expected timestamp to increase over time")
	}
}

func TestConstants(t *testing.T) {
	if Frequency <= 0 {
		t.Errorf("Expected positive frequency, got %d", Frequency)
	}
	
	if !IsHighResolution {
		t.Error("Expected high resolution timer")
	}
}