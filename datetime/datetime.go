package datetime

import (
	"fmt"
	"strings"
	"time"

	"github.com/utils-go/ngo/timespan"
)

// DateTime represents an instant in time, typically expressed as a date and time of day
// Equivalent to System.DateTime in .NET
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.datetime?view=netframework-4.7.2

// DateTimeKind specifies whether a DateTime object represents a local time, UTC time, or neither
type DateTimeKind int

const (
	// Unspecified indicates that the time represented is not specified as either local time or UTC
	Unspecified DateTimeKind = iota
	// UTC indicates that the time represented is UTC
	UTC
	// Local indicates that the time represented is local time
	Local
)

// DateTime represents a date and time
type DateTime struct {
	time time.Time
	kind DateTimeKind
}

// Static DateTime values
var (
	// MinValue represents the smallest possible value of DateTime (January 1, 0001 00:00:00.000)
	MinValue = DateTime{time: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), kind: Unspecified}
	
	// MaxValue represents the largest possible value of DateTime (December 31, 9999 23:59:59.999)
	MaxValue = DateTime{time: time.Date(9999, 12, 31, 23, 59, 59, 999999999, time.UTC), kind: Unspecified}
)

// Now gets a DateTime object that is set to the current date and time on this computer, expressed as the local time
func Now() *DateTime {
	return &DateTime{time: time.Now(), kind: Local}
}

// UtcNow gets a DateTime object that is set to the current date and time on this computer, expressed as UTC
func UtcNow() *DateTime {
	return &DateTime{time: time.Now().UTC(), kind: UTC}
}

// Today gets the current date with the time component set to 00:00:00
func Today() *DateTime {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return &DateTime{time: today, kind: Local}
}

// NewDateTime creates a new DateTime instance
func NewDateTime(year, month, day, hour, minute, second, millisecond int) *DateTime {
	t := time.Date(year, time.Month(month), day, hour, minute, second, millisecond*1000000, time.Local)
	return &DateTime{time: t, kind: Local}
}

// NewDateTimeWithKind creates a new DateTime instance with the specified DateTimeKind
func NewDateTimeWithKind(year, month, day, hour, minute, second, millisecond int, kind DateTimeKind) *DateTime {
	var loc *time.Location
	switch kind {
	case UTC:
		loc = time.UTC
	case Local:
		loc = time.Local
	default:
		loc = time.UTC
	}
	
	t := time.Date(year, time.Month(month), day, hour, minute, second, millisecond*1000000, loc)
	return &DateTime{time: t, kind: kind}
}

// NewDateTimeFromTime creates a DateTime from Go's time.Time
func NewDateTimeFromTime(t time.Time) *DateTime {
	kind := Local
	if t.Location() == time.UTC {
		kind = UTC
	}
	return &DateTime{time: t, kind: kind}
}

// Parse converts the string representation of a date and time to its DateTime equivalent
func Parse(s string) (*DateTime, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, fmt.Errorf("input string is empty")
	}

	// Common date/time formats
	formats := []string{
		"2006-01-02 15:04:05.000",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.000",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"01/02/2006 15:04:05",
		"01/02/2006 15:04",
		"01/02/2006",
		"1/2/2006 3:04:05 PM",
		"1/2/2006 3:04 PM",
		"1/2/2006",
		time.RFC3339,
		time.RFC3339Nano,
		time.RFC822,
		time.RFC822Z,
		time.RFC1123,
		time.RFC1123Z,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			kind := Local
			if t.Location() == time.UTC || strings.HasSuffix(s, "Z") {
				kind = UTC
			}
			return &DateTime{time: t, kind: kind}, nil
		}
	}

	return nil, fmt.Errorf("unable to parse datetime: %s", s)
}

// ParseExact converts the string representation of a date and time to its DateTime equivalent using the specified format
func ParseExact(s, format string) (*DateTime, error) {
	// Convert .NET format to Go format
	goFormat := convertDotNetFormatToGo(format)
	
	t, err := time.Parse(goFormat, s)
	if err != nil {
		return nil, fmt.Errorf("unable to parse datetime '%s' with format '%s': %v", s, format, err)
	}

	kind := Local
	if t.Location() == time.UTC {
		kind = UTC
	}
	
	return &DateTime{time: t, kind: kind}, nil
}

// Properties

// Year gets the year component of the date represented by this instance
func (dt *DateTime) Year() int {
	return dt.time.Year()
}

// Month gets the month component of the date represented by this instance
func (dt *DateTime) Month() int {
	return int(dt.time.Month())
}

// Day gets the day of the month represented by this instance
func (dt *DateTime) Day() int {
	return dt.time.Day()
}

// Hour gets the hour component of the date represented by this instance
func (dt *DateTime) Hour() int {
	return dt.time.Hour()
}

// Minute gets the minute component of the date represented by this instance
func (dt *DateTime) Minute() int {
	return dt.time.Minute()
}

// Second gets the second component of the date represented by this instance
func (dt *DateTime) Second() int {
	return dt.time.Second()
}

// Millisecond gets the milliseconds component of the date represented by this instance
func (dt *DateTime) Millisecond() int {
	return dt.time.Nanosecond() / 1000000
}

// DayOfWeek gets the day of the week represented by this instance
func (dt *DateTime) DayOfWeek() time.Weekday {
	return dt.time.Weekday()
}

// DayOfYear gets the day of the year represented by this instance
func (dt *DateTime) DayOfYear() int {
	return dt.time.YearDay()
}

// Kind gets a value that indicates whether the time represented by this instance is based on local time, UTC, or neither
func (dt *DateTime) Kind() DateTimeKind {
	return dt.kind
}

// Date gets the date component of this instance
func (dt *DateTime) Date() *DateTime {
	date := time.Date(dt.time.Year(), dt.time.Month(), dt.time.Day(), 0, 0, 0, 0, dt.time.Location())
	return &DateTime{time: date, kind: dt.kind}
}

// TimeOfDay gets the time of day for this instance
func (dt *DateTime) TimeOfDay() *timespan.TimeSpan {
	totalNanoseconds := int64(dt.time.Hour())*int64(time.Hour) +
		int64(dt.time.Minute())*int64(time.Minute) +
		int64(dt.time.Second())*int64(time.Second) +
		int64(dt.time.Nanosecond())
	
	// Convert nanoseconds to ticks (1 tick = 100 nanoseconds)
	ticks := totalNanoseconds / 100
	return timespan.FromTicks(ticks)
}

// Ticks gets the number of ticks that represent the date and time of this instance
func (dt *DateTime) Ticks() int64 {
	// .NET ticks start from January 1, 0001 00:00:00 UTC
	// Go's Unix time starts from January 1, 1970 00:00:00 UTC
	// We need to add the offset between these two epochs
	
	// Number of ticks from 0001-01-01 to 1970-01-01
	const ticksTo1970 int64 = 621355968000000000
	
	unixNano := dt.time.UnixNano()
	// Convert nanoseconds to ticks (1 tick = 100 nanoseconds)
	ticks := unixNano/100 + ticksTo1970
	
	return ticks
}

// Methods

// Add returns a new DateTime that adds the value of the specified TimeSpan to the value of this instance
func (dt *DateTime) Add(ts *timespan.TimeSpan) *DateTime {
	duration := ts.ToDuration()
	newTime := dt.time.Add(duration)
	return &DateTime{time: newTime, kind: dt.kind}
}

// AddYears returns a new DateTime that adds the specified number of years to the value of this instance
func (dt *DateTime) AddYears(years int) *DateTime {
	newTime := dt.time.AddDate(years, 0, 0)
	return &DateTime{time: newTime, kind: dt.kind}
}

// AddMonths returns a new DateTime that adds the specified number of months to the value of this instance
func (dt *DateTime) AddMonths(months int) *DateTime {
	newTime := dt.time.AddDate(0, months, 0)
	return &DateTime{time: newTime, kind: dt.kind}
}

// AddDays returns a new DateTime that adds the specified number of days to the value of this instance
func (dt *DateTime) AddDays(days float64) *DateTime {
	duration := time.Duration(days * float64(24*time.Hour))
	newTime := dt.time.Add(duration)
	return &DateTime{time: newTime, kind: dt.kind}
}

// AddHours returns a new DateTime that adds the specified number of hours to the value of this instance
func (dt *DateTime) AddHours(hours float64) *DateTime {
	duration := time.Duration(hours * float64(time.Hour))
	newTime := dt.time.Add(duration)
	return &DateTime{time: newTime, kind: dt.kind}
}

// AddMinutes returns a new DateTime that adds the specified number of minutes to the value of this instance
func (dt *DateTime) AddMinutes(minutes float64) *DateTime {
	duration := time.Duration(minutes * float64(time.Minute))
	newTime := dt.time.Add(duration)
	return &DateTime{time: newTime, kind: dt.kind}
}

// AddSeconds returns a new DateTime that adds the specified number of seconds to the value of this instance
func (dt *DateTime) AddSeconds(seconds float64) *DateTime {
	duration := time.Duration(seconds * float64(time.Second))
	newTime := dt.time.Add(duration)
	return &DateTime{time: newTime, kind: dt.kind}
}

// AddMilliseconds returns a new DateTime that adds the specified number of milliseconds to the value of this instance
func (dt *DateTime) AddMilliseconds(milliseconds float64) *DateTime {
	duration := time.Duration(milliseconds * float64(time.Millisecond))
	newTime := dt.time.Add(duration)
	return &DateTime{time: newTime, kind: dt.kind}
}

// Subtract returns a new TimeSpan that represents the difference between two DateTime instances
func (dt *DateTime) Subtract(other *DateTime) *timespan.TimeSpan {
	duration := dt.time.Sub(other.time)
	return timespan.NewTimeSpanFromDuration(duration)
}

// SubtractTimeSpan returns a new DateTime that subtracts the specified TimeSpan from this instance
func (dt *DateTime) SubtractTimeSpan(ts *timespan.TimeSpan) *DateTime {
	duration := ts.ToDuration()
	newTime := dt.time.Add(-duration)
	return &DateTime{time: newTime, kind: dt.kind}
}

// ToLocalTime converts the value of the current DateTime object to local time
func (dt *DateTime) ToLocalTime() *DateTime {
	localTime := dt.time.Local()
	return &DateTime{time: localTime, kind: Local}
}

// ToUniversalTime converts the value of the current DateTime object to UTC
func (dt *DateTime) ToUniversalTime() *DateTime {
	utcTime := dt.time.UTC()
	return &DateTime{time: utcTime, kind: UTC}
}

// Equals determines whether two specified instances of DateTime are equal
func (dt *DateTime) Equals(other *DateTime) bool {
	if other == nil {
		return false
	}
	return dt.time.Equal(other.time)
}

// CompareTo compares the value of this instance to a specified DateTime value
func (dt *DateTime) CompareTo(other *DateTime) int {
	if other == nil {
		return 1
	}
	
	if dt.time.Before(other.time) {
		return -1
	} else if dt.time.After(other.time) {
		return 1
	}
	return 0
}

// ToString converts the value of the current DateTime object to its equivalent string representation
func (dt *DateTime) ToString() string {
	return dt.time.Format("2006-01-02 15:04:05")
}

// ToStringWithFormat converts the value of the current DateTime object to its equivalent string representation using the specified format
func (dt *DateTime) ToStringWithFormat(format string) string {
	goFormat := convertDotNetFormatToGo(format)
	return dt.time.Format(goFormat)
}

// ToTime converts the DateTime to Go's time.Time
func (dt *DateTime) ToTime() time.Time {
	return dt.time
}

// String returns a string representation of the DateTime
func (dt *DateTime) String() string {
	return dt.ToString()
}

// Static methods

// Compare compares two instances of DateTime and returns an integer that indicates their relationship
func Compare(t1, t2 *DateTime) int {
	if t1 == nil && t2 == nil {
		return 0
	}
	if t1 == nil {
		return -1
	}
	if t2 == nil {
		return 1
	}
	return t1.CompareTo(t2)
}

// IsLeapYear returns an indication whether the specified year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInMonth returns the number of days in the specified month and year
func DaysInMonth(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
	}
}

// Helper function to convert .NET datetime format to Go format
func convertDotNetFormatToGo(dotNetFormat string) string {
	// We need to replace in a specific order to avoid conflicts
	// Replace longer patterns first to avoid partial matches
	
	result := dotNetFormat
	
	// Replace in order from longest to shortest to avoid conflicts
	replacements := []struct {
		dotNet string
		goFmt  string
	}{
		{"yyyy", "2006"},
		{"MM", "01"},
		{"dd", "02"},
		{"HH", "15"},
		{"mm", "04"},
		{"ss", "05"},
		{"fff", "000"},
		{"ff", "00"},
		{"yy", "06"},
		{"M", "1"},
		{"d", "2"},
		{"H", "15"},
		{"hh", "03"},
		{"h", "3"},
		{"m", "4"},
		{"s", "5"},
		{"f", "0"},
		{"tt", "PM"},
		{"t", "PM"},
	}

	for _, replacement := range replacements {
		result = strings.ReplaceAll(result, replacement.dotNet, replacement.goFmt)
	}

	return result
}