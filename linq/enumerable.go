package linq

import (
	"fmt"
	"sort"
)

// Enumerable provides LINQ-style extension methods for slices
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.linq.enumerable?view=netframework-4.7.2

// Predicate represents a method that defines a set of criteria and determines whether the specified object meets those criteria
type Predicate[T any] func(T) bool

// Selector represents a transform function to apply to each element
type Selector[T, R any] func(T) R

// Comparer represents a method that compares two objects
type Comparer[T any] func(T, T) int

// Enumerable wraps a slice to provide LINQ-style methods
type Enumerable[T any] struct {
	items []T
}

// From creates an Enumerable from a slice
func From[T any](items []T) *Enumerable[T] {
	return &Enumerable[T]{items: items}
}

// Range generates a sequence of integral numbers within a specified range
func Range(start, count int) *Enumerable[int] {
	items := make([]int, count)
	for i := 0; i < count; i++ {
		items[i] = start + i
	}
	return &Enumerable[int]{items: items}
}

// Repeat generates a sequence that contains one repeated value
func Repeat[T any](element T, count int) *Enumerable[T] {
	items := make([]T, count)
	for i := 0; i < count; i++ {
		items[i] = element
	}
	return &Enumerable[T]{items: items}
}

// Empty returns an empty Enumerable
func Empty[T any]() *Enumerable[T] {
	return &Enumerable[T]{items: []T{}}
}

// ToSlice returns the underlying slice
func (e *Enumerable[T]) ToSlice() []T {
	result := make([]T, len(e.items))
	copy(result, e.items)
	return result
}

// Count returns the number of elements in the sequence
func (e *Enumerable[T]) Count() int {
	return len(e.items)
}

// CountWhere returns the number of elements in the sequence that satisfy a condition
func (e *Enumerable[T]) CountWhere(predicate Predicate[T]) int {
	count := 0
	for _, item := range e.items {
		if predicate(item) {
			count++
		}
	}
	return count
}

// Where filters a sequence of values based on a predicate
func (e *Enumerable[T]) Where(predicate Predicate[T]) *Enumerable[T] {
	var result []T
	for _, item := range e.items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return &Enumerable[T]{items: result}
}

// Select projects each element of a sequence into a new form
func Select[T, R any](e *Enumerable[T], selector Selector[T, R]) *Enumerable[R] {
	result := make([]R, len(e.items))
	for i, item := range e.items {
		result[i] = selector(item)
	}
	return &Enumerable[R]{items: result}
}

// SelectMany projects each element of a sequence to an Enumerable and flattens the resulting sequences into one sequence
func SelectMany[T, R any](e *Enumerable[T], selector Selector[T, []R]) *Enumerable[R] {
	var result []R
	for _, item := range e.items {
		selected := selector(item)
		result = append(result, selected...)
	}
	return &Enumerable[R]{items: result}
}

// First returns the first element of a sequence
func (e *Enumerable[T]) First() (T, bool) {
	if len(e.items) == 0 {
		var zero T
		return zero, false
	}
	return e.items[0], true
}

// FirstWhere returns the first element of a sequence that satisfies a specified condition
func (e *Enumerable[T]) FirstWhere(predicate Predicate[T]) (T, bool) {
	for _, item := range e.items {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// FirstOrDefault returns the first element of a sequence, or a default value if no element is found
func (e *Enumerable[T]) FirstOrDefault(defaultValue T) T {
	if len(e.items) == 0 {
		return defaultValue
	}
	return e.items[0]
}

// Last returns the last element of a sequence
func (e *Enumerable[T]) Last() (T, bool) {
	if len(e.items) == 0 {
		var zero T
		return zero, false
	}
	return e.items[len(e.items)-1], true
}

// LastWhere returns the last element of a sequence that satisfies a specified condition
func (e *Enumerable[T]) LastWhere(predicate Predicate[T]) (T, bool) {
	for i := len(e.items) - 1; i >= 0; i-- {
		if predicate(e.items[i]) {
			return e.items[i], true
		}
	}
	var zero T
	return zero, false
}

// Single returns the only element of a sequence, and throws an exception if there is not exactly one element
func (e *Enumerable[T]) Single() (T, error) {
	if len(e.items) == 0 {
		var zero T
		return zero, fmt.Errorf("sequence contains no elements")
	}
	if len(e.items) > 1 {
		var zero T
		return zero, fmt.Errorf("sequence contains more than one element")
	}
	return e.items[0], nil
}

// Any determines whether any element of a sequence satisfies a condition
func (e *Enumerable[T]) Any(predicate Predicate[T]) bool {
	for _, item := range e.items {
		if predicate(item) {
			return true
		}
	}
	return false
}

// AnyElements determines whether a sequence contains any elements
func (e *Enumerable[T]) AnyElements() bool {
	return len(e.items) > 0
}

// All determines whether all elements of a sequence satisfy a condition
func (e *Enumerable[T]) All(predicate Predicate[T]) bool {
	for _, item := range e.items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Contains determines whether a sequence contains a specified element
func (e *Enumerable[T]) Contains(value T) bool {
	for _, item := range e.items {
		if any(item) == any(value) {
			return true
		}
	}
	return false
}

// Distinct returns distinct elements from a sequence
func (e *Enumerable[T]) Distinct() *Enumerable[T] {
	seen := make(map[any]bool)
	var result []T
	
	for _, item := range e.items {
		key := any(item)
		if !seen[key] {
			seen[key] = true
			result = append(result, item)
		}
	}
	
	return &Enumerable[T]{items: result}
}

// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements
func (e *Enumerable[T]) Skip(count int) *Enumerable[T] {
	if count >= len(e.items) {
		return &Enumerable[T]{items: []T{}}
	}
	if count <= 0 {
		return &Enumerable[T]{items: e.items}
	}
	
	result := make([]T, len(e.items)-count)
	copy(result, e.items[count:])
	return &Enumerable[T]{items: result}
}

// Take returns a specified number of contiguous elements from the start of a sequence
func (e *Enumerable[T]) Take(count int) *Enumerable[T] {
	if count <= 0 {
		return &Enumerable[T]{items: []T{}}
	}
	if count >= len(e.items) {
		result := make([]T, len(e.items))
		copy(result, e.items)
		return &Enumerable[T]{items: result}
	}
	
	result := make([]T, count)
	copy(result, e.items[:count])
	return &Enumerable[T]{items: result}
}

// SkipWhile bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements
func (e *Enumerable[T]) SkipWhile(predicate Predicate[T]) *Enumerable[T] {
	startIndex := 0
	for i, item := range e.items {
		if !predicate(item) {
			startIndex = i
			break
		}
		if i == len(e.items)-1 {
			// All elements satisfy the predicate
			return &Enumerable[T]{items: []T{}}
		}
	}
	
	result := make([]T, len(e.items)-startIndex)
	copy(result, e.items[startIndex:])
	return &Enumerable[T]{items: result}
}

// TakeWhile returns elements from a sequence as long as a specified condition is true
func (e *Enumerable[T]) TakeWhile(predicate Predicate[T]) *Enumerable[T] {
	var result []T
	for _, item := range e.items {
		if !predicate(item) {
			break
		}
		result = append(result, item)
	}
	return &Enumerable[T]{items: result}
}

// Reverse inverts the order of the elements in a sequence
func (e *Enumerable[T]) Reverse() *Enumerable[T] {
	result := make([]T, len(e.items))
	for i, item := range e.items {
		result[len(e.items)-1-i] = item
	}
	return &Enumerable[T]{items: result}
}

// OrderBy sorts the elements of a sequence in ascending order according to a key
func OrderBy[T any, K any](e *Enumerable[T], keySelector Selector[T, K], comparer Comparer[K]) *Enumerable[T] {
	result := make([]T, len(e.items))
	copy(result, e.items)
	
	sort.Slice(result, func(i, j int) bool {
		keyI := keySelector(result[i])
		keyJ := keySelector(result[j])
		return comparer(keyI, keyJ) < 0
	})
	
	return &Enumerable[T]{items: result}
}

// OrderByDescending sorts the elements of a sequence in descending order according to a key
func OrderByDescending[T any, K any](e *Enumerable[T], keySelector Selector[T, K], comparer Comparer[K]) *Enumerable[T] {
	result := make([]T, len(e.items))
	copy(result, e.items)
	
	sort.Slice(result, func(i, j int) bool {
		keyI := keySelector(result[i])
		keyJ := keySelector(result[j])
		return comparer(keyI, keyJ) > 0
	})
	
	return &Enumerable[T]{items: result}
}

// GroupBy groups the elements of a sequence according to a specified key selector function
func GroupBy[T any, K comparable](e *Enumerable[T], keySelector Selector[T, K]) map[K][]T {
	groups := make(map[K][]T)
	
	for _, item := range e.items {
		key := keySelector(item)
		groups[key] = append(groups[key], item)
	}
	
	return groups
}

// Concat concatenates two sequences
func (e *Enumerable[T]) Concat(other *Enumerable[T]) *Enumerable[T] {
	result := make([]T, len(e.items)+len(other.items))
	copy(result, e.items)
	copy(result[len(e.items):], other.items)
	return &Enumerable[T]{items: result}
}

// Union produces the set union of two sequences
func (e *Enumerable[T]) Union(other *Enumerable[T]) *Enumerable[T] {
	return e.Concat(other).Distinct()
}

// Intersect produces the set intersection of two sequences
func (e *Enumerable[T]) Intersect(other *Enumerable[T]) *Enumerable[T] {
	otherSet := make(map[any]bool)
	for _, item := range other.items {
		otherSet[any(item)] = true
	}
	
	var result []T
	seen := make(map[any]bool)
	
	for _, item := range e.items {
		key := any(item)
		if otherSet[key] && !seen[key] {
			seen[key] = true
			result = append(result, item)
		}
	}
	
	return &Enumerable[T]{items: result}
}

// Except produces the set difference of two sequences
func (e *Enumerable[T]) Except(other *Enumerable[T]) *Enumerable[T] {
	otherSet := make(map[any]bool)
	for _, item := range other.items {
		otherSet[any(item)] = true
	}
	
	var result []T
	seen := make(map[any]bool)
	
	for _, item := range e.items {
		key := any(item)
		if !otherSet[key] && !seen[key] {
			seen[key] = true
			result = append(result, item)
		}
	}
	
	return &Enumerable[T]{items: result}
}

// Aggregate applies an accumulator function over a sequence
func Aggregate[T, R any](e *Enumerable[T], seed R, func_ func(R, T) R) R {
	result := seed
	for _, item := range e.items {
		result = func_(result, item)
	}
	return result
}

// ForEach performs the specified action on each element of the sequence
func (e *Enumerable[T]) ForEach(action func(T)) {
	for _, item := range e.items {
		action(item)
	}
}