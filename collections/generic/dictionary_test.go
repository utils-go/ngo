package generic

import (
	"testing"
)

func TestNewDictionary(t *testing.T) {
	dict := NewDictionary[string, int]()
	
	if dict.Count() != 0 {
		t.Errorf("Expected count 0, got %d", dict.Count())
	}
	
	if !dict.IsEmpty() {
		t.Error("Expected dictionary to be empty")
	}
}

func TestAdd(t *testing.T) {
	dict := NewDictionary[string, int]()
	
	dict.Add("one", 1)
	dict.Add("two", 2)
	
	if dict.Count() != 2 {
		t.Errorf("Expected count 2, got %d", dict.Count())
	}
	
	if value, exists := dict.Get("one"); !exists || value != 1 {
		t.Errorf("Expected value 1 for key 'one', got %d (exists: %t)", value, exists)
	}
}

func TestSet(t *testing.T) {
	dict := NewDictionary[string, int]()
	
	dict.Set("key", 1)
	dict.Set("key", 2) // Should update existing key
	
	if dict.Count() != 1 {
		t.Errorf("Expected count 1, got %d", dict.Count())
	}
	
	if value := dict.GetValue("key"); value != 2 {
		t.Errorf("Expected value 2, got %d", value)
	}
}

func TestContainsKey(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("exists", 1)
	
	if !dict.ContainsKey("exists") {
		t.Error("Expected ContainsKey('exists') to be true")
	}
	
	if dict.ContainsKey("notexists") {
		t.Error("Expected ContainsKey('notexists') to be false")
	}
}

func TestContainsValue(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("key1", 1)
	dict.Add("key2", 2)
	
	if !dict.ContainsValue(1) {
		t.Error("Expected ContainsValue(1) to be true")
	}
	
	if dict.ContainsValue(3) {
		t.Error("Expected ContainsValue(3) to be false")
	}
}

func TestRemove(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("key1", 1)
	dict.Add("key2", 2)
	
	removed := dict.Remove("key1")
	if !removed {
		t.Error("Expected Remove('key1') to return true")
	}
	
	if dict.Count() != 1 {
		t.Errorf("Expected count 1 after removal, got %d", dict.Count())
	}
	
	if dict.ContainsKey("key1") {
		t.Error("Expected key1 to be removed")
	}
	
	removed = dict.Remove("nonexistent")
	if removed {
		t.Error("Expected Remove('nonexistent') to return false")
	}
}

func TestKeys(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("key1", 1)
	dict.Add("key2", 2)
	
	keys := dict.Keys()
	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}
	
	// Check that both keys are present (order doesn't matter)
	keyMap := make(map[string]bool)
	for _, key := range keys {
		keyMap[key] = true
	}
	
	if !keyMap["key1"] || !keyMap["key2"] {
		t.Error("Expected keys 'key1' and 'key2'")
	}
}

func TestValues(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("key1", 1)
	dict.Add("key2", 2)
	
	values := dict.Values()
	if len(values) != 2 {
		t.Errorf("Expected 2 values, got %d", len(values))
	}
	
	// Check that both values are present (order doesn't matter)
	valueMap := make(map[int]bool)
	for _, value := range values {
		valueMap[value] = true
	}
	
	if !valueMap[1] || !valueMap[2] {
		t.Error("Expected values 1 and 2")
	}
}

func TestClear(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("key1", 1)
	dict.Add("key2", 2)
	
	dict.Clear()
	
	if dict.Count() != 0 {
		t.Errorf("Expected count 0 after clear, got %d", dict.Count())
	}
	
	if !dict.IsEmpty() {
		t.Error("Expected dictionary to be empty after clear")
	}
}

func TestWhere(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("one", 1)
	dict.Add("two", 2)
	dict.Add("three", 3)
	dict.Add("four", 4)
	
	// Filter even values
	filtered := dict.Where(func(k string, v int) bool {
		return v%2 == 0
	})
	
	if filtered.Count() != 2 {
		t.Errorf("Expected 2 filtered items, got %d", filtered.Count())
	}
	
	if !filtered.ContainsKey("two") || !filtered.ContainsKey("four") {
		t.Error("Expected filtered dictionary to contain 'two' and 'four'")
	}
}

func TestAny(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("one", 1)
	dict.Add("two", 2)
	
	hasEven := dict.Any(func(k string, v int) bool {
		return v%2 == 0
	})
	
	if !hasEven {
		t.Error("Expected Any to return true for even values")
	}
	
	hasLarge := dict.Any(func(k string, v int) bool {
		return v > 10
	})
	
	if hasLarge {
		t.Error("Expected Any to return false for values > 10")
	}
}

func TestAll(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("one", 1)
	dict.Add("two", 2)
	
	allPositive := dict.All(func(k string, v int) bool {
		return v > 0
	})
	
	if !allPositive {
		t.Error("Expected All to return true for positive values")
	}
	
	allEven := dict.All(func(k string, v int) bool {
		return v%2 == 0
	})
	
	if allEven {
		t.Error("Expected All to return false for even values")
	}
}

func TestClone(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Add("one", 1)
	dict.Add("two", 2)
	
	cloned := dict.Clone()
	
	if cloned.Count() != dict.Count() {
		t.Errorf("Expected cloned count %d, got %d", dict.Count(), cloned.Count())
	}
	
	if !cloned.ContainsKey("one") || !cloned.ContainsKey("two") {
		t.Error("Expected cloned dictionary to contain all keys")
	}
	
	// Modify original, clone should be unaffected
	dict.Add("three", 3)
	
	if cloned.Count() != 2 {
		t.Error("Expected clone to be unaffected by changes to original")
	}
}

func TestEquals(t *testing.T) {
	dict1 := NewDictionary[string, int]()
	dict1.Add("one", 1)
	dict1.Add("two", 2)
	
	dict2 := NewDictionary[string, int]()
	dict2.Add("one", 1)
	dict2.Add("two", 2)
	
	dict3 := NewDictionary[string, int]()
	dict3.Add("one", 1)
	dict3.Add("two", 3) // Different value
	
	if !dict1.Equals(dict2) {
		t.Error("Expected dict1 to equal dict2")
	}
	
	if dict1.Equals(dict3) {
		t.Error("Expected dict1 not to equal dict3")
	}
}