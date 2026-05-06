package generic

import (
	"sort"
)

// SortedDictionaryEntry represents a key/value pair in a SortedDictionary.
type SortedDictionaryEntry[K comparable, V any] struct {
	Key   K
	Value V
}

// SortedDictionary represents a collection of key/value pairs that are sorted on the key.
// Equivalent to System.Collections.Generic.SortedDictionary<TKey,TValue> in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.sorteddictionary-2?view=netframework-4.7.2
type SortedDictionary[K comparable, V any] struct {
	entries []SortedDictionaryEntry[K, V]
	dict    map[K]V
	less    func(K, K) bool
}

// NewSortedDictionary creates a new SortedDictionary with the specified comparator.
func NewSortedDictionary[K comparable, V any](less func(K, K) bool) *SortedDictionary[K, V] {
	return &SortedDictionary[K, V]{
		entries: make([]SortedDictionaryEntry[K, V], 0),
		dict:    make(map[K]V),
		less:    less,
	}
}

// NewSortedDictionaryWithCapacity creates a new SortedDictionary with initial capacity.
func NewSortedDictionaryWithCapacity[K comparable, V any](capacity int, less func(K, K) bool) *SortedDictionary[K, V] {
	return &SortedDictionary[K, V]{
		entries: make([]SortedDictionaryEntry[K, V], 0, capacity),
		dict:    make(map[K]V, capacity),
		less:    less,
	}
}

// Count gets the number of key/value pairs.
func (sd *SortedDictionary[K, V]) Count() int {
	return len(sd.entries)
}

// Add adds an element with the specified key and value.
func (sd *SortedDictionary[K, V]) Add(key K, value V) {
	if _, exists := sd.dict[key]; exists {
		sd.Set(key, value)
		return
	}

	entry := SortedDictionaryEntry[K, V]{Key: key, Value: value}

	// Find insertion position
	idx := sort.Search(len(sd.entries), func(i int) bool {
		return sd.less(key, sd.entries[i].Key)
	})

	sd.entries = append(sd.entries, entry)
	copy(sd.entries[idx+1:], sd.entries[idx:])
	sd.entries[idx] = entry

	sd.dict[key] = value
}

// Set sets the value for the specified key.
func (sd *SortedDictionary[K, V]) Set(key K, value V) {
	if _, exists := sd.dict[key]; exists {
		for i, e := range sd.entries {
			if e.Key == key {
				sd.entries[i].Value = value
				break
			}
		}
		sd.dict[key] = value
	} else {
		sd.Add(key, value)
	}
}

// Get gets the value associated with the specified key.
func (sd *SortedDictionary[K, V]) Get(key K) (V, bool) {
	value, exists := sd.dict[key]
	return value, exists
}

// GetValue gets the value (panics if key not found).
func (sd *SortedDictionary[K, V]) GetValue(key K) V {
	if value, exists := sd.dict[key]; exists {
		return value
	}
	panic("key not found")
}

// TryGetValue gets the value associated with the specified key.
func (sd *SortedDictionary[K, V]) TryGetValue(key K) (V, bool) {
	return sd.Get(key)
}

// Remove removes the element with the specified key.
func (sd *SortedDictionary[K, V]) Remove(key K) bool {
	idx := -1
	for i, e := range sd.entries {
		if e.Key == key {
			idx = i
			break
		}
	}
	if idx < 0 {
		return false
	}

	sd.entries = append(sd.entries[:idx], sd.entries[idx+1:]...)
	delete(sd.dict, key)
	return true
}

// ContainsKey determines whether the dictionary contains a specific key.
func (sd *SortedDictionary[K, V]) ContainsKey(key K) bool {
	_, exists := sd.dict[key]
	return exists
}

// ContainsValue determines whether the dictionary contains a specific value.
func (sd *SortedDictionary[K, V]) ContainsValue(value V) bool {
	for _, e := range sd.entries {
		if any(e.Value) == any(value) {
			return true
		}
	}
	return false
}

// Keys gets a collection containing the keys.
func (sd *SortedDictionary[K, V]) Keys() []K {
	result := make([]K, len(sd.entries))
	for i, e := range sd.entries {
		result[i] = e.Key
	}
	return result
}

// Values gets a collection containing the values.
func (sd *SortedDictionary[K, V]) Values() []V {
	result := make([]V, len(sd.entries))
	for i, e := range sd.entries {
		result[i] = e.Value
	}
	return result
}

// Entries returns the sorted entries.
func (sd *SortedDictionary[K, V]) Entries() []SortedDictionaryEntry[K, V] {
	result := make([]SortedDictionaryEntry[K, V], len(sd.entries))
	copy(result, sd.entries)
	return result
}

// First returns the first key/value pair.
func (sd *SortedDictionary[K, V]) First() (SortedDictionaryEntry[K, V], bool) {
	if len(sd.entries) == 0 {
		var zero SortedDictionaryEntry[K, V]
		return zero, false
	}
	return sd.entries[0], true
}

// Last returns the last key/value pair.
func (sd *SortedDictionary[K, V]) Last() (SortedDictionaryEntry[K, V], bool) {
	if len(sd.entries) == 0 {
		var zero SortedDictionaryEntry[K, V]
		return zero, false
	}
	return sd.entries[len(sd.entries)-1], true
}

// Clear removes all elements.
func (sd *SortedDictionary[K, V]) Clear() {
	sd.entries = sd.entries[:0]
	sd.dict = make(map[K]V)
}

// ForEach performs the specified action on each key/value pair.
func (sd *SortedDictionary[K, V]) ForEach(action func(K, V)) {
	for _, e := range sd.entries {
		action(e.Key, e.Value)
	}
}

// TrimExcess sets the capacity to the actual number of elements.
func (sd *SortedDictionary[K, V]) TrimExcess() {
	if len(sd.entries) < cap(sd.entries) {
		newEntries := make([]SortedDictionaryEntry[K, V], len(sd.entries))
		copy(newEntries, sd.entries)
		sd.entries = newEntries
	}
}
