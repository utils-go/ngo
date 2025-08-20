package converter

import (
	"testing"
)

func TestToInt32(t *testing.T) {
	result, err := ToInt32(42)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
	
	result, err = ToInt32("42")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
}

func TestToBoolean(t *testing.T) {
	result, err := ToBoolean(true)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !result {
		t.Error("Expected true")
	}
	
	result, err = ToBoolean("true")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !result {
		t.Error("Expected true")
	}
}

func TestToString(t *testing.T) {
	result := ToString(42)
	if result != "42" {
		t.Errorf("Expected '42', got '%s'", result)
	}
	
	result = ToString(true)
	if result != "True" {
		t.Errorf("Expected 'True', got '%s'", result)
	}
}

func TestToBase64String(t *testing.T) {
	data := []byte("Hello")
	result := ToBase64String(data)
	expected := "SGVsbG8="
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFromBase64String(t *testing.T) {
	base64Str := "SGVsbG8="
	result, err := FromBase64String(base64Str)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := "Hello"
	if string(result) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, string(result))
	}
}