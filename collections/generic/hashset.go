package generic

// HashSet represents a set of values, equivalent to System.Collections.Generic.HashSet<T> in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.hashset-1?view=netframework-4.7.2
type HashSet[T comparable] struct {
	items map[T]struct{}
}

// NewHashSet creates a new empty HashSet.
func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		items: make(map[T]struct{}),
	}
}

// NewHashSetWithCapacity creates a new HashSet with the specified initial capacity.
func NewHashSetWithCapacity[T comparable](capacity int) *HashSet[T] {
	return &HashSet[T]{
		items: make(map[T]struct{}, capacity),
	}
}

// NewHashSetFromSlice creates a new HashSet containing elements from the given slice.
func NewHashSetFromSlice[T comparable](items []T) *HashSet[T] {
	hs := NewHashSetWithCapacity[T](len(items))
	for _, item := range items {
		hs.items[item] = struct{}{}
	}
	return hs
}

// Add adds an element to the HashSet. Returns true if added; false if already present.
func (hs *HashSet[T]) Add(item T) bool {
	if _, exists := hs.items[item]; exists {
		return false
	}
	hs.items[item] = struct{}{}
	return true
}

// Remove removes the specified element from the HashSet.
func (hs *HashSet[T]) Remove(item T) bool {
	if _, exists := hs.items[item]; exists {
		delete(hs.items, item)
		return true
	}
	return false
}

// Contains determines whether the HashSet contains the specified element.
func (hs *HashSet[T]) Contains(item T) bool {
	_, exists := hs.items[item]
	return exists
}

// Count gets the number of elements in the HashSet.
func (hs *HashSet[T]) Count() int {
	return len(hs.items)
}

// Clear removes all elements from the HashSet.
func (hs *HashSet[T]) Clear() {
	hs.items = make(map[T]struct{})
}

// ToSlice returns the elements of the HashSet as a slice.
func (hs *HashSet[T]) ToSlice() []T {
	result := make([]T, 0, len(hs.items))
	for item := range hs.items {
		result = append(result, item)
	}
	return result
}

// UnionWith modifies the HashSet to contain all elements present in itself and the other collection.
func (hs *HashSet[T]) UnionWith(other *HashSet[T]) {
	for item := range other.items {
		hs.items[item] = struct{}{}
	}
}

// IntersectWith modifies the HashSet to contain only elements present in both collections.
func (hs *HashSet[T]) IntersectWith(other *HashSet[T]) {
	for item := range hs.items {
		if !other.Contains(item) {
			delete(hs.items, item)
		}
	}
}

// ExceptWith removes all elements in the other collection from the HashSet.
func (hs *HashSet[T]) ExceptWith(other *HashSet[T]) {
	for item := range other.items {
		delete(hs.items, item)
	}
}

// SymmetricExceptWith modifies the HashSet to contain only elements present in either
// the HashSet or the other collection, but not both.
func (hs *HashSet[T]) SymmetricExceptWith(other *HashSet[T]) {
	for item := range other.items {
		if hs.Contains(item) {
			delete(hs.items, item)
		} else {
			hs.items[item] = struct{}{}
		}
	}
}

// IsSubsetOf determines whether the HashSet is a subset of the specified collection.
func (hs *HashSet[T]) IsSubsetOf(other *HashSet[T]) bool {
	if hs.Count() > other.Count() {
		return false
	}
	for item := range hs.items {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// IsSupersetOf determines whether the HashSet is a superset of the specified collection.
func (hs *HashSet[T]) IsSupersetOf(other *HashSet[T]) bool {
	if hs.Count() < other.Count() {
		return false
	}
	for item := range other.items {
		if !hs.Contains(item) {
			return false
		}
	}
	return true
}

// IsProperSubsetOf determines whether the HashSet is a proper subset of the specified collection.
func (hs *HashSet[T]) IsProperSubsetOf(other *HashSet[T]) bool {
	return hs.IsSubsetOf(other) && hs.Count() < other.Count()
}

// IsProperSupersetOf determines whether the HashSet is a proper superset of the specified collection.
func (hs *HashSet[T]) IsProperSupersetOf(other *HashSet[T]) bool {
	return hs.IsSupersetOf(other) && hs.Count() > other.Count()
}

// Overlaps determines whether the HashSet and the specified collection share common elements.
func (hs *HashSet[T]) Overlaps(other *HashSet[T]) bool {
	for item := range other.items {
		if hs.Contains(item) {
			return true
		}
	}
	return false
}

// SetEquals determines whether the HashSet and the specified collection contain the same elements.
func (hs *HashSet[T]) SetEquals(other *HashSet[T]) bool {
	if hs.Count() != other.Count() {
		return false
	}
	return hs.IsSubsetOf(other)
}

// ForEach performs the specified action on each element.
func (hs *HashSet[T]) ForEach(action func(T)) {
	for item := range hs.items {
		action(item)
	}
}
