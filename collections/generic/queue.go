package generic

import (
	"fmt"
)

// Queue represents a first-in, first-out collection of objects.
// Equivalent to System.Collections.Generic.Queue<T> in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.queue-1?view=netframework-4.7.2
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new empty Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// NewQueueWithCapacity creates a new Queue with the specified initial capacity.
func NewQueueWithCapacity[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, capacity),
	}
}

// NewQueueFromSlice creates a new Queue from a slice.
func NewQueueFromSlice[T any](items []T) *Queue[T] {
	q := NewQueueWithCapacity[T](len(items))
	q.items = append(q.items, items...)
	return q
}

// Count gets the number of elements in the Queue.
func (q *Queue[T]) Count() int {
	return len(q.items)
}

// IsEmpty returns whether the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Enqueue adds an object to the end of the Queue.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the object at the beginning of the Queue.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Peek returns the object at the beginning of the Queue without removing it.
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

// Clear removes all objects from the Queue.
func (q *Queue[T]) Clear() {
	q.items = q.items[:0]
}

// Contains determines whether an element is in the Queue.
func (q *Queue[T]) Contains(item T) bool {
	for _, v := range q.items {
		if any(v) == any(item) {
			return true
		}
	}
	return false
}

// ToSlice copies the Queue elements to a new slice.
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, len(q.items))
	copy(result, q.items)
	return result
}

// ForEach performs the specified action on each element.
func (q *Queue[T]) ForEach(action func(T)) {
	for _, item := range q.items {
		action(item)
	}
}

// TrimExcess sets the capacity to the actual number of elements.
func (q *Queue[T]) TrimExcess() {
	if len(q.items) < cap(q.items) {
		newSlice := make([]T, len(q.items))
		copy(newSlice, q.items)
		q.items = newSlice
	}
}
