package main

import (
	"fmt"
	"time"

	"github.com/utils-go/ngo/collections/generic"
	"github.com/utils-go/ngo/diagnostics"
	"github.com/utils-go/ngo/environment"
	"github.com/utils-go/ngo/linq"
	"github.com/utils-go/ngo/text"
)

func main() {
	fmt.Println("=== NGo New APIs Examples ===\n")

	// 1. StringBuilder Examples
	fmt.Println("1. StringBuilder Operations:")
	sb := text.NewStringBuilder()
	sb.AppendString("Hello").AppendString(" ").AppendString("World")
	sb.AppendLine().AppendLineString("This is a new line")
	sb.AppendFormat("Formatted: %s = %d", "answer", 42)
	
	fmt.Printf("StringBuilder result:\n%s\n", sb.ToString())
	fmt.Printf("Length: %d, Capacity: %d\n", sb.Length(), sb.Capacity())
	
	// String manipulation
	sb.Insert(6, "Beautiful ")
	sb.Replace("World", "Go")
	fmt.Printf("After manipulation: %s\n", sb.ToString())
	fmt.Println()

	// 2. Dictionary Examples
	fmt.Println("2. Dictionary Operations:")
	dict := generic.NewDictionary[string, int]()
	
	// Add items
	dict.Add("apple", 5)
	dict.Add("banana", 3)
	dict.Add("cherry", 8)
	dict.Set("date", 2) // Set can add or update
	
	fmt.Printf("Dictionary count: %d\n", dict.Count())
	fmt.Printf("Contains 'apple': %t\n", dict.ContainsKey("apple"))
	
	if value, exists := dict.TryGetValue("banana"); exists {
		fmt.Printf("Banana count: %d\n", value)
	}
	
	// Dictionary operations
	fmt.Printf("Keys: %v\n", dict.Keys())
	fmt.Printf("Values: %v\n", dict.Values())
	
	// LINQ-style operations on dictionary
	expensive := dict.Where(func(k string, v int) bool { return v > 5 })
	fmt.Printf("Expensive items: %s\n", expensive.String())
	
	hasLowStock := dict.Any(func(k string, v int) bool { return v < 3 })
	fmt.Printf("Has low stock items: %t\n", hasLowStock)
	fmt.Println()

	// 3. LINQ Examples
	fmt.Println("3. LINQ Operations:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	enum := linq.From(numbers)
	
	fmt.Printf("Original: %v\n", enum.ToSlice())
	
	// Filter even numbers
	evens := enum.Where(func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens.ToSlice())
	
	// Square all numbers
	squares := linq.Select(enum, func(x int) int { return x * x })
	fmt.Printf("Squares: %v\n", squares.ToSlice())
	
	// Chain operations
	result := enum.Where(func(x int) bool { return x > 5 }).
		Take(3).
		ToSlice()
	fmt.Printf("Numbers > 5, take 3: %v\n", result)
	
	// Aggregations
	sum := linq.Aggregate(enum, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("Sum of all numbers: %d\n", sum)
	
	// Check conditions
	hasLarge := enum.Any(func(x int) bool { return x > 8 })
	allPositive := enum.All(func(x int) bool { return x > 0 })
	fmt.Printf("Has numbers > 8: %t, All positive: %t\n", hasLarge, allPositive)
	
	// Set operations
	enum1 := linq.From([]int{1, 2, 3, 4})
	enum2 := linq.From([]int{3, 4, 5, 6})
	
	union := enum1.Union(enum2)
	intersection := enum1.Intersect(enum2)
	fmt.Printf("Union: %v\n", union.ToSlice())
	fmt.Printf("Intersection: %v\n", intersection.ToSlice())
	
	// Order operations
	shuffled := linq.From([]int{5, 2, 8, 1, 9, 3})
	ordered := linq.OrderBy(shuffled, func(x int) int { return x }, func(a, b int) int {
		if a < b { return -1 }
		if a > b { return 1 }
		return 0
	})
	fmt.Printf("Ordered: %v\n", ordered.ToSlice())
	fmt.Println()

	// 4. Environment Examples
	fmt.Println("4. Environment Information:")
	fmt.Printf("Machine Name: %s\n", environment.MachineName())
	fmt.Printf("OS Version: %s\n", environment.OSVersion())
	fmt.Printf("User Name: %s\n", environment.UserName())
	fmt.Printf("Processor Count: %d\n", environment.ProcessorCount())
	fmt.Printf("Is 64-bit Process: %t\n", environment.Is64BitProcess())
	fmt.Printf("Current Directory: %s\n", environment.CurrentDirectory())
	
	// Environment variables
	fmt.Printf("PATH exists: %t\n", environment.GetEnvironmentVariable("PATH") != "")
	
	// Special folders
	userProfile := environment.GetFolderPath(environment.UserProfile)
	desktop := environment.GetFolderPath(environment.Desktop)
	fmt.Printf("User Profile: %s\n", userProfile)
	fmt.Printf("Desktop: %s\n", desktop)
	
	// Command line args
	args := environment.GetCommandLineArgs()
	fmt.Printf("Command line args count: %d\n", len(args))
	if len(args) > 0 {
		fmt.Printf("Program name: %s\n", args[0])
	}
	fmt.Println()

	// 5. Stopwatch Examples
	fmt.Println("5. Stopwatch Performance Measurement:")
	
	// Measure a simple operation
	sw := diagnostics.StartNew()
	
	// Simulate some work
	time.Sleep(50 * time.Millisecond)
	
	sw.Stop()
	fmt.Printf("Operation took: %d ms\n", sw.ElapsedMilliseconds())
	fmt.Printf("Operation took: %v\n", sw.Elapsed())
	
	// Measure multiple operations
	sw.Restart()
	
	// Simulate more work
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}
	
	sw.Stop()
	fmt.Printf("Loop took: %d ms\n", sw.ElapsedMilliseconds())
	
	// Accumulative timing
	sw.Reset()
	
	for i := 0; i < 3; i++ {
		sw.Start()
		time.Sleep(10 * time.Millisecond)
		sw.Stop()
		fmt.Printf("Iteration %d: %d ms (total: %d ms)\n", 
			i+1, 10, sw.ElapsedMilliseconds())
	}
	
	fmt.Printf("Total accumulated time: %d ms\n", sw.ElapsedMilliseconds())
	fmt.Println()

	// 6. Combined Example: Processing Data with Performance Monitoring
	fmt.Println("6. Combined Example - Data Processing:")
	
	// Create sample data
	data := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape"}
	
	// Use StringBuilder to build a report
	report := text.NewStringBuilder()
	report.AppendLineString("=== Fruit Analysis Report ===")
	
	// Use Stopwatch to measure processing time
	processingTimer := diagnostics.StartNew()
	
	// Use LINQ to process data
	fruits := linq.From(data)
	
	// Find long fruit names
	longNames := fruits.Where(func(s string) bool { return len(s) > 5 })
	report.AppendFormat("Fruits with names longer than 5 characters: %d\n", longNames.Count())
	
	// Get fruit lengths using LINQ
	lengths := linq.Select(fruits, func(s string) int { return len(s) })
	avgLength := linq.Aggregate(lengths, 0, func(acc, x int) int { return acc + x }) / lengths.Count()
	report.AppendFormat("Average fruit name length: %d characters\n", avgLength)
	
	// Use Dictionary to count by length
	lengthCounts := generic.NewDictionary[int, int]()
	fruits.ForEach(func(fruit string) {
		length := len(fruit)
		if count, exists := lengthCounts.TryGetValue(length); exists {
			lengthCounts.Set(length, count+1)
		} else {
			lengthCounts.Set(length, 1)
		}
	})
	
	report.AppendLineString("Length distribution:")
	lengthCounts.ForEach(func(length, count int) {
		report.AppendFormat("  %d characters: %d fruits\n", length, count)
	})
	
	processingTimer.Stop()
	
	report.AppendFormat("Processing completed in: %d ms\n", processingTimer.ElapsedMilliseconds())
	report.AppendFormat("Environment: %s on %s\n", 
		environment.UserName(), environment.MachineName())
	
	fmt.Print(report.ToString())
	
	fmt.Println("\n=== All New APIs Working Perfectly! ===")
}