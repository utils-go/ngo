package main

import (
	"fmt"
	"log"

	"github.com/utils-go/ngo/reflection"
	"github.com/utils-go/ngo/text"
)

// Sample structures for reflection demonstration
type Person struct {
	Name    string
	Age     int
	Email   string
	IsAdmin bool
}

func (p *Person) GetFullInfo() string {
	return fmt.Sprintf("%s (%d years old) - %s", p.Name, p.Age, p.Email)
}

func (p *Person) SetAge(newAge int) {
	p.Age = newAge
}

func (p *Person) IsAdult() bool {
	return p.Age >= 18
}

type Company struct {
	Name      string
	Employees []*Person
	Founded   int
}

func (c *Company) AddEmployee(person *Person) {
	c.Employees = append(c.Employees, person)
}

func (c *Company) GetEmployeeCount() int {
	return len(c.Employees)
}

func main() {
	fmt.Println("=== NGo Text.Encoding and Reflection Examples ===\n")

	// 1. Text Encoding Examples
	fmt.Println("1. Text Encoding Operations:")
	
	// Test different encodings with international text
	originalText := "Hello, 世界! 🌍 Привет мир! مرحبا بالعالم!"
	fmt.Printf("Original text: %s\n", originalText)
	
	// UTF-8 encoding (default in Go)
	utf8Bytes := text.UTF8.GetBytes(originalText)
	fmt.Printf("UTF-8 encoded bytes: %d bytes\n", len(utf8Bytes))
	
	decoded, err := text.UTF8.GetString(utf8Bytes)
	if err != nil {
		log.Printf("UTF-8 decoding error: %v", err)
	} else {
		fmt.Printf("UTF-8 decoded: %s\n", decoded)
	}
	
	// UTF-16 encoding
	utf16Bytes := text.Unicode.GetBytes(originalText)
	fmt.Printf("UTF-16 encoded bytes: %d bytes\n", len(utf16Bytes))
	
	utf16Decoded, err := text.Unicode.GetString(utf16Bytes)
	if err != nil {
		log.Printf("UTF-16 decoding error: %v", err)
	} else {
		fmt.Printf("UTF-16 decoded: %s\n", utf16Decoded)
	}
	
	// ASCII encoding (will replace non-ASCII characters)
	asciiText := "Hello, World!"
	asciiBytes := text.ASCII.GetBytes(asciiText)
	fmt.Printf("ASCII text '%s' -> %d bytes\n", asciiText, len(asciiBytes))
	
	// Encoding conversion
	converted, err := text.Convert(text.UTF8, text.Unicode, utf8Bytes)
	if err != nil {
		log.Printf("Conversion error: %v", err)
	} else {
		fmt.Printf("Converted UTF-8 to UTF-16: %d bytes\n", len(converted))
	}
	
	// Get encoding by name
	encoding, err := text.GetEncoding("utf-8")
	if err != nil {
		log.Printf("Get encoding error: %v", err)
	} else {
		fmt.Printf("Retrieved encoding: %s (Code Page: %d)\n", 
			encoding.EncodingName(), encoding.CodePage())
	}
	fmt.Println()

	// 2. Reflection Examples
	fmt.Println("2. Reflection Operations:")
	
	// Create sample objects
	person := &Person{
		Name:    "John Doe",
		Age:     30,
		Email:   "john.doe@example.com",
		IsAdmin: false,
	}
	
	company := &Company{
		Name:    "Tech Corp",
		Founded: 2010,
	}
	company.AddEmployee(person)
	
	// Get type information
	personType := reflection.GetType(person)
	fmt.Printf("Person type: %s\n", personType.Name())
	fmt.Printf("Full name: %s\n", personType.FullName())
	fmt.Printf("Is class: %t\n", personType.IsClass())
	fmt.Printf("Is pointer: %t\n", personType.IsPointer())
	fmt.Printf("Is value type: %t\n", personType.IsValueType())
	
	// Get and display fields
	fmt.Println("\nPerson Fields:")
	fields := personType.GetFields()
	for _, field := range fields {
		value, err := field.GetValue(person)
		if err != nil {
			log.Printf("Error getting field %s: %v", field.Name, err)
			continue
		}
		fmt.Printf("  %s (%s): %v\n", field.Name, field.FieldType.Name(), value)
	}
	
	// Modify a field using reflection
	nameField := personType.GetField("Name")
	if nameField != nil {
		err := nameField.SetValue(person, "Jane Smith")
		if err != nil {
			log.Printf("Error setting Name field: %v", err)
		} else {
			fmt.Printf("Changed name to: %s\n", person.Name)
		}
	}
	
	// Get and invoke methods
	fmt.Println("\nPerson Methods:")
	methods := personType.GetMethods()
	for _, method := range methods {
		fmt.Printf("  %s (", method.Name)
		for i, param := range method.Parameters {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s %s", param.Name, param.ParameterType.Name())
		}
		fmt.Print(")")
		if method.ReturnType != nil {
			fmt.Printf(" -> %s", method.ReturnType.Name())
		}
		fmt.Println()
	}
	
	// Invoke a method
	getFullInfoMethod := personType.GetMethod("GetFullInfo")
	if getFullInfoMethod != nil {
		results, err := getFullInfoMethod.Invoke(person)
		if err != nil {
			log.Printf("Error invoking GetFullInfo: %v", err)
		} else if len(results) > 0 {
			fmt.Printf("GetFullInfo result: %s\n", results[0])
		}
	}
	
	// Invoke method with parameters
	setAgeMethod := personType.GetMethod("SetAge")
	if setAgeMethod != nil {
		_, err := setAgeMethod.Invoke(person, 35)
		if err != nil {
			log.Printf("Error invoking SetAge: %v", err)
		} else {
			fmt.Printf("Age changed to: %d\n", person.Age)
		}
	}
	
	// Check boolean method
	isAdultMethod := personType.GetMethod("IsAdult")
	if isAdultMethod != nil {
		results, err := isAdultMethod.Invoke(person)
		if err != nil {
			log.Printf("Error invoking IsAdult: %v", err)
		} else if len(results) > 0 {
			fmt.Printf("Is adult: %v\n", results[0])
		}
	}
	
	// Create instance using reflection
	fmt.Println("\nCreating new instance using reflection:")
	newPersonType := reflection.GetType(Person{})
	newInstance, err := newPersonType.CreateInstance()
	if err != nil {
		log.Printf("Error creating instance: %v", err)
	} else {
		fmt.Printf("Created new instance: %T\n", newInstance)
		
		// Set fields on new instance
		if newPerson, ok := newInstance.(Person); ok {
			newPersonPtr := &newPerson
			nameField := newPersonType.GetField("Name")
			if nameField != nil {
				nameField.SetValue(newPersonPtr, "Bob Johnson")
				fmt.Printf("New person name: %s\n", newPerson.Name)
			}
		}
	}
	fmt.Println()

	// 3. Combined Example: Serialization using Reflection and Encoding
	fmt.Println("3. Combined Example - Object Serialization:")
	
	// Simple JSON-like serialization using reflection
	serialized := serializeObject(person)
	fmt.Printf("Serialized person: %s\n", serialized)
	
	// Encode the serialized string in different formats
	serializedBytes := text.UTF8.GetBytes(serialized)
	
	// Convert to UTF-16
	utf16Serialized, _ := text.Convert(text.UTF8, text.Unicode, serializedBytes)
	fmt.Printf("UTF-16 serialized size: %d bytes\n", len(utf16Serialized))
	
	// Convert to UTF-32
	utf32Serialized, _ := text.Convert(text.UTF8, text.UTF32, serializedBytes)
	fmt.Printf("UTF-32 serialized size: %d bytes\n", len(utf32Serialized))
	
	// Demonstrate encoding efficiency
	fmt.Printf("Encoding efficiency comparison:\n")
	fmt.Printf("  UTF-8:  %d bytes\n", len(serializedBytes))
	fmt.Printf("  UTF-16: %d bytes (%.1fx)\n", len(utf16Serialized), 
		float64(len(utf16Serialized))/float64(len(serializedBytes)))
	fmt.Printf("  UTF-32: %d bytes (%.1fx)\n", len(utf32Serialized), 
		float64(len(utf32Serialized))/float64(len(serializedBytes)))
	
	fmt.Println("\n=== Text.Encoding and Reflection Working Perfectly! ===")
}

// Simple object serialization using reflection
func serializeObject(obj interface{}) string {
	if obj == nil {
		return "null"
	}
	
	objType := reflection.GetType(obj)
	if objType == nil {
		return "null"
	}
	
	result := "{"
	fields := objType.GetFields()
	
	for i, field := range fields {
		if i > 0 {
			result += ", "
		}
		
		value, err := field.GetValue(obj)
		if err != nil {
			continue
		}
		
		result += fmt.Sprintf(`"%s": `, field.Name)
		
		switch v := value.(type) {
		case string:
			result += fmt.Sprintf(`"%s"`, v)
		case bool:
			result += fmt.Sprintf("%t", v)
		default:
			result += fmt.Sprintf("%v", v)
		}
	}
	
	result += "}"
	return result
}