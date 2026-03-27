package generic

import (
	"fmt"
	"sort"
)

// List represents a strongly typed list of objects that can be accessed by index
// Provides methods to search, sort, and manipulate lists
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.list-1?view=netframework-4.7.2
type List[T comparable] struct {
	items    []T
	capacity int
}

// Predicate represents a method that defines a set of criteria and determines whether the specified object meets those criteria
type Predicate[T any] func(T) bool

// Action represents a method that performs an action on the specified object
type Action[T any] func(T)

// Comparison represents a method that compares two objects of the same type
type Comparison[T any] func(T, T) int

// NewList creates a new instance of the List class that is empty and has the default initial capacity
func NewList[T comparable]() *List[T] {
	return &List[T]{
		items:    make([]T, 0),
		capacity: 0,
	}
}

// NewListWithCapacity creates a new instance of the List class that is empty and has the specified initial capacity
func NewListWithCapacity[T comparable](capacity int) *List[T] {
	return &List[T]{
		items:    make([]T, 0, capacity),
		capacity: capacity,
	}
}

// NewListFromSlice creates a new instance of the List class that contains elements copied from the specified slice
func NewListFromSlice[T comparable](items []T) *List[T] {
	newItems := make([]T, len(items))
	copy(newItems, items)
	return &List[T]{
		items:    newItems,
		capacity: len(items),
	}
}

// Count gets the number of elements contained in the List
func (l *List[T]) Count() int {
	return len(l.items)
}

// Capacity gets or sets the total number of elements the internal data structure can hold without resizing
func (l *List[T]) Capacity() int {
	return cap(l.items)
}

// Add adds an object to the end of the List
func (l *List[T]) Add(item T) {
	l.items = append(l.items, item)
}

// AddRange adds the elements of the specified collection to the end of the List
func (l *List[T]) AddRange(items []T) {
	l.items = append(l.items, items...)
}

// Insert inserts an element into the List at the specified index
func (l *List[T]) Insert(index int, item T) error {
	if index < 0 || index > len(l.items) {
		return fmt.Errorf("index %d is out of range", index)
	}
	
	// Expand slice if necessary
	l.items = append(l.items, *new(T))
	
	// Shift elements to the right
	copy(l.items[index+1:], l.items[index:])
	
	// Insert the new item
	l.items[index] = item
	
	return nil
}

// InsertRange inserts the elements of a collection into the List at the specified index
func (l *List[T]) InsertRange(index int, items []T) error {
	if index < 0 || index > len(l.items) {
		return fmt.Errorf("index %d is out of range", index)
	}
	
	if len(items) == 0 {
		return nil
	}
	
	// Expand slice
	oldLen := len(l.items)
	newLen := oldLen + len(items)
	
	// Ensure capacity
	if newLen > cap(l.items) {
		newSlice := make([]T, newLen, newLen*2)
		copy(newSlice, l.items)
		l.items = newSlice
	} else {
		l.items = l.items[:newLen]
	}
	
	// Shift existing elements to the right
	copy(l.items[index+len(items):], l.items[index:oldLen])
	
	// Insert new items
	copy(l.items[index:], items)
	
	return nil
}

// Remove removes the first occurrence of a specific object from the List
func (l *List[T]) Remove(item T) bool {
	for i, v := range l.items {
		if v == item {
			l.RemoveAt(i)
			return true
		}
	}
	return false
}

// RemoveAt removes the element at the specified index of the List
func (l *List[T]) RemoveAt(index int) error {
	if index < 0 || index >= len(l.items) {
		return fmt.Errorf("index %d is out of range", index)
	}
	
	copy(l.items[index:], l.items[index+1:])
	l.items = l.items[:len(l.items)-1]
	
	return nil
}

// RemoveAll removes all the elements that match the conditions defined by the specified predicate
func (l *List[T]) RemoveAll(predicate Predicate[T]) int {
	removed := 0
	for i := len(l.items) - 1; i >= 0; i-- {
		if predicate(l.items[i]) {
			l.RemoveAt(i)
			removed++
		}
	}
	return removed
}

// RemoveRange removes a range of elements from the List
func (l *List[T]) RemoveRange(index, count int) error {
	if index < 0 || index >= len(l.items) {
		return fmt.Errorf("index %d is out of range", index)
	}
	if count < 0 || index+count > len(l.items) {
		return fmt.Errorf("count %d is invalid", count)
	}
	
	copy(l.items[index:], l.items[index+count:])
	l.items = l.items[:len(l.items)-count]
	
	return nil
}

// Contains determines whether an element is in the List
func (l *List[T]) Contains(item T) bool {
	for _, v := range l.items {
		if v == item {
			return true
		}
	}
	return false
}

// IndexOf searches for the specified object and returns the zero-based index of the first occurrence
func (l *List[T]) IndexOf(item T) int {
	for i, v := range l.items {
		if v == item {
			return i
		}
	}
	return -1
}

// LastIndexOf searches for the specified object and returns the zero-based index of the last occurrence
func (l *List[T]) LastIndexOf(item T) int {
	for i := len(l.items) - 1; i >= 0; i-- {
		if l.items[i] == item {
			return i
		}
	}
	return -1
}

// Find searches for an element that matches the conditions defined by the specified predicate
func (l *List[T]) Find(predicate Predicate[T]) (T, bool) {
	for _, item := range l.items {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// FindAll retrieves all the elements that match the conditions defined by the specified predicate
func (l *List[T]) FindAll(predicate Predicate[T]) *List[T] {
	result := NewList[T]()
	for _, item := range l.items {
		if predicate(item) {
			result.Add(item)
		}
	}
	return result
}

// FindIndex searches for an element that matches the conditions defined by the specified predicate
func (l *List[T]) FindIndex(predicate Predicate[T]) int {
	for i, item := range l.items {
		if predicate(item) {
			return i
		}
	}
	return -1
}

// FindLast searches for an element that matches the conditions defined by the specified predicate
func (l *List[T]) FindLast(predicate Predicate[T]) (T, bool) {
	for i := len(l.items) - 1; i >= 0; i-- {
		if predicate(l.items[i]) {
			return l.items[i], true
		}
	}
	var zero T
	return zero, false
}

// Get returns the element at the specified index
func (l *List[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(l.items) {
		var zero T
		return zero, fmt.Errorf("index %d is out of range", index)
	}
	return l.items[index], nil
}

// Set sets the element at the specified index
func (l *List[T]) Set(index int, item T) error {
	if index < 0 || index >= len(l.items) {
		return fmt.Errorf("index %d is out of range", index)
	}
	l.items[index] = item
	return nil
}

// Sort sorts the elements in the entire List using the default comparer
func (l *List[T]) Sort(comparison Comparison[T]) {
	sort.Slice(l.items, func(i, j int) bool {
		return comparison(l.items[i], l.items[j]) < 0
	})
}

// Reverse reverses the order of the elements in the entire List
func (l *List[T]) Reverse() {
	for i, j := 0, len(l.items)-1; i < j; i, j = i+1, j-1 {
		l.items[i], l.items[j] = l.items[j], l.items[i]
	}
}

// ForEach performs the specified action on each element of the List
func (l *List[T]) ForEach(action Action[T]) {
	for _, item := range l.items {
		action(item)
	}
}

// ToArray copies the elements of the List to a new array
func (l *List[T]) ToArray() []T {
	result := make([]T, len(l.items))
	copy(result, l.items)
	return result
}

// Clear removes all elements from the List
func (l *List[T]) Clear() {
	l.items = l.items[:0]
}

// TrimExcess sets the capacity to the actual number of elements in the List
func (l *List[T]) TrimExcess() {
	if len(l.items) < cap(l.items) {
		newSlice := make([]T, len(l.items))
		copy(newSlice, l.items)
		l.items = newSlice
	}
}

// Exists determines whether the List contains elements that match the conditions defined by the specified predicate
func (l *List[T]) Exists(predicate Predicate[T]) bool {
	for _, item := range l.items {
		if predicate(item) {
			return true
		}
	}
	return false
}

// TrueForAll determines whether every element in the List matches the conditions defined by the specified predicate
func (l *List[T]) TrueForAll(predicate Predicate[T]) bool {
	for _, item := range l.items {
		if !predicate(item) {
			return false
		}
	}
	return true
}