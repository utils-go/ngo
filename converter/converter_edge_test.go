package converter

import (
	"math"
	"testing"
)

func TestEdgeCases(t *testing.T) {
	// 测试边界值
	t.Run("Int32Boundaries", func(t *testing.T) {
		// 最小值
		result, err := ToInt32(math.MinInt32)
		if err != nil {
			t.Errorf("MinInt32 error: %v", err)
		}
		if result != math.MinInt32 {
			t.Errorf("Expected %d, got %d", math.MinInt32, result)
		}

		// 最大值
		result, err = ToInt32(math.MaxInt32)
		if err != nil {
			t.Errorf("MaxInt32 error: %v", err)
		}
		if result != math.MaxInt32 {
			t.Errorf("Expected %d, got %d", math.MaxInt32, result)
		}

		// 溢出测试
		_, err = ToInt32(int64(math.MinInt32 - 1))
		if err == nil {
			t.Error("Expected overflow error for MinInt32-1")
		}

		_, err = ToInt32(int64(math.MaxInt32 + 1))
		if err == nil {
			t.Error("Expected overflow error for MaxInt32+1")
		}
	})

	t.Run("UInt32Boundaries", func(t *testing.T) {
		// 最小值
		result, err := ToUInt32(0)
		if err != nil {
			t.Errorf("0 error: %v", err)
		}
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}

		// 最大值
		result, err = ToUInt32(math.MaxUint32)
		if err != nil {
			t.Errorf("MaxUint32 error: %v", err)
		}
		if result != math.MaxUint32 {
			t.Errorf("Expected %d, got %d", math.MaxUint32, result)
		}

		// 溢出测试
		_, err = ToUInt32(-1)
		if err == nil {
			t.Error("Expected overflow error for -1")
		}

		_, err = ToUInt32(uint64(math.MaxUint32 + 1))
		if err == nil {
			t.Error("Expected overflow error for MaxUint32+1")
		}
	})

	t.Run("FloatBoundaries", func(t *testing.T) {
		// 单精度浮点数边界
		result, err := ToSingle(math.MaxFloat32)
		if err != nil {
			t.Errorf("MaxFloat32 error: %v", err)
		}
		if result != math.MaxFloat32 {
			t.Errorf("Expected %f, got %f", math.MaxFloat32, result)
		}

		// 双精度浮点数
		result64, err := ToDouble(math.MaxFloat64)
		if err != nil {
			t.Errorf("MaxFloat64 error: %v", err)
		}
		if result64 != math.MaxFloat64 {
			t.Errorf("Expected %f, got %f", math.MaxFloat64, result64)
		}
	})

	t.Run("NestedPointers", func(t *testing.T) {
		// 嵌套指针测试
		val := 42
		ptr1 := &val
		ptr2 := &ptr1
		ptr3 := &ptr2

		result, err := ToInt32(ptr3)
		if err != nil {
			t.Errorf("Nested pointer error: %v", err)
		}
		if result != 42 {
			t.Errorf("Expected 42, got %d", result)
		}
	})

	t.Run("EmptyString", func(t *testing.T) {
		// 空字符串测试
		result, err := ToInt32("")
		if err != nil {
			t.Errorf("Empty string error: %v", err)
		}
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}

		resultBool, err := ToBoolean("")
		if err != nil {
			t.Errorf("Empty string to bool error: %v", err)
		}
		if resultBool {
			t.Error("Expected false for empty string")
		}
	})

	t.Run("WhitespaceString", func(t *testing.T) {
		// 空白字符串测试
		result, err := ToInt32("  42  ")
		if err != nil {
			t.Errorf("Whitespace string error: %v", err)
		}
		if result != 42 {
			t.Errorf("Expected 42, got %d", result)
		}

		resultBool, err := ToBoolean("  true  ")
		if err != nil {
			t.Errorf("Whitespace bool error: %v", err)
		}
		if !resultBool {
			t.Error("Expected true for '  true  '")
		}
	})

	t.Run("HexStringEdgeCases", func(t *testing.T) {
		// 空字节数组
		result := ToHexString([]byte{})
		if result != "" {
			t.Errorf("Expected empty string, got '%s'", result)
		}

		resultLower := ToHexStringLower([]byte{})
		if resultLower != "" {
			t.Errorf("Expected empty string, got '%s'", resultLower)
		}

		// 单字节
		result = ToHexString([]byte{0x00})
		if result != "00" {
			t.Errorf("Expected '00', got '%s'", result)
		}

		result = ToHexString([]byte{0xFF})
		if result != "FF" {
			t.Errorf("Expected 'FF', got '%s'", result)
		}
	})

	t.Run("Base64EdgeCases", func(t *testing.T) {
		// 空字节数组
		result := ToBase64String([]byte{})
		if result != "" {
			t.Errorf("Expected empty string, got '%s'", result)
		}

		// 单字节
		result = ToBase64String([]byte{0x00})
		if result != "AA==" {
			t.Errorf("Expected 'AA==', got '%s'", result)
		}

		// 无效 Base64
		_, err := FromBase64String("Invalid!")
		if err == nil {
			t.Error("Expected error for invalid base64")
		}
	})

	t.Run("CharEdgeCases", func(t *testing.T) {
		// 空字符串
		result, err := ToChar("")
		if err != nil {
			t.Errorf("Empty string to char error: %v", err)
		}
		if result != 0 {
			t.Errorf("Expected 0, got %c", result)
		}

		// 多字符字符串
		result, err = ToChar("Hello")
		if err != nil {
			t.Errorf("Multi-char string error: %v", err)
		}
		if result != 'H' {
			t.Errorf("Expected 'H', got %c", result)
		}

		// Unicode 字符
		result, err = ToChar("世界")
		if err != nil {
			t.Errorf("Unicode char error: %v", err)
		}
		if result != '世' {
			t.Errorf("Expected '世', got %c", result)
		}
	})
}

func TestTypeCompatibility(t *testing.T) {
	// 测试各种类型之间的兼容性
	testCases := []struct {
		name  string
		value interface{}
		check func(t *testing.T, result interface{}, err error)
	}{
		{
			name:  "IntToAll",
			value: 42,
			check: func(t *testing.T, result interface{}, err error) {
				if err != nil {
					t.Errorf("Int conversion error: %v", err)
				}
			},
		},
		{
			name:  "FloatToAll",
			value: 3.14,
			check: func(t *testing.T, result interface{}, err error) {
				if err != nil {
					t.Errorf("Float conversion error: %v", err)
				}
			},
		},
		{
			name:  "StringToAll",
			value: "123",
			check: func(t *testing.T, result interface{}, err error) {
				// "123" 转换为 boolean 应该失败，其他转换应该成功
				// 这个检查会在具体的测试中处理
			},
		},
		{
			name:  "BoolToAll",
			value: true,
			check: func(t *testing.T, result interface{}, err error) {
				if err != nil {
					t.Errorf("Bool conversion error: %v", err)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 测试所有转换函数
			_, err := ToInt32(tc.value)
			tc.check(t, nil, err)

			_, err = ToInt64(tc.value)
			tc.check(t, nil, err)

			_, err = ToDouble(tc.value)
			tc.check(t, nil, err)

			_, err = ToSingle(tc.value)
			tc.check(t, nil, err)

			// 对于字符串 "123"，转换为 boolean 应该失败
			if tc.name == "StringToAll" {
				_, err = ToBoolean(tc.value)
				if err == nil {
					t.Error("Expected error when converting '123' to boolean")
				}
			} else {
				_, err = ToBoolean(tc.value)
				tc.check(t, nil, err)
			}

			// ToString 应该总是成功
			_ = ToString(tc.value)
		})
	}
}