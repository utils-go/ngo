package datetime

import (
	"testing"
	"time"

	"github.com/utils-go/ngo/timespan"
)

func TestNewDateTime(t *testing.T) {
	dt := NewDateTime(2023, 12, 25, 15, 30, 45, 123)
	
	if dt.Year() != 2023 {
		t.Errorf("Expected Year() = 2023, got %d", dt.Year())
	}
	if dt.Month() != 12 {
		t.Errorf("Expected Month() = 12, got %d", dt.Month())
	}
	if dt.Day() != 25 {
		t.Errorf("Expected Day() = 25, got %d", dt.Day())
	}
	if dt.Hour() != 15 {
		t.Errorf("Expected Hour() = 15, got %d", dt.Hour())
	}
	if dt.Minute() != 30 {
		t.Errorf("Expected Minute() = 30, got %d", dt.Minute())
	}
	if dt.Second() != 45 {
		t.Errorf("Expected Second() = 45, got %d", dt.Second())
	}
	if dt.Millisecond() != 123 {
		t.Errorf("Expected Millisecond() = 123, got %d", dt.Millisecond())
	}
}

func TestNow(t *testing.T) {
	before := time.Now()
	dt := Now()
	after := time.Now()
	
	dtTime := dt.ToTime()
	if dtTime.Before(before) || dtTime.After(after) {
		t.Error("Now() should return current time")
	}
	
	if dt.Kind() != Local {
		t.Errorf("Now() should have Local kind, got %v", dt.Kind())
	}
}

func TestUtcNow(t *testing.T) {
	dt := UtcNow()
	
	if dt.Kind() != UTC {
		t.Errorf("UtcNow() should have UTC kind, got %v", dt.Kind())
	}
}

func TestToday(t *testing.T) {
	dt := Today()
	
	if dt.Hour() != 0 || dt.Minute() != 0 || dt.Second() != 0 || dt.Millisecond() != 0 {
		t.Error("Today() should have time component set to 00:00:00.000")
	}
	
	now := time.Now()
	if dt.Year() != now.Year() || dt.Month() != int(now.Month()) || dt.Day() != now.Day() {
		t.Error("Today() should have current date")
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		hasError bool
	}{
		{"2023-12-25 15:30:45", false},
		{"2023-12-25T15:30:45", false},
		{"2023-12-25", false},
		{"12/25/2023", false},
		{"12/25/2023 3:30:45 PM", false},
		{"invalid date", true},
		{"", true},
	}
	
	for _, test := range tests {
		dt, err := Parse(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for input '%s', but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input '%s': %v", test.input, err)
			}
			if dt == nil {
				t.Errorf("Expected valid DateTime for input '%s', got nil", test.input)
			}
		}
	}
}

func TestParseExact(t *testing.T) {
	dt, err := ParseExact("2023-12-25 15:30:45", "yyyy-MM-dd HH:mm:ss")
	if err != nil {
		t.Errorf("ParseExact failed: %v", err)
	}
	
	if dt.Year() != 2023 || dt.Month() != 12 || dt.Day() != 25 {
		t.Errorf("ParseExact result date = %d-%d-%d, expected 2023-12-25", dt.Year(), dt.Month(), dt.Day())
	}
	
	if dt.Hour() != 15 || dt.Minute() != 30 || dt.Second() != 45 {
		t.Errorf("ParseExact result time = %d:%d:%d, expected 15:30:45", dt.Hour(), dt.Minute(), dt.Second())
	}
}

func TestAddMethods(t *testing.T) {
	dt := NewDateTime(2023, 1, 1, 12, 0, 0, 0)
	
	// Test AddYears
	result := dt.AddYears(1)
	if result.Year() != 2024 {
		t.Errorf("AddYears(1) result year = %d, expected 2024", result.Year())
	}
	
	// Test AddMonths
	result = dt.AddMonths(6)
	if result.Month() != 7 {
		t.Errorf("AddMonths(6) result month = %d, expected 7", result.Month())
	}
	
	// Test AddDays
	result = dt.AddDays(15)
	if result.Day() != 16 {
		t.Errorf("AddDays(15) result day = %d, expected 16", result.Day())
	}
	
	// Test AddHours
	result = dt.AddHours(6)
	if result.Hour() != 18 {
		t.Errorf("AddHours(6) result hour = %d, expected 18", result.Hour())
	}
	
	// Test AddMinutes
	result = dt.AddMinutes(30)
	if result.Minute() != 30 {
		t.Errorf("AddMinutes(30) result minute = %d, expected 30", result.Minute())
	}
	
	// Test AddSeconds
	result = dt.AddSeconds(45)
	if result.Second() != 45 {
		t.Errorf("AddSeconds(45) result second = %d, expected 45", result.Second())
	}
	
	// Test AddMilliseconds
	result = dt.AddMilliseconds(500)
	if result.Millisecond() != 500 {
		t.Errorf("AddMilliseconds(500) result millisecond = %d, expected 500", result.Millisecond())
	}
}

func TestAddTimeSpan(t *testing.T) {
	dt := NewDateTime(2023, 1, 1, 12, 0, 0, 0)
	ts := timespan.NewTimeSpan(0, 2, 30, 45, 0) // 2 hours, 30 minutes, 45 seconds
	
	result := dt.Add(ts)
	
	if result.Hour() != 14 || result.Minute() != 30 || result.Second() != 45 {
		t.Errorf("Add(TimeSpan) result = %d:%d:%d, expected 14:30:45", result.Hour(), result.Minute(), result.Second())
	}
}

func TestSubtract(t *testing.T) {
	dt1 := NewDateTime(2023, 1, 2, 15, 30, 0, 0)
	dt2 := NewDateTime(2023, 1, 1, 12, 0, 0, 0)
	
	ts := dt1.Subtract(dt2)
	
	expectedHours := 27.5 // 1 day + 3.5 hours
	if ts.TotalHours() != expectedHours {
		t.Errorf("Subtract result = %f hours, expected %f", ts.TotalHours(), expectedHours)
	}
}

func TestSubtractTimeSpan(t *testing.T) {
	dt := NewDateTime(2023, 1, 1, 15, 30, 45, 0)
	ts := timespan.NewTimeSpan(0, 2, 15, 30, 0) // 2 hours, 15 minutes, 30 seconds
	
	result := dt.SubtractTimeSpan(ts)
	
	if result.Hour() != 13 || result.Minute() != 15 || result.Second() != 15 {
		t.Errorf("SubtractTimeSpan result = %d:%d:%d, expected 13:15:15", result.Hour(), result.Minute(), result.Second())
	}
}

func TestDate(t *testing.T) {
	dt := NewDateTime(2023, 12, 25, 15, 30, 45, 123)
	date := dt.Date()
	
	if date.Year() != 2023 || date.Month() != 12 || date.Day() != 25 {
		t.Errorf("Date() result date = %d-%d-%d, expected 2023-12-25", date.Year(), date.Month(), date.Day())
	}
	
	if date.Hour() != 0 || date.Minute() != 0 || date.Second() != 0 || date.Millisecond() != 0 {
		t.Error("Date() should have time component set to 00:00:00.000")
	}
}

func TestTimeOfDay(t *testing.T) {
	dt := NewDateTime(2023, 12, 25, 15, 30, 45, 123)
	timeOfDay := dt.TimeOfDay()
	
	expectedHours := 15
	expectedMinutes := 30
	expectedSeconds := 45
	
	if timeOfDay.Hours() != expectedHours {
		t.Errorf("TimeOfDay().Hours() = %d, expected %d", timeOfDay.Hours(), expectedHours)
	}
	if timeOfDay.Minutes() != expectedMinutes {
		t.Errorf("TimeOfDay().Minutes() = %d, expected %d", timeOfDay.Minutes(), expectedMinutes)
	}
	if timeOfDay.Seconds() != expectedSeconds {
		t.Errorf("TimeOfDay().Seconds() = %d, expected %d", timeOfDay.Seconds(), expectedSeconds)
	}
}

func TestToLocalTime(t *testing.T) {
	dt := NewDateTimeWithKind(2023, 12, 25, 15, 30, 45, 0, UTC)
	local := dt.ToLocalTime()
	
	if local.Kind() != Local {
		t.Errorf("ToLocalTime() should have Local kind, got %v", local.Kind())
	}
}

func TestToUniversalTime(t *testing.T) {
	dt := NewDateTimeWithKind(2023, 12, 25, 15, 30, 45, 0, Local)
	utc := dt.ToUniversalTime()
	
	if utc.Kind() != UTC {
		t.Errorf("ToUniversalTime() should have UTC kind, got %v", utc.Kind())
	}
}

func TestEquals(t *testing.T) {
	dt1 := NewDateTime(2023, 12, 25, 15, 30, 45, 0)
	dt2 := NewDateTime(2023, 12, 25, 15, 30, 45, 0)
	dt3 := NewDateTime(2023, 12, 25, 15, 30, 46, 0)
	
	if !dt1.Equals(dt2) {
		t.Error("Equal DateTimes should be equal")
	}
	
	if dt1.Equals(dt3) {
		t.Error("Different DateTimes should not be equal")
	}
}

func TestCompareTo(t *testing.T) {
	dt1 := NewDateTime(2023, 12, 25, 15, 30, 45, 0)
	dt2 := NewDateTime(2023, 12, 25, 15, 30, 46, 0)
	dt3 := NewDateTime(2023, 12, 25, 15, 30, 45, 0)
	
	if dt1.CompareTo(dt2) >= 0 {
		t.Error("dt1 should be less than dt2")
	}
	
	if dt2.CompareTo(dt1) <= 0 {
		t.Error("dt2 should be greater than dt1")
	}
	
	if dt1.CompareTo(dt3) != 0 {
		t.Error("dt1 should be equal to dt3")
	}
}

func TestToString(t *testing.T) {
	dt := NewDateTime(2023, 12, 25, 15, 30, 45, 0)
	
	result := dt.ToString()
	expected := "2023-12-25 15:30:45"
	
	if result != expected {
		t.Errorf("ToString() = '%s', expected '%s'", result, expected)
	}
}

func TestToStringWithFormat(t *testing.T) {
	dt := NewDateTime(2023, 12, 25, 15, 30, 45, 0)
	
	result := dt.ToStringWithFormat("yyyy-MM-dd HH:mm:ss")
	expected := "2023-12-25 15:30:45"
	
	if result != expected {
		t.Errorf("ToStringWithFormat() = '%s', expected '%s'", result, expected)
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year     int
		expected bool
	}{
		{2000, true},  // Divisible by 400
		{2004, true},  // Divisible by 4, not by 100
		{1900, false}, // Divisible by 100, not by 400
		{2001, false}, // Not divisible by 4
	}
	
	for _, test := range tests {
		result := IsLeapYear(test.year)
		if result != test.expected {
			t.Errorf("IsLeapYear(%d) = %t, expected %t", test.year, result, test.expected)
		}
	}
}

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		year, month int
		expected    int
	}{
		{2023, 1, 31},  // January
		{2023, 2, 28},  // February (non-leap year)
		{2024, 2, 29},  // February (leap year)
		{2023, 4, 30},  // April
		{2023, 12, 31}, // December
	}
	
	for _, test := range tests {
		result := DaysInMonth(test.year, test.month)
		if result != test.expected {
			t.Errorf("DaysInMonth(%d, %d) = %d, expected %d", test.year, test.month, result, test.expected)
		}
	}
}

func TestDayOfWeek(t *testing.T) {
	// January 1, 2023 was a Sunday
	dt := NewDateTime(2023, 1, 1, 0, 0, 0, 0)
	
	if dt.DayOfWeek() != time.Sunday {
		t.Errorf("DayOfWeek() = %v, expected %v", dt.DayOfWeek(), time.Sunday)
	}
}

func TestDayOfYear(t *testing.T) {
	// January 1st should be day 1
	dt := NewDateTime(2023, 1, 1, 0, 0, 0, 0)
	if dt.DayOfYear() != 1 {
		t.Errorf("DayOfYear() for Jan 1 = %d, expected 1", dt.DayOfYear())
	}
	
	// December 31st in non-leap year should be day 365
	dt = NewDateTime(2023, 12, 31, 0, 0, 0, 0)
	if dt.DayOfYear() != 365 {
		t.Errorf("DayOfYear() for Dec 31 (non-leap) = %d, expected 365", dt.DayOfYear())
	}
}