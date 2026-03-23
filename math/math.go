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
// 参数:
//   value: 要计算绝对值的数字
// 返回值:
//   float64: 绝对值
func Abs(value float64) float64 {
	return math.Abs(value)
}

// AbsInt32 returns the absolute value of a 32-bit signed integer
// 参数:
//   value: 要计算绝对值的32位整数
// 返回值:
//   int32: 绝对值
func AbsInt32(value int32) int32 {
	if value < 0 {
		return -value
	}
	return value
}

// AbsInt64 returns the absolute value of a 64-bit signed integer
// 参数:
//   value: 要计算绝对值的64位整数
// 返回值:
//   int64: 绝对值
func AbsInt64(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}

// Max returns the larger of two numbers
// 参数:
//   val1: 第一个数字
//   val2: 第二个数字
// 返回值:
//   float64: 较大的数字
func Max(val1, val2 float64) float64 {
	return math.Max(val1, val2)
}

// MaxInt32 returns the larger of two 32-bit signed integers
// 参数:
//   val1: 第一个32位整数
//   val2: 第二个32位整数
// 返回值:
//   int32: 较大的整数
func MaxInt32(val1, val2 int32) int32 {
	if val1 > val2 {
		return val1
	}
	return val2
}

// MaxInt64 returns the larger of two 64-bit signed integers
// 参数:
//   val1: 第一个64位整数
//   val2: 第二个64位整数
// 返回值:
//   int64: 较大的整数
func MaxInt64(val1, val2 int64) int64 {
	if val1 > val2 {
		return val1
	}
	return val2
}

// Min returns the smaller of two numbers
// 参数:
//   val1: 第一个数字
//   val2: 第二个数字
// 返回值:
//   float64: 较小的数字
func Min(val1, val2 float64) float64 {
	return math.Min(val1, val2)
}

// MinInt32 returns the smaller of two 32-bit signed integers
// 参数:
//   val1: 第一个32位整数
//   val2: 第二个32位整数
// 返回值:
//   int32: 较小的整数
func MinInt32(val1, val2 int32) int32 {
	if val1 < val2 {
		return val1
	}
	return val2
}

// MinInt64 returns the smaller of two 64-bit signed integers
// 参数:
//   val1: 第一个64位整数
//   val2: 第二个64位整数
// 返回值:
//   int64: 较小的整数
func MinInt64(val1, val2 int64) int64 {
	if val1 < val2 {
		return val1
	}
	return val2
}

// Round rounds a value to the nearest integer or to the specified number of fractional digits
// 参数:
//   value: 要舍入的数字
// 返回值:
//   float64: 舍入后的数字
func Round(value float64) float64 {
	return math.Round(value)
}

// RoundToDigits rounds a value to the specified number of fractional digits
// 参数:
//   value: 要舍入的数字
//   digits: 小数位数
// 返回值:
//   float64: 舍入后的数字
func RoundToDigits(value float64, digits int) float64 {
	multiplier := math.Pow(10, float64(digits))
	return math.Round(value*multiplier) / multiplier
}

// Ceiling returns the smallest integral value that is greater than or equal to the specified number
// 参数:
//   value: 要向上取整的数字
// 返回值:
//   float64: 向上取整后的数字
func Ceiling(value float64) float64 {
	return math.Ceil(value)
}

// Floor returns the largest integral value less than or equal to the specified number
// 参数:
//   value: 要向下取整的数字
// 返回值:
//   float64: 向下取整后的数字
func Floor(value float64) float64 {
	return math.Floor(value)
}

// Truncate calculates the integral part of a specified number
// 参数:
//   value: 要截断的数字
// 返回值:
//   float64: 截断后的数字
func Truncate(value float64) float64 {
	return math.Trunc(value)
}

// Pow returns a specified number raised to the specified power
// 参数:
//   x: 底数
//   y: 指数
// 返回值:
//   float64: x的y次幂
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Sqrt returns the square root of a specified number
// 参数:
//   value: 要计算平方根的数字
// 返回值:
//   float64: 平方根
func Sqrt(value float64) float64 {
	return math.Sqrt(value)
}

// Exp returns e raised to the specified power
// 参数:
//   value: 指数
// 返回值:
//   float64: e的value次幂
func Exp(value float64) float64 {
	return math.Exp(value)
}

// Log returns the natural (base e) logarithm of a specified number
// 参数:
//   value: 要计算对数的数字
// 返回值:
//   float64: 自然对数
func Log(value float64) float64 {
	return math.Log(value)
}

// LogWithBase returns the logarithm of a specified number in a specified base
// 参数:
//   value: 要计算对数的数字
//   base: 对数的底数
// 返回值:
//   float64: 指定底数的对数
func LogWithBase(value, base float64) float64 {
	return math.Log(value) / math.Log(base)
}

// Log10 returns the base 10 logarithm of a specified number
// 参数:
//   value: 要计算对数的数字
// 返回值:
//   float64: 以10为底的对数
func Log10(value float64) float64 {
	return math.Log10(value)
}

// Sin returns the sine of the specified angle
// 参数:
//   value: 角度（弧度）
// 返回值:
//   float64: 正弦值
func Sin(value float64) float64 {
	return math.Sin(value)
}

// Cos returns the cosine of the specified angle
// 参数:
//   value: 角度（弧度）
// 返回值:
//   float64: 余弦值
func Cos(value float64) float64 {
	return math.Cos(value)
}

// Tan returns the tangent of the specified angle
// 参数:
//   value: 角度（弧度）
// 返回值:
//   float64: 正切值
func Tan(value float64) float64 {
	return math.Tan(value)
}

// Asin returns the angle whose sine is the specified number
// 参数:
//   value: 正弦值
// 返回值:
//   float64: 角度（弧度）
func Asin(value float64) float64 {
	return math.Asin(value)
}

// Acos returns the angle whose cosine is the specified number
// 参数:
//   value: 余弦值
// 返回值:
//   float64: 角度（弧度）
func Acos(value float64) float64 {
	return math.Acos(value)
}

// Atan returns the angle whose tangent is the specified number
// 参数:
//   value: 正切值
// 返回值:
//   float64: 角度（弧度）
func Atan(value float64) float64 {
	return math.Atan(value)
}

// Atan2 returns the angle whose tangent is the quotient of two specified numbers
// 参数:
//   y: 点的y坐标
//   x: 点的x坐标
// 返回值:
//   float64: 角度（弧度）
func Atan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

// Sign returns an integer that indicates the sign of a number
// 参数:
//   value: 要判断符号的数字
// 返回值:
//   int: 如果value>0返回1，如果value<0返回-1，否则返回0
func Sign(value float64) int {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	}
	return 0
}

// SignInt32 returns an integer that indicates the sign of a 32-bit signed integer
// 参数:
//   value: 要判断符号的32位整数
// 返回值:
//   int: 如果value>0返回1，如果value<0返回-1，否则返回0
func SignInt32(value int32) int {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	}
	return 0
}

// SignInt64 returns an integer that indicates the sign of a 64-bit signed integer
// 参数:
//   value: 要判断符号的64位整数
// 返回值:
//   int: 如果value>0返回1，如果value<0返回-1，否则返回0
func SignInt64(value int64) int {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	}
	return 0
}

// IEEERemainder returns the remainder resulting from the division of a specified number by another specified number
// 参数:
//   x: 被除数
//   y: 除数
// 返回值:
//   float64: 余数
func IEEERemainder(x, y float64) float64 {
	return math.Remainder(x, y)
}

// IsNaN returns a value indicating whether the specified number evaluates to not a number (NaN)
// 参数:
//   value: 要判断的数字
// 返回值:
//   bool: 如果是NaN返回true，否则返回false
func IsNaN(value float64) bool {
	return math.IsNaN(value)
}

// IsInfinity returns a value indicating whether the specified number evaluates to negative or positive infinity
// 参数:
//   value: 要判断的数字
// 返回值:
//   bool: 如果是无穷大返回true，否则返回false
func IsInfinity(value float64) bool {
	return math.IsInf(value, 0)
}

// IsPositiveInfinity returns a value indicating whether the specified number evaluates to positive infinity
// 参数:
//   value: 要判断的数字
// 返回值:
//   bool: 如果是正无穷大返回true，否则返回false
func IsPositiveInfinity(value float64) bool {
	return math.IsInf(value, 1)
}

// IsNegativeInfinity returns a value indicating whether the specified number evaluates to negative infinity
// 参数:
//   value: 要判断的数字
// 返回值:
//   bool: 如果是负无穷大返回true，否则返回false
func IsNegativeInfinity(value float64) bool {
	return math.IsInf(value, -1)
}
