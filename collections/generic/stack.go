package generic

import (
	"fmt"
)

// Stack represents a last-in, first-out collection of objects.
// Equivalent to System.Collections.Generic.Stack<T> in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.stack-1?view=netframework-4.7.2
type Stack[T any] struct {
	items []T
}

// NewStack creates a new empty Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// NewStackWithCapacity creates a new Stack with the specified initial capacity.
func NewStackWithCapacity[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, capacity),
	}
}

// NewStackFromSlice creates a new Stack from a slice.
func NewStackFromSlice[T any](items []T) *Stack[T] {
	s := NewStackWithCapacity[T](len(items))
	s.items = append(s.items, items...)
	return s
}

// Count gets the number of elements in the Stack.
func (s *Stack[T]) Count() int {
	return len(s.items)
}

// IsEmpty returns whether the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Push inserts an object at the top of the Stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the object at the top of the Stack.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Peek returns the object at the top of the Stack without removing it.
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// Clear removes all objects from the Stack.
func (s *Stack[T]) Clear() {
	s.items = s.items[:0]
}

// Contains determines whether an element is in the Stack.
func (s *Stack[T]) Contains(item T) bool {
	for _, v := range s.items {
		if any(v) == any(item) {
			return true
		}
	}
	return false
}

// ToSlice copies the Stack elements to a new slice (top of stack is last element).
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, len(s.items))
	copy(result, s.items)
	return result
}

// ForEach performs the specified action on each element.
func (s *Stack[T]) ForEach(action func(T)) {
	for _, item := range s.items {
		action(item)
	}
}

// TrimExcess sets the capacity to the actual number of elements.
func (s *Stack[T]) TrimExcess() {
	if len(s.items) < cap(s.items) {
		newSlice := make([]T, len(s.items))
		copy(newSlice, s.items)
		s.items = newSlice
	}
}
