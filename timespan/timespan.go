package timespan

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// TimeSpan represents a time interval equivalent to System.TimeSpan in .NET
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.timespan?view=netframework-4.7.2

// Constants for time calculations (matching .NET TimeSpan)
const (
	// Number of ticks per time unit (1 tick = 100 nanoseconds in .NET)
	TicksPerMillisecond int64 = 10000
	TicksPerSecond      int64 = TicksPerMillisecond * 1000
	TicksPerMinute      int64 = TicksPerSecond * 60
	TicksPerHour        int64 = TicksPerMinute * 60
	TicksPerDay         int64 = TicksPerHour * 24

	// Conversion factors
	MillisecondsPerTick float64 = 1.0 / float64(TicksPerMillisecond)
	SecondsPerTick      float64 = 1.0 / float64(TicksPerSecond)
	MinutesPerTick      float64 = 1.0 / float64(TicksPerMinute)
	HoursPerTick        float64 = 1.0 / float64(TicksPerHour)
	DaysPerTick         float64 = 1.0 / float64(TicksPerDay)

	// Milliseconds per time unit
	MillisecondsPerSecond = 1000
	MillisecondsPerMinute = MillisecondsPerSecond * 60
	MillisecondsPerHour   = MillisecondsPerMinute * 60
	MillisecondsPerDay    = MillisecondsPerHour * 24

	// Maximum and minimum values
	MaxValue int64 = math.MaxInt64
	MinValue int64 = math.MinInt64
)

// TimeSpan represents a time interval
type TimeSpan struct {
	ticks int64 // Internal representation in ticks (100-nanosecond units)
}

// Zero represents a TimeSpan with zero duration
var Zero = TimeSpan{ticks: 0}

// MaxTimeSpan represents the maximum TimeSpan value
var MaxTimeSpan = TimeSpan{ticks: MaxValue}

// MinTimeSpan represents the minimum TimeSpan value
var MinTimeSpan = TimeSpan{ticks: MinValue}

// NewTimeSpan creates a new TimeSpan from days, hours, minutes, seconds, and milliseconds
func NewTimeSpan(days, hours, minutes, seconds, milliseconds int) *TimeSpan {
	totalTicks := int64(days)*TicksPerDay +
		int64(hours)*TicksPerHour +
		int64(minutes)*TicksPerMinute +
		int64(seconds)*TicksPerSecond +
		int64(milliseconds)*TicksPerMillisecond

	return &TimeSpan{ticks: totalTicks}
}

// NewTimeSpanFromTicks creates a new TimeSpan from the specified number of ticks
func NewTimeSpanFromTicks(ticks int64) *TimeSpan {
	return &TimeSpan{ticks: ticks}
}

// NewTimeSpanFromDuration creates a TimeSpan from Go's time.Duration
func NewTimeSpanFromDuration(d time.Duration) *TimeSpan {
	// Convert nanoseconds to ticks (1 tick = 100 nanoseconds)
	ticks := d.Nanoseconds() / 100
	return &TimeSpan{ticks: ticks}
}

// Days gets the days component of the time interval
func (ts *TimeSpan) Days() int {
	return int(ts.ticks / TicksPerDay)
}

// Hours gets the hours component of the time interval (0-23)
func (ts *TimeSpan) Hours() int {
	return int((ts.ticks / TicksPerHour) % 24)
}

// Minutes gets the minutes component of the time interval (0-59)
func (ts *TimeSpan) Minutes() int {
	return int((ts.ticks / TicksPerMinute) % 60)
}

// Seconds gets the seconds component of the time interval (0-59)
func (ts *TimeSpan) Seconds() int {
	return int((ts.ticks / TicksPerSecond) % 60)
}

// Milliseconds gets the milliseconds component of the time interval (0-999)
func (ts *TimeSpan) Milliseconds() int {
	return int((ts.ticks / TicksPerMillisecond) % 1000)
}

// Ticks gets the number of ticks that represent the value of the current TimeSpan
func (ts *TimeSpan) Ticks() int64 {
	return ts.ticks
}

// TotalDays gets the value of the current TimeSpan expressed in whole and fractional days
func (ts *TimeSpan) TotalDays() float64 {
	return float64(ts.ticks) * DaysPerTick
}

// TotalHours gets the value of the current TimeSpan expressed in whole and fractional hours
func (ts *TimeSpan) TotalHours() float64 {
	return float64(ts.ticks) * HoursPerTick
}

// TotalMinutes gets the value of the current TimeSpan expressed in whole and fractional minutes
func (ts *TimeSpan) TotalMinutes() float64 {
	return float64(ts.ticks) * MinutesPerTick
}

// TotalSeconds gets the value of the current TimeSpan expressed in whole and fractional seconds
func (ts *TimeSpan) TotalSeconds() float64 {
	return float64(ts.ticks) * SecondsPerTick
}

// TotalMilliseconds gets the value of the current TimeSpan expressed in whole and fractional milliseconds
func (ts *TimeSpan) TotalMilliseconds() float64 {
	return float64(ts.ticks) * MillisecondsPerTick
}

// Add returns a new TimeSpan whose value is the sum of the specified TimeSpan and this instance
func (ts *TimeSpan) Add(other *TimeSpan) *TimeSpan {
	return &TimeSpan{ticks: ts.ticks + other.ticks}
}

// Subtract returns a new TimeSpan whose value is the difference between the specified TimeSpan and this instance
func (ts *TimeSpan) Subtract(other *TimeSpan) *TimeSpan {
	return &TimeSpan{ticks: ts.ticks - other.ticks}
}

// Negate returns a new TimeSpan whose value is the negated value of this instance
func (ts *TimeSpan) Negate() *TimeSpan {
	return &TimeSpan{ticks: -ts.ticks}
}

// Duration returns a new TimeSpan whose value is the absolute value of the current TimeSpan
func (ts *TimeSpan) Duration() *TimeSpan {
	if ts.ticks < 0 {
		return &TimeSpan{ticks: -ts.ticks}
	}
	return &TimeSpan{ticks: ts.ticks}
}

// Equals determines whether two specified instances of TimeSpan are equal
func (ts *TimeSpan) Equals(other *TimeSpan) bool {
	if other == nil {
		return false
	}
	return ts.ticks == other.ticks
}

// CompareTo compares two TimeSpan values and returns an integer that indicates their relationship
func (ts *TimeSpan) CompareTo(other *TimeSpan) int {
	if other == nil {
		return 1
	}
	if ts.ticks < other.ticks {
		return -1
	} else if ts.ticks > other.ticks {
		return 1
	}
	return 0
}

// String returns a string representation of the TimeSpan
func (ts *TimeSpan) String() string {
	if ts.ticks == 0 {
		return "00:00:00"
	}

	negative := ts.ticks < 0
	ticks := ts.ticks
	if negative {
		ticks = -ticks
	}

	days := ticks / TicksPerDay
	hours := (ticks / TicksPerHour) % 24
	minutes := (ticks / TicksPerMinute) % 60
	seconds := (ticks / TicksPerSecond) % 60
	milliseconds := (ticks / TicksPerMillisecond) % 1000

	var result string
	if negative {
		result = "-"
	}

	if days > 0 {
		if milliseconds > 0 {
			result += fmt.Sprintf("%d.%02d:%02d:%02d.%03d", days, hours, minutes, seconds, milliseconds)
		} else {
			result += fmt.Sprintf("%d.%02d:%02d:%02d", days, hours, minutes, seconds)
		}
	} else {
		if milliseconds > 0 {
			result += fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, milliseconds)
		} else {
			result += fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
		}
	}

	return result
}

// ToDuration converts the TimeSpan to Go's time.Duration
func (ts *TimeSpan) ToDuration() time.Duration {
	// Convert ticks to nanoseconds (1 tick = 100 nanoseconds)
	return time.Duration(ts.ticks * 100)
}

// Static factory methods

// FromDays returns a TimeSpan that represents a specified number of days
func FromDays(value float64) *TimeSpan {
	return &TimeSpan{ticks: int64(value * float64(TicksPerDay))}
}

// FromHours returns a TimeSpan that represents a specified number of hours
func FromHours(value float64) *TimeSpan {
	return &TimeSpan{ticks: int64(value * float64(TicksPerHour))}
}

// FromMinutes returns a TimeSpan that represents a specified number of minutes
func FromMinutes(value float64) *TimeSpan {
	return &TimeSpan{ticks: int64(value * float64(TicksPerMinute))}
}

// FromSeconds returns a TimeSpan that represents a specified number of seconds
func FromSeconds(value float64) *TimeSpan {
	return &TimeSpan{ticks: int64(value * float64(TicksPerSecond))}
}

// FromMilliseconds returns a TimeSpan that represents a specified number of milliseconds
func FromMilliseconds(value float64) *TimeSpan {
	return &TimeSpan{ticks: int64(value * float64(TicksPerMillisecond))}
}

// FromTicks returns a TimeSpan that represents a specified time, where the specification is in units of ticks
func FromTicks(value int64) *TimeSpan {
	return &TimeSpan{ticks: value}
}

// Parse converts the string representation of a time interval to its TimeSpan equivalent
func Parse(s string) (*TimeSpan, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, fmt.Errorf("input string is empty")
	}

	// Handle negative values
	negative := false
	if strings.HasPrefix(s, "-") {
		negative = true
		s = s[1:]
	}

	// Try to parse as Go duration first (e.g., "1h30m", "2.5s")
	if duration, err := time.ParseDuration(s); err == nil {
		ts := NewTimeSpanFromDuration(duration)
		if negative {
			ts = ts.Negate()
		}
		return ts, nil
	}

	// Parse .NET TimeSpan format: [d.]hh:mm:ss[.fffffff]
	parts := strings.Split(s, ".")
	var days int64 = 0
	var timeStr string
	var milliseconds int64 = 0

	if len(parts) == 1 {
		// No dots, just time part
		timeStr = parts[0]
	} else if len(parts) == 2 {
		// Could be days.time or time.milliseconds
		if strings.Contains(parts[1], ":") {
			// days.time format
			if d, err := strconv.ParseInt(parts[0], 10, 64); err == nil {
				days = d
				timeStr = parts[1]
			} else {
				return nil, fmt.Errorf("invalid days part: %s", parts[0])
			}
		} else {
			// time.milliseconds format
			timeStr = parts[0]
			// Pad or truncate to 3 digits
			msStr := parts[1]
			if len(msStr) > 3 {
				msStr = msStr[:3]
			} else {
				for len(msStr) < 3 {
					msStr += "0"
				}
			}
			if ms, err := strconv.ParseInt(msStr, 10, 64); err == nil {
				milliseconds = ms
			} else {
				return nil, fmt.Errorf("invalid milliseconds part: %s", parts[1])
			}
		}
	} else if len(parts) == 3 {
		// days.time.milliseconds format
		if d, err := strconv.ParseInt(parts[0], 10, 64); err == nil {
			days = d
			timeStr = parts[1]
			if ms, err := strconv.ParseInt(parts[2], 10, 64); err == nil {
				milliseconds = ms
			} else {
				return nil, fmt.Errorf("invalid milliseconds part: %s", parts[2])
			}
		} else {
			return nil, fmt.Errorf("invalid days part: %s", parts[0])
		}
	} else {
		return nil, fmt.Errorf("invalid TimeSpan format: %s", s)
	}

	// Parse time part (hh:mm:ss)
	timeParts := strings.Split(timeStr, ":")
	if len(timeParts) != 3 {
		return nil, fmt.Errorf("invalid time format: %s", timeStr)
	}

	hours, err := strconv.ParseInt(timeParts[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid hours: %s", timeParts[0])
	}

	minutes, err := strconv.ParseInt(timeParts[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid minutes: %s", timeParts[1])
	}

	seconds, err := strconv.ParseInt(timeParts[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid seconds: %s", timeParts[2])
	}

	// Calculate total ticks
	totalTicks := days*TicksPerDay +
		hours*TicksPerHour +
		minutes*TicksPerMinute +
		seconds*TicksPerSecond +
		milliseconds*TicksPerMillisecond

	if negative {
		totalTicks = -totalTicks
	}

	return &TimeSpan{ticks: totalTicks}, nil
}

// Compare compares two TimeSpan values and returns an integer that indicates their relationship
func Compare(t1, t2 *TimeSpan) int {
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