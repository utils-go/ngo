package converter

import (
	"testing"
	"time"
)

func BenchmarkToInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToInt32(42)
	}
}

func BenchmarkToInt32Pointer(b *testing.B) {
	val := 42
	for i := 0; i < b.N; i++ {
		ToInt32(&val)
	}
}

func BenchmarkToInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToInt64(int64(1234567890123))
	}
}

func BenchmarkToDouble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToDouble(3.14159)
	}
}

func BenchmarkToSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToSingle(float32(3.14))
	}
}

func BenchmarkToByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToByte(byte(255))
	}
}

func BenchmarkToBoolean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToBoolean(true)
	}
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToString(42)
	}
}

func BenchmarkToBase64String(b *testing.B) {
	data := []byte("Hello, World!")
	for i := 0; i < b.N; i++ {
		ToBase64String(data)
	}
}

func BenchmarkFromBase64String(b *testing.B) {
	base64Str := "SGVsbG8sIFdvcmxkIQ=="
	for i := 0; i < b.N; i++ {
		FromBase64String(base64Str)
	}
}

func BenchmarkToHexString(b *testing.B) {
	data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F}
	for i := 0; i < b.N; i++ {
		ToHexString(data)
	}
}

func BenchmarkToHexStringLower(b *testing.B) {
	data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F}
	for i := 0; i < b.N; i++ {
		ToHexStringLower(data)
	}
}

func BenchmarkFromHexString(b *testing.B) {
	hexStr := "48656C6C6F"
	for i := 0; i < b.N; i++ {
		FromHexString(hexStr)
	}
}

func BenchmarkToDateTime(b *testing.B) {
	now := time.Now()
	for i := 0; i < b.N; i++ {
		ToDateTime(now)
	}
}

func BenchmarkToDateTimeString(b *testing.B) {
	strTime := "2023-12-25 15:30:45"
	for i := 0; i < b.N; i++ {
		ToDateTime(strTime)
	}
}