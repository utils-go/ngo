package generic

import (
	"fmt"
)

// KeyValuePair represents a key/value pair
type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

// Dictionary represents a collection of keys and values equivalent to System.Collections.Generic.Dictionary<TKey,TValue> in .NET
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.dictionary-2?view=netframework-4.7.2
type Dictionary[K comparable, V any] struct {
	items map[K]V
}

// NewDictionary creates a new Dictionary instance
func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{
		items: make(map[K]V),
	}
}

// NewDictionaryWithCapacity creates a new Dictionary instance with the specified initial capacity
func NewDictionaryWithCapacity[K comparable, V any](capacity int) *Dictionary[K, V] {
	return &Dictionary[K, V]{
		items: make(map[K]V, capacity),
	}
}

// NewDictionaryFromMap creates a new Dictionary instance from an existing map
func NewDictionaryFromMap[K comparable, V any](m map[K]V) *Dictionary[K, V] {
	dict := &Dictionary[K, V]{
		items: make(map[K]V, len(m)),
	}
	for k, v := range m {
		dict.items[k] = v
	}
	return dict
}

// Count gets the number of key/value pairs contained in the Dictionary
func (d *Dictionary[K, V]) Count() int {
	return len(d.items)
}

// IsEmpty gets a value indicating whether the Dictionary is empty
func (d *Dictionary[K, V]) IsEmpty() bool {
	return len(d.items) == 0
}

// Add adds the specified key and value to the dictionary
func (d *Dictionary[K, V]) Add(key K, value V) {
	if _, exists := d.items[key]; exists {
		panic(fmt.Sprintf("key already exists: %v", key))
	}
	d.items[key] = value
}

// Set sets the value for the specified key (adds if not exists, updates if exists)
func (d *Dictionary[K, V]) Set(key K, value V) {
	d.items[key] = value
}

// Get gets the value associated with the specified key
func (d *Dictionary[K, V]) Get(key K) (V, bool) {
	value, exists := d.items[key]
	return value, exists
}

// GetValue gets the value associated with the specified key, panics if key doesn't exist
func (d *Dictionary[K, V]) GetValue(key K) V {
	if value, exists := d.items[key]; exists {
		return value
	}
	panic(fmt.Sprintf("key not found: %v", key))
}

// TryGetValue gets the value associated with the specified key
func (d *Dictionary[K, V]) TryGetValue(key K) (V, bool) {
	value, exists := d.items[key]
	return value, exists
}

// ContainsKey determines whether the Dictionary contains the specified key
func (d *Dictionary[K, V]) ContainsKey(key K) bool {
	_, exists := d.items[key]
	return exists
}

// ContainsValue determines whether the Dictionary contains a specific value
func (d *Dictionary[K, V]) ContainsValue(value V) bool {
	for _, v := range d.items {
		if any(v) == any(value) {
			return true
		}
	}
	return false
}

// Remove removes the value with the specified key from the Dictionary
func (d *Dictionary[K, V]) Remove(key K) bool {
	if _, exists := d.items[key]; exists {
		delete(d.items, key)
		return true
	}
	return false
}

// Clear removes all keys and values from the Dictionary
func (d *Dictionary[K, V]) Clear() {
	d.items = make(map[K]V)
}

// Keys gets a collection containing the keys in the Dictionary
func (d *Dictionary[K, V]) Keys() []K {
	keys := make([]K, 0, len(d.items))
	for k := range d.items {
		keys = append(keys, k)
	}
	return keys
}

// Values gets a collection containing the values in the Dictionary
func (d *Dictionary[K, V]) Values() []V {
	values := make([]V, 0, len(d.items))
	for _, v := range d.items {
		values = append(values, v)
	}
	return values
}

// KeyValuePairs gets a collection containing the key/value pairs in the Dictionary
func (d *Dictionary[K, V]) KeyValuePairs() []KeyValuePair[K, V] {
	pairs := make([]KeyValuePair[K, V], 0, len(d.items))
	for k, v := range d.items {
		pairs = append(pairs, KeyValuePair[K, V]{Key: k, Value: v})
	}
	return pairs
}

// ForEach performs the specified action on each key/value pair of the Dictionary
func (d *Dictionary[K, V]) ForEach(action func(K, V)) {
	for k, v := range d.items {
		action(k, v)
	}
}

// Where filters the Dictionary based on a predicate
func (d *Dictionary[K, V]) Where(predicate func(K, V) bool) *Dictionary[K, V] {
	result := NewDictionary[K, V]()
	for k, v := range d.items {
		if predicate(k, v) {
			result.Set(k, v)
		}
	}
	return result
}

// Select projects each key/value pair of the Dictionary into a new form
func (d *Dictionary[K, V]) Select(selector func(K, V) KeyValuePair[K, V]) *Dictionary[K, V] {
	result := NewDictionary[K, V]()
	for k, v := range d.items {
		pair := selector(k, v)
		result.Set(pair.Key, pair.Value)
	}
	return result
}

// Any determines whether any key/value pair satisfies a condition
func (d *Dictionary[K, V]) Any(predicate func(K, V) bool) bool {
	for k, v := range d.items {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

// All determines whether all key/value pairs satisfy a condition
func (d *Dictionary[K, V]) All(predicate func(K, V) bool) bool {
	for k, v := range d.items {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

// First returns the first key/value pair that satisfies a condition
func (d *Dictionary[K, V]) First(predicate func(K, V) bool) (KeyValuePair[K, V], bool) {
	for k, v := range d.items {
		if predicate(k, v) {
			return KeyValuePair[K, V]{Key: k, Value: v}, true
		}
	}
	var zero KeyValuePair[K, V]
	return zero, false
}

// ToMap converts the Dictionary to a Go map
func (d *Dictionary[K, V]) ToMap() map[K]V {
	result := make(map[K]V, len(d.items))
	for k, v := range d.items {
		result[k] = v
	}
	return result
}

// Clone creates a shallow copy of the Dictionary
func (d *Dictionary[K, V]) Clone() *Dictionary[K, V] {
	result := NewDictionaryWithCapacity[K, V](len(d.items))
	for k, v := range d.items {
		result.items[k] = v
	}
	return result
}

// Equals determines whether this Dictionary and another Dictionary contain the same key/value pairs
func (d *Dictionary[K, V]) Equals(other *Dictionary[K, V]) bool {
	if other == nil || d.Count() != other.Count() {
		return false
	}
	
	for k, v := range d.items {
		if otherValue, exists := other.items[k]; !exists || any(v) != any(otherValue) {
			return false
		}
	}
	return true
}

// String returns a string representation of the Dictionary
func (d *Dictionary[K, V]) String() string {
	if d.IsEmpty() {
		return "{}"
	}
	
	result := "{"
	first := true
	for k, v := range d.items {
		if !first {
			result += ", "
		}
		result += fmt.Sprintf("%v: %v", k, v)
		first = false
	}
	result += "}"
	return result
}