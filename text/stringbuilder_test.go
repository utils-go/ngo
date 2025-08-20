package text

import (
	"testing"
)

func TestNewStringBuilder(t *testing.T) {
	sb := NewStringBuilder()
	
	if sb.Length() != 0 {
		t.Errorf("Expected length 0, got %d", sb.Length())
	}
	
	if sb.ToString() != "" {
		t.Errorf("Expected empty string, got '%s'", sb.ToString())
	}
}

func TestAppendString(t *testing.T) {
	sb := NewStringBuilder()
	
	sb.AppendString("Hello").AppendString(" ").AppendString("World")
	
	expected := "Hello World"
	if sb.ToString() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, sb.ToString())
	}
	
	if sb.Length() != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), sb.Length())
	}
}

func TestAppendLine(t *testing.T) {
	sb := NewStringBuilder()
	
	sb.AppendLineString("Line 1").AppendLineString("Line 2")
	
	expected := "Line 1\nLine 2\n"
	if sb.ToString() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, sb.ToString())
	}
}

func TestAppendFormat(t *testing.T) {
	sb := NewStringBuilder()
	
	sb.AppendFormat("Hello %s, you are %d years old", "John", 25)
	
	expected := "Hello John, you are 25 years old"
	if sb.ToString() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, sb.ToString())
	}
}

func TestInsert(t *testing.T) {
	sb := NewStringBuilder()
	sb.AppendString("Hello World")
	
	sb.Insert(6, "Beautiful ")
	
	expected := "Hello Beautiful World"
	if sb.ToString() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, sb.ToString())
	}
}

func TestRemove(t *testing.T) {
	sb := NewStringBuilder()
	sb.AppendString("Hello Beautiful World")
	
	sb.Remove(6, 10) // Remove "Beautiful "
	
	expected := "Hello World"
	if sb.ToString() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, sb.ToString())
	}
}

func TestReplace(t *testing.T) {
	sb := NewStringBuilder()
	sb.AppendString("Hello World World")
	
	sb.Replace("World", "Go")
	
	expected := "Hello Go Go"
	if sb.ToString() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, sb.ToString())
	}
}

func TestClear(t *testing.T) {
	sb := NewStringBuilder()
	sb.AppendString("Hello World")
	
	sb.Clear()
	
	if sb.Length() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", sb.Length())
	}
	
	if sb.ToString() != "" {
		t.Errorf("Expected empty string after clear, got '%s'", sb.ToString())
	}
}

func TestEquals(t *testing.T) {
	sb1 := NewStringBuilder()
	sb1.AppendString("Hello World")
	
	sb2 := NewStringBuilder()
	sb2.AppendString("Hello World")
	
	sb3 := NewStringBuilder()
	sb3.AppendString("Hello Go")
	
	if !sb1.Equals(sb2) {
		t.Error("Expected sb1 to equal sb2")
	}
	
	if sb1.Equals(sb3) {
		t.Error("Expected sb1 not to equal sb3")
	}
}

func TestGetCharAt(t *testing.T) {
	sb := NewStringBuilder()
	sb.AppendString("Hello")
	
	if sb.GetCharAt(0) != 'H' {
		t.Errorf("Expected 'H', got '%c'", sb.GetCharAt(0))
	}
	
	if sb.GetCharAt(4) != 'o' {
		t.Errorf("Expected 'o', got '%c'", sb.GetCharAt(4))
	}
}

func TestSetCharAt(t *testing.T) {
	sb := NewStringBuilder()
	sb.AppendString("Hello")
	
	sb.SetCharAt(0, 'h')
	
	expected := "hello"
	if sb.ToString() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, sb.ToString())
	}
}