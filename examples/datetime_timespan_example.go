package main

import (
	"fmt"
	"github.com/utils-go/ngo/datetime"
	"github.com/utils-go/ngo/timespan"
)

func main1() {
	fmt.Println("=== NGo DateTime and TimeSpan Examples ===\n")

	// TimeSpan examples
	fmt.Println("1. TimeSpan Operations:")
	
	// Create TimeSpan instances
	ts1 := timespan.NewTimeSpan(1, 2, 30, 45, 500) // 1 day, 2 hours, 30 minutes, 45 seconds, 500 milliseconds
	fmt.Printf("TimeSpan 1: %s\n", ts1.String())
	fmt.Printf("  Days: %d, Hours: %d, Minutes: %d, Seconds: %d, Milliseconds: %d\n", 
		ts1.Days(), ts1.Hours(), ts1.Minutes(), ts1.Seconds(), ts1.Milliseconds())
	fmt.Printf("  Total Hours: %.2f, Total Minutes: %.2f\n", ts1.TotalHours(), ts1.TotalMinutes())
	
	// Create TimeSpan from different units
	ts2 := timespan.FromHours(3.5) // 3.5 hours
	fmt.Printf("TimeSpan from 3.5 hours: %s\n", ts2.String())
	
	ts3 := timespan.FromMinutes(90) // 90 minutes
	fmt.Printf("TimeSpan from 90 minutes: %s\n", ts3.String())
	
	// TimeSpan arithmetic
	sum := ts2.Add(ts3)
	fmt.Printf("3.5 hours + 90 minutes = %s\n", sum.String())
	
	diff := ts1.Subtract(ts2)
	fmt.Printf("TimeSpan1 - 3.5 hours = %s\n", diff.String())
	
	// Parse TimeSpan from string
	parsed, _ := timespan.Parse("2.15:30:45.123") // 2 days, 15 hours, 30 minutes, 45 seconds, 123 milliseconds
	fmt.Printf("Parsed TimeSpan '2.15:30:45.123': %s\n", parsed.String())
	
	// Parse Go duration format
	goDuration, _ := timespan.Parse("2h30m45s")
	fmt.Printf("Parsed Go duration '2h30m45s': %s\n", goDuration.String())
	fmt.Println()

	// DateTime examples
	fmt.Println("2. DateTime Operations:")
	
	// Current date and time
	now := datetime.Now()
	fmt.Printf("Current local time: %s\n", now.ToString())
	
	utcNow := datetime.UtcNow()
	fmt.Printf("Current UTC time: %s\n", utcNow.ToString())
	
	today := datetime.Today()
	fmt.Printf("Today (date only): %s\n", today.ToString())
	
	// Create specific DateTime
	dt := datetime.NewDateTime(2023, 12, 25, 15, 30, 45, 123)
	fmt.Printf("Christmas 2023: %s\n", dt.ToString())
	fmt.Printf("  Year: %d, Month: %d, Day: %d\n", dt.Year(), dt.Month(), dt.Day())
	fmt.Printf("  Hour: %d, Minute: %d, Second: %d, Millisecond: %d\n", 
		dt.Hour(), dt.Minute(), dt.Second(), dt.Millisecond())
	fmt.Printf("  Day of Week: %s, Day of Year: %d\n", dt.DayOfWeek(), dt.DayOfYear())
	
	// DateTime formatting
	fmt.Printf("Custom format (yyyy-MM-dd HH:mm:ss): %s\n", dt.ToStringWithFormat("yyyy-MM-dd HH:mm:ss"))
	fmt.Printf("Custom format (MM/dd/yyyy h:mm tt): %s\n", dt.ToStringWithFormat("MM/dd/yyyy h:mm tt"))
	
	// Parse DateTime from string
	parsed2, _ := datetime.Parse("2023-06-15 14:30:00")
	fmt.Printf("Parsed DateTime: %s\n", parsed2.ToString())
	
	// Parse with exact format
	parsed3, _ := datetime.ParseExact("25/12/2023 15:30:45", "dd/MM/yyyy HH:mm:ss")
	fmt.Printf("Parsed with exact format: %s\n", parsed3.ToString())
	fmt.Println()

	// DateTime arithmetic
	fmt.Println("3. DateTime Arithmetic:")
	
	baseDate := datetime.NewDateTime(2023, 1, 1, 12, 0, 0, 0)
	fmt.Printf("Base date: %s\n", baseDate.ToString())
	
	// Add different time units
	fmt.Printf("Add 1 year: %s\n", baseDate.AddYears(1).ToString())
	fmt.Printf("Add 6 months: %s\n", baseDate.AddMonths(6).ToString())
	fmt.Printf("Add 15 days: %s\n", baseDate.AddDays(15).ToString())
	fmt.Printf("Add 8 hours: %s\n", baseDate.AddHours(8).ToString())
	fmt.Printf("Add 30 minutes: %s\n", baseDate.AddMinutes(30).ToString())
	
	// Add TimeSpan to DateTime
	timeSpanToAdd := timespan.NewTimeSpan(0, 2, 15, 30, 0) // 2 hours, 15 minutes, 30 seconds
	result := baseDate.Add(timeSpanToAdd)
	fmt.Printf("Add TimeSpan %s: %s\n", timeSpanToAdd.String(), result.ToString())
	
	// Subtract DateTimes to get TimeSpan
	dt1 := datetime.NewDateTime(2023, 6, 15, 14, 30, 0, 0)
	dt2 := datetime.NewDateTime(2023, 6, 10, 10, 15, 0, 0)
	difference := dt1.Subtract(dt2)
	fmt.Printf("Difference between %s and %s: %s\n", 
		dt1.ToString(), dt2.ToString(), difference.String())
	fmt.Printf("  Total days: %.2f, Total hours: %.2f\n", 
		difference.TotalDays(), difference.TotalHours())
	fmt.Println()

	// Date components
	fmt.Println("4. Date Components:")
	
	sampleDate := datetime.NewDateTime(2023, 7, 4, 16, 45, 30, 250)
	fmt.Printf("Sample date: %s\n", sampleDate.ToString())
	
	// Get date part only
	dateOnly := sampleDate.Date()
	fmt.Printf("Date part only: %s\n", dateOnly.ToString())
	
	// Get time of day as TimeSpan
	timeOfDay := sampleDate.TimeOfDay()
	fmt.Printf("Time of day: %s\n", timeOfDay.String())
	
	// Time zone conversions
	fmt.Printf("Local time: %s\n", sampleDate.ToString())
	fmt.Printf("UTC time: %s\n", sampleDate.ToUniversalTime().ToString())
	fmt.Printf("Back to local: %s\n", sampleDate.ToUniversalTime().ToLocalTime().ToString())
	fmt.Println()

	// Utility functions
	fmt.Println("5. Utility Functions:")
	
	// Leap year check
	fmt.Printf("Is 2024 a leap year? %t\n", datetime.IsLeapYear(2024))
	fmt.Printf("Is 2023 a leap year? %t\n", datetime.IsLeapYear(2023))
	
	// Days in month
	fmt.Printf("Days in February 2024: %d\n", datetime.DaysInMonth(2024, 2))
	fmt.Printf("Days in February 2023: %d\n", datetime.DaysInMonth(2023, 2))
	
	// Comparisons
	dt3 := datetime.NewDateTime(2023, 6, 15, 10, 0, 0, 0)
	dt4 := datetime.NewDateTime(2023, 6, 20, 10, 0, 0, 0)
	
	fmt.Printf("Compare %s with %s: %d\n", 
		dt3.ToString(), dt4.ToString(), dt3.CompareTo(dt4))
	fmt.Printf("Are they equal? %t\n", dt3.Equals(dt4))
	
	fmt.Println("\n=== DateTime and TimeSpan Working Perfectly! ===")
}