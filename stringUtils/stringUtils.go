package stringUtils

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Length returns the number of characters in the string
func Length(s string) int {
	return len([]rune(s))
}

// Split divides a string into substrings based on specified delimiters
func Split(s, separator string) []string {
	if separator == "" {
		return []string{s}
	}
	return strings.Split(s, separator)
}

// SplitWithOptions splits a string with additional options
func SplitWithOptions(s, separator string, removeEmptyEntries bool) []string {
	if separator == "" {
		return []string{s}
	}

	parts := strings.Split(s, separator)
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
func Replace(s, oldValue, newValue string) string {
	return strings.ReplaceAll(s, oldValue, newValue)
}

// Trim removes leading and trailing whitespace characters
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// TrimStart removes leading whitespace characters
func TrimStart(s string) string {
	return strings.TrimLeftFunc(s, unicode.IsSpace)
}

// TrimEnd removes trailing whitespace characters
func TrimEnd(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// TrimChars removes leading and trailing occurrences of specified characters
func TrimChars(s, chars string) string {
	return strings.Trim(s, chars)
}

// Contains determines whether the string contains the specified substring
func Contains(s, value string) bool {
	return strings.Contains(s, value)
}

// StartsWith determines whether the string starts with the specified prefix
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// EndsWith determines whether the string ends with the specified suffix
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// IndexOf returns the zero-based index of the first occurrence of the specified string
func IndexOf(s, value string) int {
	index := strings.Index(s, value)
	if index == -1 {
		return -1
	}
	// Convert byte index to rune index
	return len([]rune(s[:index]))
}

// LastIndexOf returns the zero-based index of the last occurrence of the specified string
func LastIndexOf(s, value string) int {
	index := strings.LastIndex(s, value)
	if index == -1 {
		return -1
	}
	// Convert byte index to rune index
	return len([]rune(s[:index]))
}

// ToUpper returns a copy of the string converted to uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower returns a copy of the string converted to lowercase
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Substring returns a substring starting at the specified index
func Substring(s string, startIndex int) string {
	runes := []rune(s)
	if startIndex < 0 || startIndex >= len(runes) {
		return ""
	}
	return string(runes[startIndex:])
}

// SubstringWithLength returns a substring starting at the specified index with specified length
func SubstringWithLength(s string, startIndex, length int) string {
	runes := []rune(s)
	if startIndex < 0 || startIndex >= len(runes) || length <= 0 {
		return ""
	}

	endIndex := startIndex + length
	if endIndex > len(runes) {
		endIndex = len(runes)
	}

	return string(runes[startIndex:endIndex])
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
func PadLeft(s string, totalWidth int) string {
	return PadLeftWithChar(s, totalWidth, ' ')
}

// PadLeftWithChar pads the string on the left with the specified character
func PadLeftWithChar(s string, totalWidth int, paddingChar rune) string {
	runes := []rune(s)
	currentLength := len(runes)

	if totalWidth <= currentLength {
		return s
	}

	padding := strings.Repeat(string(paddingChar), totalWidth-currentLength)
	return padding + s
}

// PadRight pads the string on the right with spaces to reach the specified total length
func PadRight(s string, totalWidth int) string {
	return PadRightWithChar(s, totalWidth, ' ')
}

// PadRightWithChar pads the string on the right with the specified character
func PadRightWithChar(s string, totalWidth int, paddingChar rune) string {
	runes := []rune(s)
	currentLength := len(runes)

	if totalWidth <= currentLength {
		return s
	}

	padding := strings.Repeat(string(paddingChar), totalWidth-currentLength)
	return s + padding
}

// IsNullOrEmpty determines whether the specified string is empty
func IsNullOrEmpty(s string) bool {
	return s == ""
}

// IsNullOrWhiteSpace determines whether the specified string is empty or consists only of white-space characters
func IsNullOrWhiteSpace(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Equals determines whether two strings have the same value
func Equals(s1, s2 string) bool {
	return s1 == s2
}

// EqualsIgnoreCase determines whether two strings have the same value (case-insensitive)
func EqualsIgnoreCase(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

// CompareTo compares two strings and returns an integer that indicates their relative position
func CompareTo(s1, s2 string) int {
	if s1 < s2 {
		return -1
	} else if s1 > s2 {
		return 1
	}
	return 0
}
