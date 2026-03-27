# NGo
`ngo` means .net api implement by go

the .net api implement by go. you can write go as like .net

this is based on .net4.7.2 api

api reference:https://learn.microsoft.com/en-us/dotnet/api/?view=netframework-4.7.2&preserve-view=true

source code reference:https://referencesource.microsoft.com/
# Usage
 ```
 go get -u github.com/lishuangquan1987/ngo
 ```
file: same as `Sytem.IO.File`
```go
package main

import (
	"fmt"
	"github.com/lishuangquan1987/ngo/io/file"
)

func main() {
	content := "this is ngo example"
	file.WriteAllText("./test.txt", content)
	result, err := file.ReadAllText("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

```
bitconverter: same as `System.BitConverter`,but it is not static methods,so you can change `LittleEndian` or `BigEndian` any where
```go
package main

import (
	"fmt"
	"github.com/lishuangquan1987/ngo/bitconverter"
)

func main() {
	converter := bitconverter.BitConverter{IsLittleEndian: true} //littleEndian at the front
	value := int64(11234567489)
	bytes, err := converter.GetBytesFromInt64E(value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bytes)
	valueNew, err := converter.ToInt64E(bytes, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value == valueNew)
}
```

strings: same as `System.String` with .NET-style string manipulation
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/stringUtils"
)

func main() {
	// String utility functions
	s := "  Hello, World!  "
	fmt.Println("Original:", s)
	fmt.Println("Trimmed:", stringUtils.Trim(s))
	fmt.Println("Upper:", stringUtils.ToUpper(s))
	fmt.Println("Contains 'World':", stringUtils.Contains(s, "World"))
	
	// String formatting like .NET
	formatted := stringUtils.Format("Hello {0}, you are {1}!", "John", 25)
	fmt.Println("Formatted:", formatted)
}
```

math: same as `System.Math` with mathematical functions and constants
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/math"
)

func main() {
	fmt.Println("Abs(-5.5):", math.Abs(-5.5))
	fmt.Println("Max(10, 20):", math.Max(10, 20))
	fmt.Println("Round(3.7):", math.Round(3.7))
	fmt.Println("Pow(2, 3):", math.Pow(2, 3))
	fmt.Println("PI:", math.PI)
}
```

converter: enhanced `System.Convert` with comprehensive type conversion
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/converter"
)

func main() {
	// Type conversions
	intVal, _ := converter.ToInt32("42")
	boolVal, _ := converter.ToBoolean("true")
	strVal := converter.ToString(123.45)
	
	fmt.Println("String to int32:", intVal)
	fmt.Println("String to bool:", boolVal)
	fmt.Println("Float to string:", strVal)
	
	// Base64 encoding
	data := []byte("Hello, NGo!")
	encoded := converter.ToBase64String(data)
	decoded, _ := converter.FromBase64String(encoded)
	fmt.Println("Base64:", encoded)
	fmt.Println("Decoded:", string(decoded))
}
```

collections: `System.Collections.Generic.List<T>` with full generic support
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/collections/generic"
)

func main() {
	// Create a generic list
	list := generic.NewList[int]()
	list.AddRange([]int{5, 2, 8, 1, 9})
	
	fmt.Println("Original:", list.ToArray())
	fmt.Println("Count:", list.Count())
	
	// Find operations
	found, exists := list.Find(func(x int) bool { return x > 5 })
	fmt.Printf("First > 5: %d (exists: %t)\n", found, exists)
	
	// Sort
	list.Sort(func(a, b int) int {
		if a < b { return -1 }
		if a > b { return 1 }
		return 0
	})
	fmt.Println("Sorted:", list.ToArray())
}
```

datetime: enhanced `System.DateTime` with comprehensive date/time operations
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/datetime"
	"github.com/utils-go/ngo/timespan"
)

func main() {
	// Current date and time
	now := datetime.Now()
	fmt.Println("Now:", now.ToString())
	
	// Create specific date
	dt := datetime.NewDateTime(2023, 12, 25, 15, 30, 45, 0)
	fmt.Println("Christmas:", dt.ToString())
	
	// Date arithmetic
	future := dt.AddDays(10).AddHours(5)
	fmt.Println("10 days + 5 hours later:", future.ToString())
	
	// Parse from string
	parsed, _ := datetime.Parse("2023-06-15 14:30:00")
	fmt.Println("Parsed:", parsed.ToString())
	
	// Custom formatting
	formatted := dt.ToStringWithFormat("yyyy-MM-dd HH:mm:ss")
	fmt.Println("Formatted:", formatted)
}
```

timespan: enhanced `System.TimeSpan` with time interval operations
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/timespan"
)

func main() {
	// Create TimeSpan
	ts := timespan.NewTimeSpan(1, 2, 30, 45, 500) // 1 day, 2h, 30m, 45s, 500ms
	fmt.Println("TimeSpan:", ts.String())
	fmt.Printf("Total hours: %.2f\n", ts.TotalHours())
	
	// Create from different units
	hours := timespan.FromHours(3.5)
	minutes := timespan.FromMinutes(90)
	
	// Arithmetic
	sum := hours.Add(minutes)
	fmt.Println("3.5h + 90m =", sum.String())
	
	// Parse from string
	parsed, _ := timespan.Parse("2.15:30:45") // 2 days, 15h, 30m, 45s
	fmt.Println("Parsed:", parsed.String())
	
	// Go duration compatibility
	goDuration, _ := timespan.Parse("2h30m45s")
	fmt.Println("Go duration:", goDuration.String())
}
```

text: `System.Text.StringBuilder` for efficient string building
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/text"
)

func main() {
	sb := text.NewStringBuilder()
	
	// Chain operations
	sb.AppendString("Hello").AppendString(" ").AppendString("World")
	sb.AppendLine().AppendLineString("New line")
	sb.AppendFormat("Number: %d", 42)
	
	fmt.Println("Result:", sb.ToString())
	fmt.Printf("Length: %d\n", sb.Length())
	
	// String manipulation
	sb.Insert(6, "Beautiful ")
	sb.Replace("World", "Go")
	fmt.Println("Modified:", sb.ToString())
}
```

dictionary: `System.Collections.Generic.Dictionary<K,V>` with full generic support
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/collections/generic"
)

func main() {
	dict := generic.NewDictionary[string, int]()
	
	// Add items
	dict.Add("apple", 5)
	dict.Set("banana", 3) // Set can add or update
	
	// Access items
	if value, exists := dict.TryGetValue("apple"); exists {
		fmt.Printf("Apple: %d\n", value)
	}
	
	// Dictionary operations
	fmt.Printf("Keys: %v\n", dict.Keys())
	fmt.Printf("Count: %d\n", dict.Count())
	
	// LINQ-style operations
	expensive := dict.Where(func(k string, v int) bool { return v > 4 })
	fmt.Printf("Expensive items: %s\n", expensive.String())
}
```

linq: LINQ-style functional programming for Go slices
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/linq"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	enum := linq.From(numbers)
	
	// Chain operations
	result := enum.Where(func(x int) bool { return x%2 == 0 }).
		Take(3).
		ToSlice()
	fmt.Printf("First 3 even numbers: %v\n", result)
	
	// Transformations
	squares := linq.Select(enum, func(x int) int { return x * x })
	fmt.Printf("Squares: %v\n", squares.ToSlice())
	
	// Aggregations
	sum := linq.Aggregate(enum, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("Sum: %d\n", sum)
	
	// Conditions
	hasLarge := enum.Any(func(x int) bool { return x > 8 })
	fmt.Printf("Has numbers > 8: %t\n", hasLarge)
}
```

environment: `System.Environment` for system information and environment variables
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/environment"
)

func main() {
	// System information
	fmt.Printf("Machine: %s\n", environment.MachineName())
	fmt.Printf("OS: %s\n", environment.OSVersion())
	fmt.Printf("User: %s\n", environment.UserName())
	fmt.Printf("CPUs: %d\n", environment.ProcessorCount())
	
	// Environment variables
	path := environment.GetEnvironmentVariable("PATH")
	fmt.Printf("PATH exists: %t\n", path != "")
	
	// Special folders
	desktop := environment.GetFolderPath(environment.Desktop)
	fmt.Printf("Desktop: %s\n", desktop)
	
	// Current directory
	fmt.Printf("Current dir: %s\n", environment.CurrentDirectory())
}
```

diagnostics: `System.Diagnostics.Stopwatch` for performance measurement
```go
package main

import (
	"fmt"
	"time"
	"github.com/utils-go/ngo/diagnostics"
)

func main() {
	// Start timing
	sw := diagnostics.StartNew()
	
	// Simulate work
	time.Sleep(100 * time.Millisecond)
	
	sw.Stop()
	fmt.Printf("Operation took: %d ms\n", sw.ElapsedMilliseconds())
	fmt.Printf("Elapsed: %v\n", sw.Elapsed())
	
	// Restart for another measurement
	sw.Restart()
	time.Sleep(50 * time.Millisecond)
	sw.Stop()
	
	fmt.Printf("Second operation: %d ms\n", sw.ElapsedMilliseconds())
}
```

encoding: `System.Text.Encoding` for character encoding conversion
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/text"
)

func main() {
	// International text with various scripts
	text := "Hello, 世界! 🌍 Привет мир!"
	
	// UTF-8 encoding (Go default)
	utf8Bytes := text.UTF8.GetBytes(text)
	fmt.Printf("UTF-8: %d bytes\n", len(utf8Bytes))
	
	// UTF-16 encoding
	utf16Bytes := text.Unicode.GetBytes(text)
	fmt.Printf("UTF-16: %d bytes\n", len(utf16Bytes))
	
	// ASCII encoding (replaces non-ASCII with '?')
	asciiBytes := text.ASCII.GetBytes("Hello, World!")
	fmt.Printf("ASCII: %d bytes\n", len(asciiBytes))
	
	// Encoding conversion
	converted, _ := text.Convert(text.UTF8, text.Unicode, utf8Bytes)
	fmt.Printf("Converted UTF-8 to UTF-16: %d bytes\n", len(converted))
	
	// Get encoding by name or code page
	encoding, _ := text.GetEncoding("utf-8")
	fmt.Printf("Encoding: %s (CP: %d)\n", 
		encoding.EncodingName(), encoding.CodePage())
}
```

reflection: `System.Reflection` for runtime type inspection and manipulation
```go
package main

import (
	"fmt"
	"github.com/utils-go/ngo/reflection"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
	person := &Person{Name: "John", Age: 30}
	
	// Get type information
	personType := reflection.GetType(person)
	fmt.Printf("Type: %s\n", personType.Name())
	fmt.Printf("Is pointer: %t\n", personType.IsPointer())
	
	// Get fields
	fields := personType.GetFields()
	for _, field := range fields {
		value, _ := field.GetValue(person)
		fmt.Printf("Field %s: %v\n", field.Name, value)
	}
	
	// Modify field
	nameField := personType.GetField("Name")
	nameField.SetValue(person, "Jane")
	fmt.Printf("Changed name to: %s\n", person.Name)
	
	// Invoke method
	method := personType.GetMethod("GetInfo")
	results, _ := method.Invoke(person)
	fmt.Printf("Method result: %s\n", results[0])
	
	// Create new instance
	newInstance, _ := personType.CreateInstance()
	fmt.Printf("Created: %T\n", newInstance)
}
```


# Finish:
- File
- Path
- BitConverter
- Directory
- DateTime (Enhanced)
- TimeSpan (Enhanced)
- Console
- String
- Math
- Convert
- Collections.Generic.List<T>
- Text.StringBuilder
- Collections.Generic.Dictionary<K,V>
- Linq (Basic Methods)
- Environment
- Diagnostics.Stopwatch
- Text.Encoding
- Reflection

# Planned (High Priority):
## System.String
String manipulation methods that .NET developers are familiar with:
- Split, Join, Replace, Trim, TrimStart, TrimEnd
- Contains, StartsWith, EndsWith, IndexOf, LastIndexOf
- ToUpper, ToLower, Substring
- Format, PadLeft, PadRight
- IsNullOrEmpty, IsNullOrWhiteSpace

## System.Math  
Mathematical functions and constants:
- Basic: Abs, Max, Min, Round, Ceiling, Floor, Truncate
- Power: Pow, Sqrt, Exp, Log, Log10
- Trigonometric: Sin, Cos, Tan, Asin, Acos, Atan, Atan2
- Constants: PI, E

## System.Convert (Enhanced)
Extended type conversion utilities:
- ToInt32, ToInt64, ToDouble, ToSingle, ToBoolean
- ToString with format providers
- ToDateTime, ToChar, ToByte
- Base64 encoding/decoding
- ChangeType for dynamic conversions

## System.Collections.Generic.List<T>
Generic dynamic array with .NET-style methods:
- Add, AddRange, Insert, InsertRange
- Remove, RemoveAt, RemoveAll, RemoveRange
- Contains, IndexOf, LastIndexOf, Find, FindAll
- Sort, Reverse, ForEach
- ToArray, Count, Capacity
