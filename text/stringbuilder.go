package text

import (
	"fmt"
	"strings"
)

// StringBuilder represents a mutable string of characters equivalent to System.Text.StringBuilder in .NET
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.text.stringbuilder?view=netframework-4.7.2
type StringBuilder struct {
	builder strings.Builder
}

// NewStringBuilder creates a new StringBuilder instance
func NewStringBuilder() *StringBuilder {
	return &StringBuilder{
		builder: strings.Builder{},
	}
}

// NewStringBuilderWithCapacity creates a new StringBuilder instance with the specified initial capacity
func NewStringBuilderWithCapacity(capacity int) *StringBuilder {
	sb := &StringBuilder{
		builder: strings.Builder{},
	}
	sb.builder.Grow(capacity)
	return sb
}

// NewStringBuilderFromString creates a new StringBuilder instance initialized with the specified string
func NewStringBuilderFromString(value string) *StringBuilder {
	sb := &StringBuilder{
		builder: strings.Builder{},
	}
	sb.builder.WriteString(value)
	return sb
}

// Length gets the length of the current StringBuilder object
func (sb *StringBuilder) Length() int {
	return sb.builder.Len()
}

// Capacity gets the maximum number of characters that can be contained in the memory allocated by the current instance
func (sb *StringBuilder) Capacity() int {
	return sb.builder.Cap()
}

// Append appends the string representation of a specified object to this instance
func (sb *StringBuilder) Append(value interface{}) *StringBuilder {
	if value == nil {
		return sb
	}
	
	switch v := value.(type) {
	case string:
		sb.builder.WriteString(v)
	case []byte:
		sb.builder.Write(v)
	case rune:
		sb.builder.WriteRune(v)
	case byte:
		sb.builder.WriteByte(v)
	default:
		sb.builder.WriteString(fmt.Sprintf("%v", value))
	}
	return sb
}

// AppendString appends a copy of the specified string to this instance
func (sb *StringBuilder) AppendString(value string) *StringBuilder {
	sb.builder.WriteString(value)
	return sb
}

// AppendRune appends the string representation of a specified Unicode character to this instance
func (sb *StringBuilder) AppendRune(value rune) *StringBuilder {
	sb.builder.WriteRune(value)
	return sb
}

// AppendByte appends the string representation of a specified byte to this instance
func (sb *StringBuilder) AppendByte(value byte) *StringBuilder {
	sb.builder.WriteByte(value)
	return sb
}

// AppendLine appends the default line terminator to the end of the current StringBuilder object
func (sb *StringBuilder) AppendLine() *StringBuilder {
	sb.builder.WriteString("\n")
	return sb
}

// AppendLineString appends a copy of the specified string followed by the default line terminator
func (sb *StringBuilder) AppendLineString(value string) *StringBuilder {
	sb.builder.WriteString(value)
	sb.builder.WriteString("\n")
	return sb
}

// AppendFormat appends the string returned by processing a composite format string
func (sb *StringBuilder) AppendFormat(format string, args ...interface{}) *StringBuilder {
	formatted := fmt.Sprintf(format, args...)
	sb.builder.WriteString(formatted)
	return sb
}

// Insert inserts the string representation of a specified object at the specified character position
func (sb *StringBuilder) Insert(index int, value interface{}) *StringBuilder {
	if index < 0 || index > sb.Length() {
		panic(fmt.Sprintf("index %d is out of range", index))
	}
	
	current := sb.ToString()
	valueStr := fmt.Sprintf("%v", value)
	
	// Rebuild the string with insertion
	sb.Clear()
	sb.builder.WriteString(current[:index])
	sb.builder.WriteString(valueStr)
	sb.builder.WriteString(current[index:])
	
	return sb
}

// Remove removes the specified range of characters from this instance
func (sb *StringBuilder) Remove(startIndex, length int) *StringBuilder {
	if startIndex < 0 || startIndex >= sb.Length() {
		panic(fmt.Sprintf("startIndex %d is out of range", startIndex))
	}
	if length < 0 || startIndex+length > sb.Length() {
		panic(fmt.Sprintf("length %d is invalid", length))
	}
	
	current := sb.ToString()
	
	// Rebuild the string without the removed part
	sb.Clear()
	sb.builder.WriteString(current[:startIndex])
	sb.builder.WriteString(current[startIndex+length:])
	
	return sb
}

// Replace replaces all occurrences of a specified string in this instance with another specified string
func (sb *StringBuilder) Replace(oldValue, newValue string) *StringBuilder {
	current := sb.ToString()
	replaced := strings.ReplaceAll(current, oldValue, newValue)
	
	sb.Clear()
	sb.builder.WriteString(replaced)
	
	return sb
}

// ReplaceRange replaces all occurrences of a specified character in a substring of this instance
func (sb *StringBuilder) ReplaceRange(oldValue, newValue string, startIndex, count int) *StringBuilder {
	if startIndex < 0 || startIndex >= sb.Length() {
		panic(fmt.Sprintf("startIndex %d is out of range", startIndex))
	}
	if count < 0 || startIndex+count > sb.Length() {
		panic(fmt.Sprintf("count %d is invalid", count))
	}
	
	current := sb.ToString()
	
	// Extract the substring to replace in
	before := current[:startIndex]
	middle := current[startIndex : startIndex+count]
	after := current[startIndex+count:]
	
	// Replace in the middle part only
	replacedMiddle := strings.ReplaceAll(middle, oldValue, newValue)
	
	// Rebuild the string
	sb.Clear()
	sb.builder.WriteString(before)
	sb.builder.WriteString(replacedMiddle)
	sb.builder.WriteString(after)
	
	return sb
}

// Clear removes all characters from the current StringBuilder instance
func (sb *StringBuilder) Clear() *StringBuilder {
	sb.builder.Reset()
	return sb
}

// ToString converts the value of this instance to a String
func (sb *StringBuilder) ToString() string {
	return sb.builder.String()
}

// ToStringRange converts the value of a substring of this instance to a String
func (sb *StringBuilder) ToStringRange(startIndex, length int) string {
	if startIndex < 0 || startIndex >= sb.Length() {
		panic(fmt.Sprintf("startIndex %d is out of range", startIndex))
	}
	if length < 0 || startIndex+length > sb.Length() {
		panic(fmt.Sprintf("length %d is invalid", length))
	}
	
	current := sb.ToString()
	return current[startIndex : startIndex+length]
}

// EnsureCapacity ensures that the capacity of this instance of StringBuilder is at least the specified value
func (sb *StringBuilder) EnsureCapacity(capacity int) int {
	if capacity > sb.Capacity() {
		sb.builder.Grow(capacity - sb.Capacity())
	}
	return sb.Capacity()
}

// Equals determines whether this instance and another specified StringBuilder object have the same value
func (sb *StringBuilder) Equals(other *StringBuilder) bool {
	if other == nil {
		return false
	}
	return sb.ToString() == other.ToString()
}

// String returns a string representation of the StringBuilder
func (sb *StringBuilder) String() string {
	return sb.ToString()
}

// CopyTo copies the characters from a specified segment of this instance to a specified segment of a destination array
func (sb *StringBuilder) CopyTo(sourceIndex int, destination []rune, destinationIndex, count int) {
	if sourceIndex < 0 || sourceIndex >= sb.Length() {
		panic(fmt.Sprintf("sourceIndex %d is out of range", sourceIndex))
	}
	if destinationIndex < 0 || destinationIndex >= len(destination) {
		panic(fmt.Sprintf("destinationIndex %d is out of range", destinationIndex))
	}
	if count < 0 || sourceIndex+count > sb.Length() || destinationIndex+count > len(destination) {
		panic(fmt.Sprintf("count %d is invalid", count))
	}
	
	source := []rune(sb.ToString())
	copy(destination[destinationIndex:destinationIndex+count], source[sourceIndex:sourceIndex+count])
}

// GetCharAt gets the character at the specified character position in this instance
func (sb *StringBuilder) GetCharAt(index int) rune {
	if index < 0 || index >= sb.Length() {
		panic(fmt.Sprintf("index %d is out of range", index))
	}
	
	runes := []rune(sb.ToString())
	return runes[index]
}

// SetCharAt sets the character at the specified character position in this instance
func (sb *StringBuilder) SetCharAt(index int, value rune) {
	if index < 0 || index >= sb.Length() {
		panic(fmt.Sprintf("index %d is out of range", index))
	}
	
	runes := []rune(sb.ToString())
	runes[index] = value
	
	sb.Clear()
	sb.builder.WriteString(string(runes))
}