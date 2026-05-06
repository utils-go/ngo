package converter

import (
	"math"
	"testing"
	"time"
)

func TestToInt32(t *testing.T) {
	result, err := ToInt32(42)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
	
	result, err = ToInt32("42")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
	
	// 测试指针
	val := 42
	result, err = ToInt32(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
	
	// 测试 nil 指针
	var nilPtr *int
	result, err = ToInt32(nilPtr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
	
	// 测试 nil
	result, err = ToInt32(nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestToInt64(t *testing.T) {
	result, err := ToInt64(int64(1234567890123))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1234567890123 {
		t.Errorf("Expected 1234567890123, got %d", result)
	}
	
	// 测试指针
	val := int64(1234567890123)
	result, err = ToInt64(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1234567890123 {
		t.Errorf("Expected 1234567890123, got %d", result)
	}
}

func TestToDouble(t *testing.T) {
	result, err := ToDouble(3.14159)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 3.14159 {
		t.Errorf("Expected 3.14159, got %f", result)
	}
	
	// 测试指针
	val := 3.14159
	result, err = ToDouble(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 3.14159 {
		t.Errorf("Expected 3.14159, got %f", result)
	}
}

func TestToSingle(t *testing.T) {
	result, err := ToSingle(float32(3.14))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 3.14 {
		t.Errorf("Expected 3.14, got %f", result)
	}
	
	// 测试指针
	val := float32(3.14)
	result, err = ToSingle(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 3.14 {
		t.Errorf("Expected 3.14, got %f", result)
	}
}

func TestToByte(t *testing.T) {
	result, err := ToByte(byte(255))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 255 {
		t.Errorf("Expected 255, got %d", result)
	}
	
	// 测试指针
	val := byte(128)
	result, err = ToByte(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 128 {
		t.Errorf("Expected 128, got %d", result)
	}
}

func TestToInt16(t *testing.T) {
	result, err := ToInt16(int16(32767))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 32767 {
		t.Errorf("Expected 32767, got %d", result)
	}
	
	// 测试指针
	val := int16(1000)
	result, err = ToInt16(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1000 {
		t.Errorf("Expected 1000, got %d", result)
	}
}

func TestToUInt16(t *testing.T) {
	result, err := ToUInt16(uint16(65535))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 65535 {
		t.Errorf("Expected 65535, got %d", result)
	}
	
	// 测试指针
	val := uint16(5000)
	result, err = ToUInt16(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5000 {
		t.Errorf("Expected 5000, got %d", result)
	}
}

func TestToUInt32(t *testing.T) {
	result, err := ToUInt32(uint32(4294967295))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 4294967295 {
		t.Errorf("Expected 4294967295, got %d", result)
	}
	
	// 测试指针
	val := uint32(1000000)
	result, err = ToUInt32(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1000000 {
		t.Errorf("Expected 1000000, got %d", result)
	}
}

func TestToUInt64(t *testing.T) {
	result, err := ToUInt64(uint64(18446744073709551615))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 18446744073709551615 {
		t.Errorf("Expected 18446744073709551615, got %d", result)
	}
	
	// 测试指针
	val := uint64(1000000000)
	result, err = ToUInt64(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1000000000 {
		t.Errorf("Expected 1000000000, got %d", result)
	}
}

func TestToDecimal(t *testing.T) {
	result, err := ToDecimal(123.456)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 123.456 {
		t.Errorf("Expected 123.456, got %f", result)
	}
	
	// 测试指针
	val := 123.456
	result, err = ToDecimal(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 123.456 {
		t.Errorf("Expected 123.456, got %f", result)
	}
}

func TestToChar(t *testing.T) {
	result, err := ToChar('A')
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 'A' {
		t.Errorf("Expected 'A', got %c", result)
	}
	
	// 测试指针
	val := 'B'
	result, err = ToChar(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 'B' {
		t.Errorf("Expected 'B', got %c", result)
	}
	
	// 测试字符串
	result, err = ToChar("Hello")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 'H' {
		t.Errorf("Expected 'H', got %c", result)
	}
}

func TestToDateTime(t *testing.T) {
	now := time.Now()
	result, err := ToDateTime(now)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !result.Equal(now) {
		t.Errorf("Expected %v, got %v", now, result)
	}
	
	// 测试指针
	result, err = ToDateTime(&now)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !result.Equal(now) {
		t.Errorf("Expected %v, got %v", now, result)
	}
	
	// 测试字符串
	strTime := "2023-12-25 15:30:45"
	result, err = ToDateTime(strTime)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected, _ := time.Parse("2006-01-02 15:04:05", strTime)
	if !result.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestToBoolean(t *testing.T) {
	result, err := ToBoolean(true)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !result {
		t.Error("Expected true")
	}
	
	result, err = ToBoolean("true")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !result {
		t.Error("Expected true")
	}
	
	// 测试指针
	val := true
	result, err = ToBoolean(&val)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !result {
		t.Error("Expected true")
	}
}

func TestToString(t *testing.T) {
	result := ToString(42)
	if result != "42" {
		t.Errorf("Expected '42', got '%s'", result)
	}
	
	result = ToString(true)
	if result != "True" {
		t.Errorf("Expected 'True', got '%s'", result)
	}
	
	// 测试指针
	val := 42
	result = ToString(&val)
	if result != "42" {
		t.Errorf("Expected '42', got '%s'", result)
	}
	
	// 测试 nil 指针
	var nilPtr *int
	result = ToString(nilPtr)
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
	
	// 测试 nil
	result = ToString(nil)
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestToBase64String(t *testing.T) {
	data := []byte("Hello")
	result := ToBase64String(data)
	expected := "SGVsbG8="
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFromBase64String(t *testing.T) {
	base64Str := "SGVsbG8="
	result, err := FromBase64String(base64Str)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := "Hello"
	if string(result) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, string(result))
	}
}

func TestToHexString(t *testing.T) {
	data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F}
	result := ToHexString(data)
	expected := "48656C6C6F"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestToHexStringLower(t *testing.T) {
	data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F}
	result := ToHexStringLower(data)
	expected := "48656c6c6f"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFromHexString(t *testing.T) {
	hexStr := "48656C6C6F"
	result, err := FromHexString(hexStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F}
	if string(result) != string(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestErrorCases(t *testing.T) {
	// 测试溢出
	_, err := ToInt32(int64(math.MaxInt64))
	if err == nil {
		t.Error("Expected overflow error")
	}
	
	// 测试无效字符串
	_, err = ToInt32("not a number")
	if err == nil {
		t.Error("Expected parse error")
	}
	
	// 测试无效类型
	_, err = ToInt32([]int{1, 2, 3})
	if err == nil {
		t.Error("Expected type error")
	}
}