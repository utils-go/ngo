package converter

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ToInt32 将指定的值转换为32位有符号整数
// 参数:
//   value: 要转换的值
// 返回值:
//   int32: 转换后的32位有符号整数
//   error: 转换过程中的错误
func ToInt32(value interface{}) (int32, error) {
	switch v := value.(type) {
	case int32:
		return v, nil
	case int:
		return int32(v), nil
	case int64:
		return int32(v), nil
	case float32:
		return int32(v), nil
	case float64:
		return int32(v), nil
	case string:
		result, err := strconv.ParseInt(strings.TrimSpace(v), 10, 32)
		return int32(result), err
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int32", value)
	}
}

// ToBoolean 将指定的值转换为等效的布尔值
// 参数:
//   value: 要转换的值
// 返回值:
//   bool: 转换后的布尔值
//   error: 转换过程中的错误
func ToBoolean(value interface{}) (bool, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case int32:
		return v != 0, nil
	case int64:
		return v != 0, nil
	case int:
		return v != 0, nil
	case float32:
		return v != 0, nil
	case float64:
		return v != 0, nil
	case string:
		s := strings.ToLower(strings.TrimSpace(v))
		switch s {
		case "true", "1", "yes", "on":
			return true, nil
		case "false", "0", "no", "off", "":
			return false, nil
		default:
			return false, fmt.Errorf("cannot convert '%s' to boolean", v)
		}
	default:
		return false, fmt.Errorf("cannot convert %T to boolean", value)
	}
}

// ToString 将指定的值转换为其等效的字符串表示形式
// 参数:
//   value: 要转换的值
// 返回值:
//   string: 转换后的字符串表示形式
func ToString(value interface{}) string {
	if value == nil {
		return ""
	}
	
	switch v := value.(type) {
	case string:
		return v
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case int:
		return strconv.Itoa(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	case bool:
		if v {
			return "True"
		}
		return "False"
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	default:
		return fmt.Sprintf("%v", value)
	}
}

// ToBase64String 将8位无符号整数数组转换为base-64字符串
// 参数:
//   data: 要转换的字节数组
// 返回值:
//   string: 转换后的base-64字符串
func ToBase64String(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// FromBase64String 将base-64字符串转换为字节数组
// 参数:
//   s: 要转换的base-64字符串
// 返回值:
//   []byte: 转换后的字节数组
//   error: 转换过程中的错误
func FromBase64String(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// ConvertToIntFromString 将字符串转换为整数（向后兼容）
// 参数:
//   s: 要转换的字符串
// 返回值:
//   int: 转换后的整数
//   error: 转换过程中的错误
func ConvertToIntFromString(s string) (int, error) {
	return strconv.Atoi(s)
}

// ConvertToStringFromInt 将整数转换为字符串（向后兼容）
// 参数:
//   i: 要转换的整数
// 返回值:
//   string: 转换后的字符串
func ConvertToStringFromInt(i int) string {
	return strconv.Itoa(i)
}

// ConvertToBoolFromString 将字符串转换为布尔值（向后兼容）
// 参数:
//   s: 要转换的字符串
// 返回值:
//   bool: 转换后的布尔值
//   error: 转换过程中的错误
func ConvertToBoolFromString(s string) (bool, error) {
	return ToBoolean(s)
}

// ConvertStringFromBool 将布尔值转换为字符串（向后兼容）
// 参数:
//   b: 要转换的布尔值
// 返回值:
//   string: 转换后的字符串
func ConvertStringFromBool(b bool) string {
	return ToString(b)
}