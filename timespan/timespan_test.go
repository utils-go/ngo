package timespan

import (
	"testing"
	"time"
)

func TestNewTimeSpan(t *testing.T) {
	ts := NewTimeSpan(1, 2, 3, 4, 500) // 1 day, 2 hours, 3 minutes, 4 seconds, 500 milliseconds
	
	if ts.Days() != 1 {
		t.Errorf("Expected Days() = 1, got %d", ts.Days())
	}
	if ts.Hours() != 2 {
		t.Errorf("Expected Hours() = 2, got %d", ts.Hours())
	}
	if ts.Minutes() != 3 {
		t.Errorf("Expected Minutes() = 3, got %d", ts.Minutes())
	}
	if ts.Seconds() != 4 {
		t.Errorf("Expected Seconds() = 4, got %d", ts.Seconds())
	}
	if ts.Milliseconds() != 500 {
		t.Errorf("Expected Milliseconds() = 500, got %d", ts.Milliseconds())
	}
}

func TestFromMethods(t *testing.T) {
	// Test FromDays
	ts := FromDays(1.5)
	expected := 1.5
	if ts.TotalDays() != expected {
		t.Errorf("FromDays(1.5).TotalDays() = %f, expected %f", ts.TotalDays(), expected)
	}
	
	// Test FromHours
	ts = FromHours(2.5)
	expected = 2.5
	if ts.TotalHours() != expected {
		t.Errorf("FromHours(2.5).TotalHours() = %f, expected %f", ts.TotalHours(), expected)
	}
	
	// Test FromMinutes
	ts = FromMinutes(90)
	expected = 90.0
	if ts.TotalMinutes() != expected {
		t.Errorf("FromMinutes(90).TotalMinutes() = %f, expected %f", ts.TotalMinutes(), expected)
	}
	
	// Test FromSeconds
	ts = FromSeconds(3661) // 1 hour, 1 minute, 1 second
	if ts.Hours() != 1 || ts.Minutes() != 1 || ts.Seconds() != 1 {
		t.Errorf("FromSeconds(3661) = %d:%d:%d, expected 1:1:1", ts.Hours(), ts.Minutes(), ts.Seconds())
	}
}

func TestAdd(t *testing.T) {
	ts1 := NewTimeSpan(0, 1, 30, 0, 0) // 1 hour 30 minutes
	ts2 := NewTimeSpan(0, 0, 45, 30, 0) // 45 minutes 30 seconds
	
	result := ts1.Add(ts2)
	
	if result.Hours() != 2 || result.Minutes() != 15 || result.Seconds() != 30 {
		t.Errorf("Add result = %d:%d:%d, expected 2:15:30", result.Hours(), result.Minutes(), result.Seconds())
	}
}

func TestSubtract(t *testing.T) {
	ts1 := NewTimeSpan(0, 2, 30, 0, 0) // 2 hours 30 minutes
	ts2 := NewTimeSpan(0, 1, 15, 0, 0) // 1 hour 15 minutes
	
	result := ts1.Subtract(ts2)
	
	if result.Hours() != 1 || result.Minutes() != 15 {
		t.Errorf("Subtract result = %d:%d, expected 1:15", result.Hours(), result.Minutes())
	}
}

func TestNegate(t *testing.T) {
	ts := NewTimeSpan(0, 1, 30, 0, 0)
	negated := ts.Negate()
	
	if negated.TotalMinutes() != -90 {
		t.Errorf("Negate result = %f minutes, expected -90", negated.TotalMinutes())
	}
}

func TestDuration(t *testing.T) {
	ts := NewTimeSpan(0, -1, -30, 0, 0) // Negative timespan
	duration := ts.Duration()
	
	if duration.TotalMinutes() != 90 {
		t.Errorf("Duration result = %f minutes, expected 90", duration.TotalMinutes())
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		timespan *TimeSpan
		expected string
	}{
		{NewTimeSpan(0, 0, 0, 0, 0), "00:00:00"},
		{NewTimeSpan(0, 1, 2, 3, 0), "01:02:03"},
		{NewTimeSpan(0, 1, 2, 3, 456), "01:02:03.456"},
		{NewTimeSpan(1, 2, 3, 4, 0), "1.02:03:04"},
		{NewTimeSpan(1, 2, 3, 4, 567), "1.02:03:04.567"},
	}
	
	for _, test := range tests {
		result := test.timespan.String()
		if result != test.expected {
			t.Errorf("TimeSpan.String() = '%s', expected '%s'", result, test.expected)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"01:02:03", "01:02:03"},
		{"1.02:03:04", "1.02:03:04"},
		{"01:02:03.456", "01:02:03.456"},
		{"1h30m", "01:30:00"},
		{"2.5s", "00:00:02.500"},
	}
	
	for _, test := range tests {
		ts, err := Parse(test.input)
		if err != nil {
			t.Errorf("Parse('%s') failed: %v", test.input, err)
			continue
		}
		
		result := ts.String()
		if result != test.expected {
			t.Errorf("Parse('%s').String() = '%s', expected '%s'", test.input, result, test.expected)
		}
	}
}

func TestEquals(t *testing.T) {
	ts1 := NewTimeSpan(0, 1, 30, 0, 0)
	ts2 := NewTimeSpan(0, 1, 30, 0, 0)
	ts3 := NewTimeSpan(0, 2, 0, 0, 0)
	
	if !ts1.Equals(ts2) {
		t.Error("Equal timespans should be equal")
	}
	
	if ts1.Equals(ts3) {
		t.Error("Different timespans should not be equal")
	}
}

func TestCompareTo(t *testing.T) {
	ts1 := NewTimeSpan(0, 1, 0, 0, 0)
	ts2 := NewTimeSpan(0, 2, 0, 0, 0)
	ts3 := NewTimeSpan(0, 1, 0, 0, 0)
	
	if ts1.CompareTo(ts2) >= 0 {
		t.Error("ts1 should be less than ts2")
	}
	
	if ts2.CompareTo(ts1) <= 0 {
		t.Error("ts2 should be greater than ts1")
	}
	
	if ts1.CompareTo(ts3) != 0 {
		t.Error("ts1 should be equal to ts3")
	}
}

func TestToDuration(t *testing.T) {
	ts := NewTimeSpan(0, 1, 30, 45, 500)
	duration := ts.ToDuration()
	
	expected := 1*time.Hour + 30*time.Minute + 45*time.Second + 500*time.Millisecond
	if duration != expected {
		t.Errorf("ToDuration() = %v, expected %v", duration, expected)
	}
}

func TestNewTimeSpanFromDuration(t *testing.T) {
	duration := 2*time.Hour + 15*time.Minute + 30*time.Second
	ts := NewTimeSpanFromDuration(duration)
	
	if ts.Hours() != 2 || ts.Minutes() != 15 || ts.Seconds() != 30 {
		t.Errorf("NewTimeSpanFromDuration result = %d:%d:%d, expected 2:15:30", ts.Hours(), ts.Minutes(), ts.Seconds())
	}
}