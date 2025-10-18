package text

import (
	"bytes"
	"testing"
)

func TestUTF8Encoding(t *testing.T) {
	encoding := UTF8
	
	// Test basic string encoding/decoding
	original := "Hello, 世界! 🌍"
	encoded := encoding.GetBytes(original)
	decoded, err := encoding.GetString(encoded)
	
	if err != nil {
		t.Errorf("UTF-8 decoding failed: %v", err)
	}
	
	if decoded != original {
		t.Errorf("UTF-8 round-trip failed: expected %q, got %q", original, decoded)
	}
	
	// Test encoding properties
	if encoding.EncodingName() != "UTF-8" {
		t.Errorf("Expected encoding name 'UTF-8', got %q", encoding.EncodingName())
	}
	
	if encoding.CodePage() != 65001 {
		t.Errorf("Expected code page 65001, got %d", encoding.CodePage())
	}
	
	if encoding.IsSingleByte() {
		t.Error("UTF-8 should not be single-byte encoding")
	}
}

func TestASCIIEncoding(t *testing.T) {
	encoding := ASCII
	
	// Test valid ASCII string
	original := "Hello, World!"
	encoded := encoding.GetBytes(original)
	decoded, err := encoding.GetString(encoded)
	
	if err != nil {
		t.Errorf("ASCII decoding failed: %v", err)
	}
	
	if decoded != original {
		t.Errorf("ASCII round-trip failed: expected %q, got %q", original, decoded)
	}
	
	// Test non-ASCII characters (should be replaced with '?')
	nonAscii := "Hello, 世界!"
	encoded = encoding.GetBytes(nonAscii)
	expected := []byte("Hello, ??!")
	
	if !bytes.Equal(encoded, expected) {
		t.Errorf("ASCII encoding of non-ASCII chars failed: expected %v, got %v", expected, encoded)
	}
	
	// Test invalid ASCII bytes
	invalidBytes := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0xFF} // "Hello" + invalid byte
	_, err = encoding.GetString(invalidBytes)
	
	if err == nil {
		t.Error("Expected error when decoding invalid ASCII bytes")
	}
	
	// Test encoding properties
	if encoding.EncodingName() != "US-ASCII" {
		t.Errorf("Expected encoding name 'US-ASCII', got %q", encoding.EncodingName())
	}
	
	if encoding.CodePage() != 20127 {
		t.Errorf("Expected code page 20127, got %d", encoding.CodePage())
	}
	
	if !encoding.IsSingleByte() {
		t.Error("ASCII should be single-byte encoding")
	}
}

func TestUnicodeEncoding(t *testing.T) {
	// Test little-endian UTF-16
	encoding := Unicode
	
	original := "Hello, 世界!"
	encoded := encoding.GetBytes(original)
	decoded, err := encoding.GetString(encoded)
	
	if err != nil {
		t.Errorf("Unicode decoding failed: %v", err)
	}
	
	if decoded != original {
		t.Errorf("Unicode round-trip failed: expected %q, got %q", original, decoded)
	}
	
	// Test big-endian UTF-16
	encodingBE := BigEndianUnicode
	encodedBE := encodingBE.GetBytes(original)
	decodedBE, err := encodingBE.GetString(encodedBE)
	
	if err != nil {
		t.Errorf("Big-endian Unicode decoding failed: %v", err)
	}
	
	if decodedBE != original {
		t.Errorf("Big-endian Unicode round-trip failed: expected %q, got %q", original, decodedBE)
	}
	
	// Encoded bytes should be different between LE and BE
	if bytes.Equal(encoded, encodedBE) {
		t.Error("Little-endian and big-endian encodings should produce different bytes")
	}
	
	// Test invalid UTF-16 (odd number of bytes)
	invalidBytes := []byte{0x48, 0x00, 0x65} // Incomplete UTF-16 sequence
	_, err = encoding.GetString(invalidBytes)
	
	if err == nil {
		t.Error("Expected error when decoding incomplete UTF-16 sequence")
	}
	
	// Test encoding properties
	if encoding.EncodingName() != "UTF-16LE" {
		t.Errorf("Expected encoding name 'UTF-16LE', got %q", encoding.EncodingName())
	}
	
	if encodingBE.EncodingName() != "UTF-16BE" {
		t.Errorf("Expected encoding name 'UTF-16BE', got %q", encodingBE.EncodingName())
	}
}

func TestUTF32Encoding(t *testing.T) {
	// Test little-endian UTF-32
	encoding := UTF32
	
	original := "Hello, 🌍!"
	encoded := encoding.GetBytes(original)
	decoded, err := encoding.GetString(encoded)
	
	if err != nil {
		t.Errorf("UTF-32 decoding failed: %v", err)
	}
	
	if decoded != original {
		t.Errorf("UTF-32 round-trip failed: expected %q, got %q", original, decoded)
	}
	
	// Test big-endian UTF-32
	encodingBE := UTF32BE
	encodedBE := encodingBE.GetBytes(original)
	decodedBE, err := encodingBE.GetString(encodedBE)
	
	if err != nil {
		t.Errorf("Big-endian UTF-32 decoding failed: %v", err)
	}
	
	if decodedBE != original {
		t.Errorf("Big-endian UTF-32 round-trip failed: expected %q, got %q", original, decodedBE)
	}
	
	// Test invalid UTF-32 (not multiple of 4 bytes)
	invalidBytes := []byte{0x48, 0x00, 0x00} // Incomplete UTF-32 sequence
	_, err = encoding.GetString(invalidBytes)
	
	if err == nil {
		t.Error("Expected error when decoding incomplete UTF-32 sequence")
	}
	
	// Test invalid Unicode code point
	invalidCodePoint := []byte{0xFF, 0xFF, 0xFF, 0xFF} // Invalid Unicode code point
	_, err = encoding.GetString(invalidCodePoint)
	
	if err == nil {
		t.Error("Expected error when decoding invalid Unicode code point")
	}
}

func TestGetEncoding(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected Encoding
		hasError bool
	}{
		{"utf-8", UTF8, false},
		{"UTF-8", UTF8, false},
		{"ascii", ASCII, false},
		{"US-ASCII", ASCII, false},
		{"unicode", Unicode, false},
		{"utf-16", Unicode, false},
		{"utf-16be", BigEndianUnicode, false},
		{"utf-32", UTF32, false},
		{"utf-32be", UTF32BE, false},
		{65001, UTF8, false},
		{20127, ASCII, false},
		{1200, Unicode, false},
		{1201, BigEndianUnicode, false},
		{12000, UTF32, false},
		{12001, UTF32BE, false},
		{"unknown", nil, true},
		{99999, nil, true},
		{3.14, nil, true}, // Invalid type
	}
	
	for _, test := range tests {
		encoding, err := GetEncoding(test.input)
		
		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for input %v, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %v: %v", test.input, err)
			}
			
			if encoding != test.expected {
				t.Errorf("GetEncoding(%v) = %v, expected %v", test.input, encoding, test.expected)
			}
		}
	}
}

func TestConvert(t *testing.T) {
	original := "Hello, World!"
	
	// Convert from UTF-8 to ASCII
	utf8Bytes := UTF8.GetBytes(original)
	asciiBytes, err := Convert(UTF8, ASCII, utf8Bytes)
	
	if err != nil {
		t.Errorf("Convert UTF-8 to ASCII failed: %v", err)
	}
	
	decoded, err := ASCII.GetString(asciiBytes)
	if err != nil {
		t.Errorf("ASCII decoding after conversion failed: %v", err)
	}
	
	if decoded != original {
		t.Errorf("Convert round-trip failed: expected %q, got %q", original, decoded)
	}
	
	// Test conversion with nil encodings
	_, err = Convert(nil, ASCII, utf8Bytes)
	if err == nil {
		t.Error("Expected error when source encoding is nil")
	}
	
	_, err = Convert(UTF8, nil, utf8Bytes)
	if err == nil {
		t.Error("Expected error when destination encoding is nil")
	}
}

func TestEncodingByteCount(t *testing.T) {
	text := "Hello, 世界! 🌍"
	
	// UTF-8: variable length
	utf8Count := UTF8.GetByteCount(text)
	utf8Bytes := UTF8.GetBytes(text)
	if utf8Count != len(utf8Bytes) {
		t.Errorf("UTF-8 GetByteCount mismatch: expected %d, got %d", len(utf8Bytes), utf8Count)
	}
	
	// ASCII: 1 byte per character (for ASCII chars)
	asciiText := "Hello"
	asciiCount := ASCII.GetByteCount(asciiText)
	if asciiCount != len(asciiText) {
		t.Errorf("ASCII GetByteCount mismatch: expected %d, got %d", len(asciiText), asciiCount)
	}
	
	// UTF-16: 2 bytes per code unit (may be more for surrogate pairs)
	unicodeCount := Unicode.GetByteCount(text)
	unicodeBytes := Unicode.GetBytes(text)
	if unicodeCount != len(unicodeBytes) {
		t.Errorf("Unicode GetByteCount mismatch: expected %d, got %d", len(unicodeBytes), unicodeCount)
	}
	
	// UTF-32: 4 bytes per character
	utf32Count := UTF32.GetByteCount(text)
	utf32Bytes := UTF32.GetBytes(text)
	if utf32Count != len(utf32Bytes) {
		t.Errorf("UTF-32 GetByteCount mismatch: expected %d, got %d", len(utf32Bytes), utf32Count)
	}
}

func TestEncodingCharCount(t *testing.T) {
	text := "Hello, 世界! 🌍"
	runes := []rune(text)
	expectedCharCount := len(runes)
	
	// Test UTF-8
	utf8Bytes := UTF8.GetBytes(text)
	utf8CharCount := UTF8.GetCharCount(utf8Bytes)
	if utf8CharCount != expectedCharCount {
		t.Errorf("UTF-8 GetCharCount mismatch: expected %d, got %d", expectedCharCount, utf8CharCount)
	}
	
	// Test ASCII (only for ASCII text)
	asciiText := "Hello"
	asciiBytes := ASCII.GetBytes(asciiText)
	asciiCharCount := ASCII.GetCharCount(asciiBytes)
	if asciiCharCount != len(asciiText) {
		t.Errorf("ASCII GetCharCount mismatch: expected %d, got %d", len(asciiText), asciiCharCount)
	}
	
	// Test UTF-16
	unicodeBytes := Unicode.GetBytes(text)
	unicodeCharCount := Unicode.GetCharCount(unicodeBytes)
	if unicodeCharCount != expectedCharCount {
		t.Errorf("Unicode GetCharCount mismatch: expected %d, got %d", expectedCharCount, unicodeCharCount)
	}
	
	// Test UTF-32
	utf32Bytes := UTF32.GetBytes(text)
	utf32CharCount := UTF32.GetCharCount(utf32Bytes)
	if utf32CharCount != expectedCharCount {
		t.Errorf("UTF-32 GetCharCount mismatch: expected %d, got %d", expectedCharCount, utf32CharCount)
	}
}

func TestEncodingFromRunes(t *testing.T) {
	runes := []rune("Hello, 世界! 🌍")
	
	// Test all encodings
	encodings := []Encoding{UTF8, ASCII, Unicode, BigEndianUnicode, UTF32, UTF32BE}
	
	for _, encoding := range encodings {
		bytes := encoding.GetBytesFromRunes(runes)
		chars, err := encoding.GetChars(bytes)
		
		if encoding == ASCII {
			// ASCII will replace non-ASCII characters, so we expect different results
			continue
		}
		
		if err != nil {
			t.Errorf("%s GetChars failed: %v", encoding.EncodingName(), err)
			continue
		}
		
		if len(chars) != len(runes) {
			t.Errorf("%s rune round-trip length mismatch: expected %d, got %d", 
				encoding.EncodingName(), len(runes), len(chars))
			continue
		}
		
		for i, expected := range runes {
			if i < len(chars) && chars[i] != expected {
				t.Errorf("%s rune round-trip failed at position %d: expected %U, got %U", 
					encoding.EncodingName(), i, expected, chars[i])
				break
			}
		}
	}
}

func TestEncodingError(t *testing.T) {
	err := &EncodingError{
		Operation: "test operation",
		Position:  5,
		Message:   "test message",
	}
	
	expected := "encoding error at position 5 during test operation: test message"
	if err.Error() != expected {
		t.Errorf("EncodingError.Error() = %q, expected %q", err.Error(), expected)
	}
}

func BenchmarkUTF8Encoding(b *testing.B) {
	text := "Hello, 世界! This is a test string with mixed ASCII and Unicode characters. 🌍🚀💻"
	
	b.Run("GetBytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			UTF8.GetBytes(text)
		}
	})
	
	bytes := UTF8.GetBytes(text)
	b.Run("GetString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			UTF8.GetString(bytes)
		}
	})
}

func BenchmarkUnicodeEncoding(b *testing.B) {
	text := "Hello, 世界! This is a test string with mixed ASCII and Unicode characters. 🌍🚀💻"
	
	b.Run("GetBytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Unicode.GetBytes(text)
		}
	})
	
	bytes := Unicode.GetBytes(text)
	b.Run("GetString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Unicode.GetString(bytes)
		}
	})
}