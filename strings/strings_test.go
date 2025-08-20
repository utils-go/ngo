package strings

import (
	"testing"
)

func TestString_Basic(t *testing.T) {
	s := NewString("Hello World")
	
	if s.Length() != 11 {
		t.Errorf("Expected length 11, got %d", s.Length())
	}
	
	if s.Value() != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", s.Value())
	}
}

func TestString_Split(t *testing.T) {
	s := NewString("apple,banana,cherry")
	parts := s.Split(",")
	
	expected := []string{"apple", "banana", "cherry"}
	if len(parts) != len(expected) {
		t.Errorf("Expected %d parts, got %d", len(expected), len(parts))
	}
	
	for i, part := range parts {
		if part != expected[i] {
			t.Errorf("Expected '%s', got '%s'", expected[i], part)
		}
	}
}

func TestString_Replace(t *testing.T) {
	s := NewString("Hello World")
	result := s.Replace("World", "Go")
	
	if result.Value() != "Hello Go" {
		t.Errorf("Expected 'Hello Go', got '%s'", result.Value())
	}
}

func TestString_Trim(t *testing.T) {
	s := NewString("  Hello World  ")
	result := s.Trim()
	
	if result.Value() != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", result.Value())
	}
}

func TestString_Contains(t *testing.T) {
	s := NewString("Hello World")
	
	if !s.Contains("World") {
		t.Error("Expected Contains('World') to be true")
	}
	
	if s.Contains("xyz") {
		t.Error("Expected Contains('xyz') to be false")
	}
}

func TestString_StartsWith_EndsWith(t *testing.T) {
	s := NewString("Hello World")
	
	if !s.StartsWith("Hello") {
		t.Error("Expected StartsWith('Hello') to be true")
	}
	
	if !s.EndsWith("World") {
		t.Error("Expected EndsWith('World') to be true")
	}
}

func TestString_IndexOf(t *testing.T) {
	s := NewString("Hello World")
	
	index := s.IndexOf("World")
	if index != 6 {
		t.Errorf("Expected index 6, got %d", index)
	}
	
	index = s.IndexOf("xyz")
	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

func TestString_ToUpper_ToLower(t *testing.T) {
	s := NewString("Hello World")
	
	upper := s.ToUpper()
	if upper.Value() != "HELLO WORLD" {
		t.Errorf("Expected 'HELLO WORLD', got '%s'", upper.Value())
	}
	
	lower := s.ToLower()
	if lower.Value() != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", lower.Value())
	}
}

func TestString_Substring(t *testing.T) {
	s := NewString("Hello World")
	
	sub := s.Substring(6)
	if sub.Value() != "World" {
		t.Errorf("Expected 'World', got '%s'", sub.Value())
	}
	
	sub = s.SubstringWithLength(0, 5)
	if sub.Value() != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", sub.Value())
	}
}

func TestString_Format(t *testing.T) {
	result := Format("Hello {0}, you are {1} years old", "John", 25)
	expected := "Hello John, you are 25 years old"
	
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestString_PadLeft_PadRight(t *testing.T) {
	s := NewString("Hello")
	
	padded := s.PadLeft(10)
	if padded.Value() != "     Hello" {
		t.Errorf("Expected '     Hello', got '%s'", padded.Value())
	}
	
	padded = s.PadRight(10)
	if padded.Value() != "Hello     " {
		t.Errorf("Expected 'Hello     ', got '%s'", padded.Value())
	}
}

func TestString_IsNullOrEmpty(t *testing.T) {
	if !IsNullOrEmpty(nil) {
		t.Error("Expected IsNullOrEmpty(nil) to be true")
	}
	
	if !IsNullOrEmpty(NewString("")) {
		t.Error("Expected IsNullOrEmpty('') to be true")
	}
	
	if IsNullOrEmpty(NewString("Hello")) {
		t.Error("Expected IsNullOrEmpty('Hello') to be false")
	}
}

func TestString_IsNullOrWhiteSpace(t *testing.T) {
	if !IsNullOrWhiteSpace(nil) {
		t.Error("Expected IsNullOrWhiteSpace(nil) to be true")
	}
	
	if !IsNullOrWhiteSpace(NewString("")) {
		t.Error("Expected IsNullOrWhiteSpace('') to be true")
	}
	
	if !IsNullOrWhiteSpace(NewString("   ")) {
		t.Error("Expected IsNullOrWhiteSpace('   ') to be true")
	}
	
	if IsNullOrWhiteSpace(NewString("Hello")) {
		t.Error("Expected IsNullOrWhiteSpace('Hello') to be false")
	}
}