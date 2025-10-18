package text

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// Encoding represents a character encoding equivalent to System.Text.Encoding in .NET
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.text.encoding?view=netframework-4.7.2
type Encoding interface {
	// GetBytes converts the specified string into a sequence of bytes
	GetBytes(s string) []byte
	
	// GetBytesFromRunes converts the specified rune slice into a sequence of bytes
	GetBytesFromRunes(chars []rune) []byte
	
	// GetString decodes a sequence of bytes into a string
	GetString(bytes []byte) (string, error)
	
	// GetChars decodes a sequence of bytes into a rune slice
	GetChars(bytes []byte) ([]rune, error)
	
	// GetByteCount returns the number of bytes produced by encoding the specified string
	GetByteCount(s string) int
	
	// GetCharCount returns the number of characters produced by decoding the specified byte sequence
	GetCharCount(bytes []byte) int
	
	// EncodingName gets the human-readable description of the current encoding
	EncodingName() string
	
	// CodePage gets the code page identifier of the current encoding
	CodePage() int
	
	// IsSingleByte gets a value indicating whether the current encoding uses single-byte code points
	IsSingleByte() bool
}

// EncodingError represents an encoding/decoding error
type EncodingError struct {
	Operation string
	Position  int
	Message   string
}

func (e *EncodingError) Error() string {
	return fmt.Sprintf("encoding error at position %d during %s: %s", e.Position, e.Operation, e.Message)
}

// Common encoding instances
var (
	// UTF8 represents UTF-8 encoding
	UTF8 Encoding = &utf8Encoding{}
	
	// ASCII represents ASCII encoding
	ASCII Encoding = &asciiEncoding{}
	
	// Unicode represents UTF-16 little-endian encoding
	Unicode Encoding = &unicodeEncoding{littleEndian: true}
	
	// BigEndianUnicode represents UTF-16 big-endian encoding
	BigEndianUnicode Encoding = &unicodeEncoding{littleEndian: false}
	
	// UTF32 represents UTF-32 little-endian encoding
	UTF32 Encoding = &utf32Encoding{littleEndian: true}
	
	// UTF32BE represents UTF-32 big-endian encoding
	UTF32BE Encoding = &utf32Encoding{littleEndian: false}
)

// GetEncoding returns an encoding associated with the specified code page identifier or name
func GetEncoding(nameOrCodePage interface{}) (Encoding, error) {
	switch v := nameOrCodePage.(type) {
	case string:
		return getEncodingByName(strings.ToLower(v))
	case int:
		return getEncodingByCodePage(v)
	default:
		return nil, fmt.Errorf("invalid encoding identifier type: %T", nameOrCodePage)
	}
}

func getEncodingByName(name string) (Encoding, error) {
	switch name {
	case "utf-8", "utf8":
		return UTF8, nil
	case "ascii", "us-ascii":
		return ASCII, nil
	case "unicode", "utf-16", "utf-16le":
		return Unicode, nil
	case "utf-16be":
		return BigEndianUnicode, nil
	case "utf-32", "utf-32le":
		return UTF32, nil
	case "utf-32be":
		return UTF32BE, nil
	default:
		return nil, fmt.Errorf("unknown encoding name: %s", name)
	}
}

func getEncodingByCodePage(codePage int) (Encoding, error) {
	switch codePage {
	case 65001:
		return UTF8, nil
	case 20127:
		return ASCII, nil
	case 1200:
		return Unicode, nil
	case 1201:
		return BigEndianUnicode, nil
	case 12000:
		return UTF32, nil
	case 12001:
		return UTF32BE, nil
	default:
		return nil, fmt.Errorf("unsupported code page: %d", codePage)
	}
}

// Convert converts bytes from one encoding to another
func Convert(srcEncoding, dstEncoding Encoding, bytes []byte) ([]byte, error) {
	if srcEncoding == nil || dstEncoding == nil {
		return nil, errors.New("encoding cannot be nil")
	}
	
	// Decode from source encoding
	str, err := srcEncoding.GetString(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode from source encoding: %w", err)
	}
	
	// Encode to destination encoding
	return dstEncoding.GetBytes(str), nil
}

// UTF-8 Encoding Implementation
type utf8Encoding struct{}

func (e *utf8Encoding) GetBytes(s string) []byte {
	return []byte(s)
}

func (e *utf8Encoding) GetBytesFromRunes(chars []rune) []byte {
	return []byte(string(chars))
}

func (e *utf8Encoding) GetString(bytes []byte) (string, error) {
	if !utf8.Valid(bytes) {
		return "", &EncodingError{
			Operation: "UTF-8 decoding",
			Position:  0,
			Message:   "invalid UTF-8 sequence",
		}
	}
	return string(bytes), nil
}

func (e *utf8Encoding) GetChars(bytes []byte) ([]rune, error) {
	str, err := e.GetString(bytes)
	if err != nil {
		return nil, err
	}
	return []rune(str), nil
}

func (e *utf8Encoding) GetByteCount(s string) int {
	return len([]byte(s))
}

func (e *utf8Encoding) GetCharCount(bytes []byte) int {
	return utf8.RuneCount(bytes)
}

func (e *utf8Encoding) EncodingName() string {
	return "UTF-8"
}

func (e *utf8Encoding) CodePage() int {
	return 65001
}

func (e *utf8Encoding) IsSingleByte() bool {
	return false
}

// ASCII Encoding Implementation
type asciiEncoding struct{}

func (e *asciiEncoding) GetBytes(s string) []byte {
	result := make([]byte, 0, len(s))
	for i, r := range s {
		if r > 127 {
			// Replace non-ASCII characters with '?'
			result = append(result, '?')
		} else {
			result = append(result, byte(r))
		}
		_ = i // Avoid unused variable warning
	}
	return result
}

func (e *asciiEncoding) GetBytesFromRunes(chars []rune) []byte {
	result := make([]byte, len(chars))
	for i, r := range chars {
		if r > 127 {
			result[i] = '?'
		} else {
			result[i] = byte(r)
		}
	}
	return result
}

func (e *asciiEncoding) GetString(bytes []byte) (string, error) {
	for i, b := range bytes {
		if b > 127 {
			return "", &EncodingError{
				Operation: "ASCII decoding",
				Position:  i,
				Message:   fmt.Sprintf("invalid ASCII byte: 0x%02X", b),
			}
		}
	}
	return string(bytes), nil
}

func (e *asciiEncoding) GetChars(bytes []byte) ([]rune, error) {
	str, err := e.GetString(bytes)
	if err != nil {
		return nil, err
	}
	return []rune(str), nil
}

func (e *asciiEncoding) GetByteCount(s string) int {
	return len([]rune(s))
}

func (e *asciiEncoding) GetCharCount(bytes []byte) int {
	return len(bytes)
}

func (e *asciiEncoding) EncodingName() string {
	return "US-ASCII"
}

func (e *asciiEncoding) CodePage() int {
	return 20127
}

func (e *asciiEncoding) IsSingleByte() bool {
	return true
}

// Unicode (UTF-16) Encoding Implementation
type unicodeEncoding struct {
	littleEndian bool
}

func (e *unicodeEncoding) GetBytes(s string) []byte {
	return e.GetBytesFromRunes([]rune(s))
}

func (e *unicodeEncoding) GetBytesFromRunes(chars []rune) []byte {
	encoded := utf16.Encode(chars)
	result := make([]byte, len(encoded)*2)
	
	for i, code := range encoded {
		if e.littleEndian {
			binary.LittleEndian.PutUint16(result[i*2:], code)
		} else {
			binary.BigEndian.PutUint16(result[i*2:], code)
		}
	}
	
	return result
}

func (e *unicodeEncoding) GetString(bytes []byte) (string, error) {
	if len(bytes)%2 != 0 {
		return "", &EncodingError{
			Operation: "UTF-16 decoding",
			Position:  len(bytes),
			Message:   "incomplete UTF-16 sequence (odd number of bytes)",
		}
	}
	
	codes := make([]uint16, len(bytes)/2)
	for i := 0; i < len(codes); i++ {
		if e.littleEndian {
			codes[i] = binary.LittleEndian.Uint16(bytes[i*2:])
		} else {
			codes[i] = binary.BigEndian.Uint16(bytes[i*2:])
		}
	}
	
	runes := utf16.Decode(codes)
	return string(runes), nil
}

func (e *unicodeEncoding) GetChars(bytes []byte) ([]rune, error) {
	str, err := e.GetString(bytes)
	if err != nil {
		return nil, err
	}
	return []rune(str), nil
}

func (e *unicodeEncoding) GetByteCount(s string) int {
	return len(utf16.Encode([]rune(s))) * 2
}

func (e *unicodeEncoding) GetCharCount(bytes []byte) int {
	if len(bytes)%2 != 0 {
		return 0
	}
	
	codes := make([]uint16, len(bytes)/2)
	for i := 0; i < len(codes); i++ {
		if e.littleEndian {
			codes[i] = binary.LittleEndian.Uint16(bytes[i*2:])
		} else {
			codes[i] = binary.BigEndian.Uint16(bytes[i*2:])
		}
	}
	
	return len(utf16.Decode(codes))
}

func (e *unicodeEncoding) EncodingName() string {
	if e.littleEndian {
		return "UTF-16LE"
	}
	return "UTF-16BE"
}

func (e *unicodeEncoding) CodePage() int {
	if e.littleEndian {
		return 1200
	}
	return 1201
}

func (e *unicodeEncoding) IsSingleByte() bool {
	return false
}

// UTF-32 Encoding Implementation
type utf32Encoding struct {
	littleEndian bool
}

func (e *utf32Encoding) GetBytes(s string) []byte {
	return e.GetBytesFromRunes([]rune(s))
}

func (e *utf32Encoding) GetBytesFromRunes(chars []rune) []byte {
	result := make([]byte, len(chars)*4)
	
	for i, r := range chars {
		if e.littleEndian {
			binary.LittleEndian.PutUint32(result[i*4:], uint32(r))
		} else {
			binary.BigEndian.PutUint32(result[i*4:], uint32(r))
		}
	}
	
	return result
}

func (e *utf32Encoding) GetString(bytes []byte) (string, error) {
	if len(bytes)%4 != 0 {
		return "", &EncodingError{
			Operation: "UTF-32 decoding",
			Position:  len(bytes),
			Message:   "incomplete UTF-32 sequence (byte count not multiple of 4)",
		}
	}
	
	runes := make([]rune, len(bytes)/4)
	for i := 0; i < len(runes); i++ {
		var code uint32
		if e.littleEndian {
			code = binary.LittleEndian.Uint32(bytes[i*4:])
		} else {
			code = binary.BigEndian.Uint32(bytes[i*4:])
		}
		
		if code > 0x10FFFF {
			return "", &EncodingError{
				Operation: "UTF-32 decoding",
				Position:  i * 4,
				Message:   fmt.Sprintf("invalid Unicode code point: U+%08X", code),
			}
		}
		
		runes[i] = rune(code)
	}
	
	return string(runes), nil
}

func (e *utf32Encoding) GetChars(bytes []byte) ([]rune, error) {
	str, err := e.GetString(bytes)
	if err != nil {
		return nil, err
	}
	return []rune(str), nil
}

func (e *utf32Encoding) GetByteCount(s string) int {
	return len([]rune(s)) * 4
}

func (e *utf32Encoding) GetCharCount(bytes []byte) int {
	return len(bytes) / 4
}

func (e *utf32Encoding) EncodingName() string {
	if e.littleEndian {
		return "UTF-32LE"
	}
	return "UTF-32BE"
}

func (e *utf32Encoding) CodePage() int {
	if e.littleEndian {
		return 12000
	}
	return 12001
}

func (e *utf32Encoding) IsSingleByte() bool {
	return false
}