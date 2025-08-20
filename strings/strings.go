package strings

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// String provides .NET-style string manipulation methods for Go
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.string?view=netframework-4.7.2
type String struct {
	value string
}

// NewString creates a new String instance from a Go string
func NewString(s string) *String {
	return &String{value: s}
}

// Value returns the underlying Go string value
func (s *String) Value() string {
	return s.value
}

// String implements the Stringer interface
func (s *String) String() string {
	return s.value
}

// Length returns the number of characters in the string
func (s *String) Length() int {
	return len([]rune(s.value))
}

// Split divides a string into substrings based on specified delimiters
func (s *String) Split(separator string) []string {
	if separator == "" {
		return []string{s.value}
	}
	return strings.Split(s.value, separator)
}

// SplitWithOptions splits a string with additional options
func (s *String) SplitWithOptions(separator string, removeEmptyEntries bool) []string {
	parts := strings.Split(s.value, separator)
	if !removeEmptyEntries {
		return parts
	}
	
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

// Join concatenates the elements of a string array using the specified separator
func Join(separator string, values []string) string {
	return strings.Join(values, separator)
}

// Replace returns a new string with all occurrences of oldValue replaced by newValue
func (s *String) Replace(oldValue, newValue string) *String {
	return NewString(strings.ReplaceAll(s.value, oldValue, newValue))
}

// Trim removes leading and trailing whitespace characters
func (s *String) Trim() *String {
	return NewString(strings.TrimSpace(s.value))
}

// TrimStart removes leading whitespace characters
func (s *String) TrimStart() *String {
	return NewString(strings.TrimLeftFunc(s.value, unicode.IsSpace))
}

// TrimEnd removes trailing whitespace characters
func (s *String) TrimEnd() *String {
	return NewString(strings.TrimRightFunc(s.value, unicode.IsSpace))
}

// TrimChars removes leading and trailing occurrences of specified characters
func (s *String) TrimChars(chars string) *String {
	return NewString(strings.Trim(s.value, chars))
}

// Contains determines whether the string contains the specified substring
func (s *String) Contains(value string) bool {
	return strings.Contains(s.value, value)
}

// StartsWith determines whether the string starts with the specified prefix
func (s *String) StartsWith(prefix string) bool {
	return strings.HasPrefix(s.value, prefix)
}

// EndsWith determines whether the string ends with the specified suffix
func (s *String) EndsWith(suffix string) bool {
	return strings.HasSuffix(s.value, suffix)
}

// IndexOf returns the zero-based index of the first occurrence of the specified string
func (s *String) IndexOf(value string) int {
	index := strings.Index(s.value, value)
	if index == -1 {
		return -1
	}
	// Convert byte index to rune index
	return len([]rune(s.value[:index]))
}

// LastIndexOf returns the zero-based index of the last occurrence of the specified string
func (s *String) LastIndexOf(value string) int {
	index := strings.LastIndex(s.value, value)
	if index == -1 {
		return -1
	}
	// Convert byte index to rune index
	return len([]rune(s.value[:index]))
}

// ToUpper returns a copy of the string converted to uppercase
func (s *String) ToUpper() *String {
	return NewString(strings.ToUpper(s.value))
}

// ToLower returns a copy of the string converted to lowercase
func (s *String) ToLower() *String {
	return NewString(strings.ToLower(s.value))
}

// Substring returns a substring starting at the specified index
func (s *String) Substring(startIndex int) *String {
	runes := []rune(s.value)
	if startIndex < 0 || startIndex >= len(runes) {
		return NewString("")
	}
	return NewString(string(runes[startIndex:]))
}

// SubstringWithLength returns a substring starting at the specified index with specified length
func (s *String) SubstringWithLength(startIndex, length int) *String {
	runes := []rune(s.value)
	if startIndex < 0 || startIndex >= len(runes) || length <= 0 {
		return NewString("")
	}
	
	endIndex := startIndex + length
	if endIndex > len(runes) {
		endIndex = len(runes)
	}
	
	return NewString(string(runes[startIndex:endIndex]))
}

// Format creates a formatted string using the specified format string and arguments
func Format(format string, args ...interface{}) string {
	// Convert .NET-style format placeholders {0}, {1}, etc. to Go-style
	re := regexp.MustCompile(`\{(\d+)\}`)
	goFormat := re.ReplaceAllStringFunc(format, func(match string) string {
		return "%v"
	})
	return fmt.Sprintf(goFormat, args...)
}

// PadLeft pads the string on the left with spaces to reach the specified total length
func (s *String) PadLeft(totalWidth int) *String {
	return s.PadLeftWithChar(totalWidth, ' ')
}

// PadLeftWithChar pads the string on the left with the specified character
func (s *String) PadLeftWithChar(totalWidth int, paddingChar rune) *String {
	runes := []rune(s.value)
	currentLength := len(runes)
	
	if totalWidth <= currentLength {
		return NewString(s.value)
	}
	
	padding := strings.Repeat(string(paddingChar), totalWidth-currentLength)
	return NewString(padding + s.value)
}

// PadRight pads the string on the right with spaces to reach the specified total length
func (s *String) PadRight(totalWidth int) *String {
	return s.PadRightWithChar(totalWidth, ' ')
}

// PadRightWithChar pads the string on the right with the specified character
func (s *String) PadRightWithChar(totalWidth int, paddingChar rune) *String {
	runes := []rune(s.value)
	currentLength := len(runes)
	
	if totalWidth <= currentLength {
		return NewString(s.value)
	}
	
	padding := strings.Repeat(string(paddingChar), totalWidth-currentLength)
	return NewString(s.value + padding)
}

// IsNullOrEmpty determines whether the specified string is null or empty
func IsNullOrEmpty(s *String) bool {
	return s == nil || s.value == ""
}

// IsNullOrWhiteSpace determines whether the specified string is null, empty, or consists only of white-space characters
func IsNullOrWhiteSpace(s *String) bool {
	if s == nil {
		return true
	}
	return strings.TrimSpace(s.value) == ""
}

// Equals determines whether two String instances have the same value
func (s *String) Equals(other *String) bool {
	if s == nil && other == nil {
		return true
	}
	if s == nil || other == nil {
		return false
	}
	return s.value == other.value
}

// EqualsIgnoreCase determines whether two String instances have the same value (case-insensitive)
func (s *String) EqualsIgnoreCase(other *String) bool {
	if s == nil && other == nil {
		return true
	}
	if s == nil || other == nil {
		return false
	}
	return strings.EqualFold(s.value, other.value)
}

// CompareTo compares this instance with another String and returns an integer that indicates their relative position
func (s *String) CompareTo(other *String) int {
	if s == nil && other == nil {
		return 0
	}
	if s == nil {
		return -1
	}
	if other == nil {
		return 1
	}
	
	if s.value < other.value {
		return -1
	} else if s.value > other.value {
		return 1
	}
	return 0
}