package math

import (
	"math"
)

// Mathematical constants equivalent to System.Math constants in .NET
const (
	// E represents the natural logarithmic base, specified by the constant, e
	E = math.E
	
	// PI represents the ratio of the circumference of a circle to its diameter
	PI = math.Pi
)

// Abs returns the absolute value of a number
// Equivalent to System.Math.Abs in .NET
func Abs(value float64) float64 {
	return math.Abs(value)
}

// AbsInt32 returns the absolute value of a 32-bit signed integer
func AbsInt32(value int32) int32 {
	if value < 0 {
		return -value
	}
	return value
}

// AbsInt64 returns the absolute value of a 64-bit signed integer
func AbsInt64(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}

// Max returns the larger of two numbers
func Max(val1, val2 float64) float64 {
	return math.Max(val1, val2)
}

// MaxInt32 returns the larger of two 32-bit signed integers
func MaxInt32(val1, val2 int32) int32 {
	if val1 > val2 {
		return val1
	}
	return val2
}

// MaxInt64 returns the larger of two 64-bit signed integers
func MaxInt64(val1, val2 int64) int64 {
	if val1 > val2 {
		return val1
	}
	return val2
}

// Min returns the smaller of two numbers
func Min(val1, val2 float64) float64 {
	return math.Min(val1, val2)
}

// MinInt32 returns the smaller of two 32-bit signed integers
func MinInt32(val1, val2 int32) int32 {
	if val1 < val2 {
		return val1
	}
	return val2
}

// MinInt64 returns the smaller of two 64-bit signed integers
func MinInt64(val1, val2 int64) int64 {
	if val1 < val2 {
		return val1
	}
	return val2
}

// Round rounds a value to the nearest integer or to the specified number of fractional digits
func Round(value float64) float64 {
	return math.Round(value)
}

// RoundToDigits rounds a value to the specified number of fractional digits
func RoundToDigits(value float64, digits int) float64 {
	multiplier := math.Pow(10, float64(digits))
	return math.Round(value*multiplier) / multiplier
}

// Ceiling returns the smallest integral value that is greater than or equal to the specified number
func Ceiling(value float64) float64 {
	return math.Ceil(value)
}

// Floor returns the largest integral value less than or equal to the specified number
func Floor(value float64) float64 {
	return math.Floor(value)
}

// Truncate calculates the integral part of a specified number
func Truncate(value float64) float64 {
	return math.Trunc(value)
}

// Pow returns a specified number raised to the specified power
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Sqrt returns the square root of a specified number
func Sqrt(value float64) float64 {
	return math.Sqrt(value)
}

// Exp returns e raised to the specified power
func Exp(value float64) float64 {
	return math.Exp(value)
}

// Log returns the natural (base e) logarithm of a specified number
func Log(value float64) float64 {
	return math.Log(value)
}

// LogWithBase returns the logarithm of a specified number in a specified base
func LogWithBase(value, base float64) float64 {
	return math.Log(value) / math.Log(base)
}

// Log10 returns the base 10 logarithm of a specified number
func Log10(value float64) float64 {
	return math.Log10(value)
}

// Sin returns the sine of the specified angle
func Sin(value float64) float64 {
	return math.Sin(value)
}

// Cos returns the cosine of the specified angle
func Cos(value float64) float64 {
	return math.Cos(value)
}

// Tan returns the tangent of the specified angle
func Tan(value float64) float64 {
	return math.Tan(value)
}

// Asin returns the angle whose sine is the specified number
func Asin(value float64) float64 {
	return math.Asin(value)
}

// Acos returns the angle whose cosine is the specified number
func Acos(value float64) float64 {
	return math.Acos(value)
}

// Atan returns the angle whose tangent is the specified number
func Atan(value float64) float64 {
	return math.Atan(value)
}

// Atan2 returns the angle whose tangent is the quotient of two specified numbers
func Atan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

// Sign returns an integer that indicates the sign of a number
func Sign(value float64) int {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	}
	return 0
}

// SignInt32 returns an integer that indicates the sign of a 32-bit signed integer
func SignInt32(value int32) int {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	}
	return 0
}

// SignInt64 returns an integer that indicates the sign of a 64-bit signed integer
func SignInt64(value int64) int {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	}
	return 0
}

// IEEERemainder returns the remainder resulting from the division of a specified number by another specified number
func IEEERemainder(x, y float64) float64 {
	return math.Remainder(x, y)
}

// IsNaN returns a value indicating whether the specified number evaluates to not a number (NaN)
func IsNaN(value float64) bool {
	return math.IsNaN(value)
}

// IsInfinity returns a value indicating whether the specified number evaluates to negative or positive infinity
func IsInfinity(value float64) bool {
	return math.IsInf(value, 0)
}

// IsPositiveInfinity returns a value indicating whether the specified number evaluates to positive infinity
func IsPositiveInfinity(value float64) bool {
	return math.IsInf(value, 1)
}

// IsNegativeInfinity returns a value indicating whether the specified number evaluates to negative infinity
func IsNegativeInfinity(value float64) bool {
	return math.IsInf(value, -1)
}