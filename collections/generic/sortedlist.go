package generic

import (
	"sort"
)

// SortedList represents a collection of key/value pairs that are sorted by key.
// Equivalent to System.Collections.Generic.SortedList<TKey,TValue> in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.sortedlist-2?view=netframework-4.7.2
type SortedList[K comparable, V any] struct {
	keys   []K
	values []V
	dict   map[K]V
	less   func(K, K) bool
}

// NewSortedList creates a new SortedList with the default comparator.
func NewSortedList[K comparable, V any](less func(K, K) bool) *SortedList[K, V] {
	return &SortedList[K, V]{
		keys: make([]K, 0),
		values: make([]V, 0),
		dict:  make(map[K]V),
		less:  less,
	}
}

// NewSortedListWithCapacity creates a new SortedList with a specified initial capacity.
func NewSortedListWithCapacity[K comparable, V any](capacity int, less func(K, K) bool) *SortedList[K, V] {
	return &SortedList[K, V]{
		keys:   make([]K, 0, capacity),
		values: make([]V, 0, capacity),
		dict:   make(map[K]V, capacity),
		less:   less,
	}
}

// Count gets the number of key/value pairs in the SortedList.
func (sl *SortedList[K, V]) Count() int {
	return len(sl.keys)
}

// Add adds an element with the specified key and value.
func (sl *SortedList[K, V]) Add(key K, value V) {
	if _, exists := sl.dict[key]; exists {
		sl.Set(key, value)
		return
	}

	// Find insertion position
	idx := sort.Search(len(sl.keys), func(i int) bool {
		return sl.less(key, sl.keys[i])
	})

	// Insert into keys slice
	sl.keys = append(sl.keys, *new(K))
	copy(sl.keys[idx+1:], sl.keys[idx:])
	sl.keys[idx] = key

	// Insert into values slice
	sl.values = append(sl.values, *new(V))
	copy(sl.values[idx+1:], sl.values[idx:])
	sl.values[idx] = value

	sl.dict[key] = value
}

// Set sets the value for the specified key.
func (sl *SortedList[K, V]) Set(key K, value V) {
	if _, exists := sl.dict[key]; exists {
		// Find index and update
		for i, k := range sl.keys {
			if k == key {
				sl.values[i] = value
				break
			}
		}
		sl.dict[key] = value
	} else {
		sl.Add(key, value)
	}
}

// Get gets the value associated with the specified key.
func (sl *SortedList[K, V]) Get(key K) (V, bool) {
	value, exists := sl.dict[key]
	return value, exists
}

// Remove removes the element with the specified key.
func (sl *SortedList[K, V]) Remove(key K) bool {
	idx := -1
	for i, k := range sl.keys {
		if k == key {
			idx = i
			break
		}
	}
	if idx < 0 {
		return false
	}

	sl.keys = append(sl.keys[:idx], sl.keys[idx+1:]...)
	sl.values = append(sl.values[:idx], sl.values[idx+1:]...)
	delete(sl.dict, key)
	return true
}

// RemoveAt removes the element at the specified index.
func (sl *SortedList[K, V]) RemoveAt(index int) bool {
	if index < 0 || index >= len(sl.keys) {
		return false
	}
	key := sl.keys[index]
	return sl.Remove(key)
}

// ContainsKey determines whether the SortedList contains a specific key.
func (sl *SortedList[K, V]) ContainsKey(key K) bool {
	_, exists := sl.dict[key]
	return exists
}

// ContainsValue determines whether the SortedList contains a specific value.
func (sl *SortedList[K, V]) ContainsValue(value V) bool {
	for _, v := range sl.values {
		if any(v) == any(value) {
			return true
		}
	}
	return false
}

// IndexOfKey returns the zero-based index of the specified key.
func (sl *SortedList[K, V]) IndexOfKey(key K) int {
	for i, k := range sl.keys {
		if k == key {
			return i
		}
	}
	return -1
}

// IndexOfValue returns the zero-based index of the first occurrence of the specified value.
func (sl *SortedList[K, V]) IndexOfValue(value V) int {
	for i, v := range sl.values {
		if any(v) == any(value) {
			return i
		}
	}
	return -1
}

// Keys gets a collection containing the keys in the SortedList.
func (sl *SortedList[K, V]) Keys() []K {
	result := make([]K, len(sl.keys))
	copy(result, sl.keys)
	return result
}

// Values gets a collection containing the values in the SortedList.
func (sl *SortedList[K, V]) Values() []V {
	result := make([]V, len(sl.values))
	copy(result, sl.values)
	return result
}

// GetKeyAtIndex gets the key at the specified index.
func (sl *SortedList[K, V]) GetKeyAtIndex(index int) (K, bool) {
	if index < 0 || index >= len(sl.keys) {
		var zero K
		return zero, false
	}
	return sl.keys[index], true
}

// GetValueAtIndex gets the value at the specified index.
func (sl *SortedList[K, V]) GetValueAtIndex(index int) (V, bool) {
	if index < 0 || index >= len(sl.values) {
		var zero V
		return zero, false
	}
	return sl.values[index], true
}

// Clear removes all elements.
func (sl *SortedList[K, V]) Clear() {
	sl.keys = sl.keys[:0]
	sl.values = sl.values[:0]
	sl.dict = make(map[K]V)
}

// ForEach performs the specified action on each key/value pair.
func (sl *SortedList[K, V]) ForEach(action func(K, V)) {
	for i, k := range sl.keys {
		action(k, sl.values[i])
	}
}

// TrimExcess sets the capacity to the actual number of elements.
func (sl *SortedList[K, V]) TrimExcess() {
	if len(sl.keys) < cap(sl.keys) {
		newKeys := make([]K, len(sl.keys))
		copy(newKeys, sl.keys)
		sl.keys = newKeys

		newValues := make([]V, len(sl.values))
		copy(newValues, sl.values)
		sl.values = newValues
	}
}
