package regex

import (
	"testing"
)

func TestCompile(t *testing.T) {
	r, err := Compile(`\d+`)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}
	if r == nil {
		t.Fatal("expected non-nil Regex")
	}
}

func TestIsMatch(t *testing.T) {
	r := MustCompile(`\d+`)
	if !r.IsMatch("abc123def") {
		t.Error("expected match")
	}
	if r.IsMatch("abcdef") {
		t.Error("expected no match")
	}
}

func TestMatch(t *testing.T) {
	r := MustCompile(`\d+`)
	m := r.Match("abc123def456")
	if m == nil {
		t.Fatal("expected match")
	}
	if m.Value() != "123" {
		t.Errorf("expected '123', got %q", m.Value())
	}
	if !m.Success() {
		t.Error("expected success")
	}
}

func TestMatches(t *testing.T) {
	r := MustCompile(`\d+`)
	matches := r.Matches("abc123def456ghi789")
	if matches.Count() != 3 {
		t.Errorf("expected 3 matches, got %d", matches.Count())
	}
	if matches.Item(0).Value() != "123" {
		t.Errorf("expected '123', got %q", matches.Item(0).Value())
	}
	if matches.Item(1).Value() != "456" {
		t.Errorf("expected '456', got %q", matches.Item(1).Value())
	}
	if matches.Item(2).Value() != "789" {
		t.Errorf("expected '789', got %q", matches.Item(2).Value())
	}
}

func TestReplace(t *testing.T) {
	r := MustCompile(`\d+`)
	result := r.Replace("abc123def456", "#")
	if result != "abc#def#" {
		t.Errorf("expected 'abc#def#', got %q", result)
	}
}

func TestSplit(t *testing.T) {
	r := MustCompile(`\s+`)
	parts := r.Split("hello world foo bar")
	if len(parts) != 4 {
		t.Errorf("expected 4 parts, got %d: %v", len(parts), parts)
	}
}

func TestGroups(t *testing.T) {
	r := MustCompile(`(\w+)@(\w+)\.(\w+)`)
	m := r.Match("user@example.com")
	if m == nil {
		t.Fatal("expected match")
	}

	// Group 0 is the full match
	if m.Group(0).Value() != "user@example.com" {
		t.Errorf("expected full match, got %q", m.Group(0).Value())
	}
	if m.Group(1).Value() != "user" {
		t.Errorf("expected 'user', got %q", m.Group(1).Value())
	}
	if m.Group(2).Value() != "example" {
		t.Errorf("expected 'example', got %q", m.Group(2).Value())
	}
	if m.Group(3).Value() != "com" {
		t.Errorf("expected 'com', got %q", m.Group(3).Value())
	}

	// Groups() should have 4 entries (full match + 3 groups)
	if len(m.Groups()) != 4 {
		t.Errorf("expected 4 groups, got %d", len(m.Groups()))
	}
}

func TestNamedGroups(t *testing.T) {
	r := MustCompile(`(?P<name>\w+)@(?P<domain>\w+)\.(?P<tld>\w+)`)
	m := r.Match("admin@company.org")
	if m == nil {
		t.Fatal("expected match")
	}
	if m.Group("name").Value() != "admin" {
		t.Errorf("expected 'admin', got %q", m.Group("name").Value())
	}
	if m.Group("domain").Value() != "company" {
		t.Errorf("expected 'company', got %q", m.Group("domain").Value())
	}
}

func TestIgnoreCase(t *testing.T) {
	r := MustCompileWithOptions(`hello`, IgnoreCase)
	if !r.IsMatch("HELLO world") {
		t.Error("expected case-insensitive match")
	}
}

func TestStaticFunctions(t *testing.T) {
	if !IsMatch("abc123", `\d+`) {
		t.Error("expected IsMatch true")
	}

	result := ReplaceString("abc123def456", `\d+`, "X")
	if result != "abcXdefX" {
		t.Errorf("expected 'abcXdefX', got %q", result)
	}
}

func TestNoMatch(t *testing.T) {
	r := MustCompile(`\d+`)
	m := r.Match("abcdef")
	if m != nil {
		t.Error("expected nil match")
	}
}

func TestEmptyMatches(t *testing.T) {
	r := MustCompile(`xyz`)
	matches := r.Matches("abc")
	if matches.Count() != 0 {
		t.Errorf("expected 0 matches, got %d", matches.Count())
	}
}

func TestEscape(t *testing.T) {
	result := Escape(`*.+?|{}[]()^$#`)
	expected := `\*\.\+\?\|\{\}\[\]\(\)\^\$#`
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
