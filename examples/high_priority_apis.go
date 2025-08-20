package main

import (
	"fmt"
	"github.com/utils-go/ngo/collections/generic"
	"github.com/utils-go/ngo/converter"
	"github.com/utils-go/ngo/math"
	"github.com/utils-go/ngo/strings"
)

func main() {
	fmt.Println("=== NGo High Priority APIs Examples ===\n")

	// System.String examples
	fmt.Println("1. String Operations:")
	str := strings.NewString("  Hello, World!  ")
	fmt.Printf("Original: '%s'\n", str.Value())
	fmt.Printf("Trimmed: '%s'\n", str.Trim().Value())
	fmt.Printf("Upper: '%s'\n", str.ToUpper().Value())
	fmt.Printf("Contains 'World': %t\n", str.Contains("World"))
	fmt.Printf("Starts with 'Hello': %t\n", str.StartsWith("  Hello"))
	
	// String formatting
	formatted := strings.Format("Hello {0}, you are {1} years old!", "John", 25)
	fmt.Printf("Formatted: %s\n", formatted)
	
	// String splitting
	csv := strings.NewString("apple,banana,cherry")
	parts := csv.Split(",")
	fmt.Printf("Split result: %v\n", parts)
	fmt.Println()

	// System.Math examples
	fmt.Println("2. Math Operations:")
	fmt.Printf("Abs(-5.5): %.2f\n", math.Abs(-5.5))
	fmt.Printf("Max(10, 20): %.2f\n", math.Max(10, 20))
	fmt.Printf("Min(10, 20): %.2f\n", math.Min(10, 20))
	fmt.Printf("Round(3.7): %.2f\n", math.Round(3.7))
	fmt.Printf("Ceiling(3.2): %.2f\n", math.Ceiling(3.2))
	fmt.Printf("Floor(3.8): %.2f\n", math.Floor(3.8))
	fmt.Printf("Pow(2, 3): %.2f\n", math.Pow(2, 3))
	fmt.Printf("Sqrt(16): %.2f\n", math.Sqrt(16))
	fmt.Printf("Sin(PI/2): %.2f\n", math.Sin(math.PI/2))
	fmt.Printf("PI constant: %.6f\n", math.PI)
	fmt.Println()

	// System.Convert examples
	fmt.Println("3. Type Conversion:")
	
	// Convert to int32
	intVal, _ := converter.ToInt32("42")
	fmt.Printf("String '42' to int32: %d\n", intVal)
	
	// Convert to boolean
	boolVal, _ := converter.ToBoolean("true")
	fmt.Printf("String 'true' to bool: %t\n", boolVal)
	
	// Convert to string
	strVal := converter.ToString(123.45)
	fmt.Printf("Float 123.45 to string: '%s'\n", strVal)
	
	// Base64 encoding
	data := []byte("Hello, NGo!")
	base64Str := converter.ToBase64String(data)
	fmt.Printf("Base64 encode: %s\n", base64Str)
	
	decoded, _ := converter.FromBase64String(base64Str)
	fmt.Printf("Base64 decode: %s\n", string(decoded))
	fmt.Println()

	// System.Collections.Generic.List<T> examples
	fmt.Println("4. Generic List Operations:")
	
	// Create a list of integers
	intList := generic.NewList[int]()
	intList.AddRange([]int{5, 2, 8, 1, 9, 3})
	fmt.Printf("Original list: %v\n", intList.ToArray())
	fmt.Printf("Count: %d\n", intList.Count())
	
	// Find operations
	found, exists := intList.Find(func(x int) bool { return x > 5 })
	fmt.Printf("First element > 5: %d (exists: %t)\n", found, exists)
	
	// Filter operations
	evenNumbers := intList.FindAll(func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evenNumbers.ToArray())
	
	// Sort the list
	intList.Sort(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	fmt.Printf("Sorted list: %v\n", intList.ToArray())
	
	// Remove operations
	intList.Remove(2)
	fmt.Printf("After removing 2: %v\n", intList.ToArray())
	
	// String list example
	stringList := generic.NewList[string]()
	stringList.AddRange([]string{"apple", "banana", "cherry"})
	stringList.Insert(1, "apricot")
	fmt.Printf("String list: %v\n", stringList.ToArray())
	
	// ForEach operation
	fmt.Print("ForEach output: ")
	stringList.ForEach(func(item string) {
		fmt.Printf("%s ", item)
	})
	fmt.Println()
	
	// Check if all items meet condition
	allLongerThan3 := stringList.TrueForAll(func(s string) bool { return len(s) > 3 })
	fmt.Printf("All strings longer than 3 chars: %t\n", allLongerThan3)
	
	fmt.Println("\n=== All High Priority APIs Working! ===")
}