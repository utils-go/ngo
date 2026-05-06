// Package threading provides .NET System.Threading-like synchronization primitives.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.threading?view=netframework-4.7.2
package threading

import (
	"fmt"
	"sync"
	"time"
)

// ---- Mutex ----

// Mutex implements a mutual exclusion lock. Equivalent to System.Threading.Mutex in .NET.
type Mutex struct {
	mu    sync.Mutex
	name  string
}

// NewMutex creates a new Mutex.
func NewMutex() *Mutex {
	return &Mutex{}
}

// NewNamedMutex creates a named Mutex.
func NewNamedMutex(name string) *Mutex {
	return &Mutex{name: name}
}

// WaitOne acquires the mutex, blocking until it is available.
func (m *Mutex) WaitOne() {
	m.mu.Lock()
}

// WaitOneWithTimeout acquires the mutex with a timeout.
func (m *Mutex) WaitOneWithTimeout(timeout time.Duration) bool {
	done := make(chan struct{})
	go func() {
		m.mu.Lock()
		close(done)
	}()
	select {
	case <-done:
		return true
	case <-time.After(timeout):
		return false
	}
}

// ReleaseMutex releases the mutex.
func (m *Mutex) ReleaseMutex() {
	m.mu.Unlock()
}

// ---- Semaphore ----

// Semaphore limits the number of threads that can access a resource concurrently.
// Equivalent to System.Threading.Semaphore in .NET.
type Semaphore struct {
	ch       chan struct{}
	maximum  int
}

// NewSemaphore creates a new Semaphore with the specified initial and maximum count.
func NewSemaphore(initialCount, maximumCount int) *Semaphore {
	if initialCount < 0 || maximumCount < 1 || initialCount > maximumCount {
		panic(fmt.Sprintf("invalid semaphore counts: initial=%d, max=%d", initialCount, maximumCount))
	}
	s := &Semaphore{
		ch:      make(chan struct{}, maximumCount),
		maximum: maximumCount,
	}
	for i := 0; i < initialCount; i++ {
		s.ch <- struct{}{}
	}
	return s
}

// WaitOne blocks until the semaphore can be entered.
func (s *Semaphore) WaitOne() {
	s.ch <- struct{}{}
}

// WaitOneWithTimeout waits with a timeout.
func (s *Semaphore) WaitOneWithTimeout(timeout time.Duration) bool {
	select {
	case s.ch <- struct{}{}:
		return true
	case <-time.After(timeout):
		return false
	}
}

// Release releases the semaphore.
func (s *Semaphore) Release() int {
	<-s.ch
	return len(s.ch) + 1
}

// ---- AutoResetEvent ----

// AutoResetEvent notifies a waiting thread that an event has occurred, resetting automatically.
// Equivalent to System.Threading.AutoResetEvent in .NET.
type AutoResetEvent struct {
	ch chan struct{}
}

// NewAutoResetEvent creates a new AutoResetEvent.
func NewAutoResetEvent(initialState bool) *AutoResetEvent {
	e := &AutoResetEvent{
		ch: make(chan struct{}, 1),
	}
	if initialState {
		e.ch <- struct{}{}
	}
	return e
}

// WaitOne waits for the event to be signaled.
func (e *AutoResetEvent) WaitOne() {
	<-e.ch
}

// WaitOneWithTimeout waits with a timeout.
func (e *AutoResetEvent) WaitOneWithTimeout(timeout time.Duration) bool {
	select {
	case <-e.ch:
		return true
	case <-time.After(timeout):
		return false
	}
}

// Set signals the event, releasing one waiting thread.
// The event is automatically reset after releasing a single thread.
func (e *AutoResetEvent) Set() bool {
	select {
	case e.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

// Reset sets the state of the event to non-signaled.
func (e *AutoResetEvent) Reset() bool {
	select {
	case <-e.ch:
		return true
	default:
		return false
	}
}

// ---- ManualResetEvent ----

// ManualResetEvent notifies one or more waiting threads that an event has occurred.
// Equivalent to System.Threading.ManualResetEvent in .NET.
type ManualResetEvent struct {
	mu   sync.Mutex
	ch   chan struct{}
	open bool
}

// NewManualResetEvent creates a new ManualResetEvent.
func NewManualResetEvent(initialState bool) *ManualResetEvent {
	e := &ManualResetEvent{
		ch: make(chan struct{}),
	}
	if initialState {
		e.Set()
	}
	return e
}

// WaitOne waits for the event to be signaled.
func (e *ManualResetEvent) WaitOne() {
	e.mu.Lock()
	ch := e.ch
	e.mu.Unlock()
	<-ch
}

// WaitOneWithTimeout waits with a timeout.
func (e *ManualResetEvent) WaitOneWithTimeout(timeout time.Duration) bool {
	e.mu.Lock()
	ch := e.ch
	e.mu.Unlock()
	select {
	case <-ch:
		return true
	case <-time.After(timeout):
		return false
	}
}

// Set signals the event, releasing all waiting threads.
func (e *ManualResetEvent) Set() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.open {
		return false
	}
	e.open = true
	close(e.ch)
	return true
}

// Reset sets the state to non-signaled, causing threads to block.
func (e *ManualResetEvent) Reset() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	if !e.open {
		return false
	}
	e.open = false
	e.ch = make(chan struct{})
	return true
}
