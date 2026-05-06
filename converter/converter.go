package converter

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"reflect"
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
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	if !v.IsValid() {
		return 0, nil
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := v.Int()
		if val < math.MinInt32 || val > math.MaxInt32 {
			return 0, fmt.Errorf("value %d out of range for int32", val)
		}
		return int32(val), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := v.Uint()
		if val > math.MaxInt32 {
			return 0, fmt.Errorf("value %d out of range for int32", val)
		}
		return int32(val), nil
	case reflect.Float32, reflect.Float64:
		val := v.Float()
		if val < math.MinInt32 || val > math.MaxInt32 {
			return 0, fmt.Errorf("value %f out of range for int32", val)
		}
		return int32(val), nil
	case reflect.String:
		s := strings.TrimSpace(v.String())
		if s == "" {
			return 0, nil
		}
		result, err := strconv.ParseInt(s, 10, 32)
		return int32(result), err
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int32", value)
	}
}

// ToInt64 将指定的值转换为64位有符号整数
func ToInt64(value interface{}) (int64, error) {
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := v.Uint()
		if val > math.MaxInt64 {
			return 0, fmt.Errorf("value %d out of range for int64", val)
		}
		return int64(val), nil
	case reflect.Float32, reflect.Float64:
		val := v.Float()
		if val < math.MinInt64 || val > math.MaxInt64 {
			return 0, fmt.Errorf("value %f out of range for int64", val)
		}
		return int64(val), nil
	case reflect.String:
		s := parseString(v.String())
		if isEmptyString(s) {
			return 0, nil
		}
		return strconv.ParseInt(s, 10, 64)
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int64", value)
	}
}

// ToDouble 将指定的值转换为双精度浮点数
func ToDouble(value interface{}) (float64, error) {
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	case reflect.String:
		s := parseString(v.String())
		if isEmptyString(s) {
			return 0, nil
		}
		return strconv.ParseFloat(s, 64)
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to double", value)
	}
}

// ToSingle 将指定的值转换为单精度浮点数
func ToSingle(value interface{}) (float32, error) {
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float32(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float32(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		val := v.Float()
		if val < -math.MaxFloat32 || val > math.MaxFloat32 {
			return 0, fmt.Errorf("value %f out of range for float32", val)
		}
		return float32(val), nil
	case reflect.String:
		val, err := strconv.ParseFloat(strings.TrimSpace(v.String()), 32)
		return float32(val), err
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to single", value)
	}
}

// ToByte 将指定的值转换为8位无符号整数
func ToByte(value interface{}) (byte, error) {
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := v.Int()
		if val < 0 || val > math.MaxUint8 {
			return 0, fmt.Errorf("value %d out of range for byte", val)
		}
		return byte(val), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := v.Uint()
		if val > math.MaxUint8 {
			return 0, fmt.Errorf("value %d out of range for byte", val)
		}
		return byte(val), nil
	case reflect.Float32, reflect.Float64:
		val := v.Float()
		if val < 0 || val > math.MaxUint8 {
			return 0, fmt.Errorf("value %f out of range for byte", val)
		}
		return byte(val), nil
	case reflect.String:
		val, err := strconv.ParseUint(strings.TrimSpace(v.String()), 10, 8)
		return byte(val), err
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to byte", value)
	}
}

// ToInt16 将指定的值转换为16位有符号整数
func ToInt16(value interface{}) (int16, error) {
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := v.Int()
		if val < math.MinInt16 || val > math.MaxInt16 {
			return 0, fmt.Errorf("value %d out of range for int16", val)
		}
		return int16(val), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := v.Uint()
		if val > math.MaxInt16 {
			return 0, fmt.Errorf("value %d out of range for int16", val)
		}
		return int16(val), nil
	case reflect.Float32, reflect.Float64:
		val := v.Float()
		if val < math.MinInt16 || val > math.MaxInt16 {
			return 0, fmt.Errorf("value %f out of range for int16", val)
		}
		return int16(val), nil
	case reflect.String:
		val, err := strconv.ParseInt(strings.TrimSpace(v.String()), 10, 16)
		return int16(val), err
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int16", value)
	}
}

// ToUInt16 将指定的值转换为16位无符号整数
func ToUInt16(value interface{}) (uint16, error) {
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := v.Int()
		if val < 0 || val > math.MaxUint16 {
			return 0, fmt.Errorf("value %d out of range for uint16", val)
		}
		return uint16(val), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := v.Uint()
		if val > math.MaxUint16 {
			return 0, fmt.Errorf("value %d out of range for uint16", val)
		}
		return uint16(val), nil
	case reflect.Float32, reflect.Float64:
		val := v.Float()
		if val < 0 || val > math.MaxUint16 {
			return 0, fmt.Errorf("value %f out of range for uint16", val)
		}
		return uint16(val), nil
	case reflect.String:
		val, err := strconv.ParseUint(strings.TrimSpace(v.String()), 10, 16)
		return uint16(val), err
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to uint16", value)
	}
}

// ToUInt32 将指定的值转换为32位无符号整数
func ToUInt32(value interface{}) (uint32, error) {
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := v.Int()
		if val < 0 || val > math.MaxUint32 {
			return 0, fmt.Errorf("value %d out of range for uint32", val)
		}
		return uint32(val), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := v.Uint()
		if val > math.MaxUint32 {
			return 0, fmt.Errorf("value %d out of range for uint32", val)
		}
		return uint32(val), nil
	case reflect.Float32, reflect.Float64:
		val := v.Float()
		if val < 0 || val > math.MaxUint32 {
			return 0, fmt.Errorf("value %f out of range for uint32", val)
		}
		return uint32(val), nil
	case reflect.String:
		val, err := strconv.ParseUint(strings.TrimSpace(v.String()), 10, 32)
		return uint32(val), err
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to uint32", value)
	}
}

// ToUInt64 将指定的值转换为64位无符号整数
func ToUInt64(value interface{}) (uint64, error) {
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := v.Int()
		if val < 0 {
			return 0, fmt.Errorf("value %d out of range for uint64", val)
		}
		return uint64(val), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint(), nil
	case reflect.Float32, reflect.Float64:
		val := v.Float()
		if val < 0 || val > math.MaxUint64 {
			return 0, fmt.Errorf("value %f out of range for uint64", val)
		}
		return uint64(val), nil
	case reflect.String:
		return strconv.ParseUint(strings.TrimSpace(v.String()), 10, 64)
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to uint64", value)
	}
}

// ToDecimal 将指定的值转换为十进制数（Go 中使用 float64 表示）
func ToDecimal(value interface{}) (float64, error) {
	return ToDouble(value)
}

// ToChar 将指定的值转换为 Unicode 字符
func ToChar(value interface{}) (rune, error) {
	v, err := derefValue(value)
	if err != nil {
		return 0, err
	}
	
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := v.Int()
		if val < 0 || val > 0x10FFFF {
			return 0, fmt.Errorf("value %d out of range for rune", val)
		}
		return rune(val), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := v.Uint()
		if val > 0x10FFFF {
			return 0, fmt.Errorf("value %d out of range for rune", val)
		}
		return rune(val), nil
	case reflect.String:
		s := strings.TrimSpace(v.String())
		if len(s) == 0 {
			return 0, nil
		}
		return []rune(s)[0], nil
	default:
		return 0, fmt.Errorf("cannot convert %T to char", value)
	}
}

// ToDateTime 将指定的值转换为时间
func ToDateTime(value interface{}) (time.Time, error) {
	v, err := derefValue(value)
	if err != nil {
		return time.Time{}, err
	}
	
	switch v.Kind() {
	case reflect.String:
		s := strings.TrimSpace(v.String())
		// 尝试常见的时间格式
		formats := []string{
			"2006-01-02 15:04:05",
			"2006-01-02",
			"2006/01/02 15:04:05",
			"2006/01/02",
			time.RFC3339,
			time.RFC3339Nano,
		}
		
		for _, format := range formats {
			if t, err := time.Parse(format, s); err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("cannot parse '%s' as datetime", s)
	case reflect.Struct:
		if t, ok := v.Interface().(time.Time); ok {
			return t, nil
		}
	}
	
	return time.Time{}, fmt.Errorf("cannot convert %T to datetime", value)
}

// ToBoolean 将指定的值转换为等效的布尔值
func ToBoolean(value interface{}) (bool, error) {
	v, err := derefValue(value)
	if err != nil {
		return false, err
	}
	
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() != 0, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() != 0, nil
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0, nil
	case reflect.String:
		s := strings.ToLower(strings.TrimSpace(v.String()))
		switch s {
		case "true", "1", "yes", "on":
			return true, nil
		case "false", "0", "no", "off", "":
			return false, nil
		default:
			return false, fmt.Errorf("cannot convert '%s' to boolean", s)
		}
	default:
		return false, fmt.Errorf("cannot convert %T to boolean", value)
	}
}

// ToString 将指定的值转换为其等效的字符串表示形式
func ToString(value interface{}) string {
	if value == nil {
		return ""
	}
	
	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64)
	case reflect.Bool:
		if v.Bool() {
			return "True"
		}
		return "False"
	case reflect.Struct:
		if t, ok := v.Interface().(time.Time); ok {
			return t.Format("2006-01-02 15:04:05")
		}
	}
	
	return fmt.Sprintf("%v", value)
}

// ToBase64String 将8位无符号整数数组转换为base-64字符串
func ToBase64String(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// FromBase64String 将base-64字符串转换为字节数组
func FromBase64String(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// ToHexString 将字节数组转换为十六进制字符串（大写）
func ToHexString(data []byte) string {
	return strings.ToUpper(hex.EncodeToString(data))
}

// ToHexStringLower 将字节数组转换为十六进制字符串（小写）
func ToHexStringLower(data []byte) string {
	return strings.ToLower(hex.EncodeToString(data))
}

// FromHexString 将十六进制字符串转换为字节数组
func FromHexString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

// derefValue 解引用指针，返回实际的值
func derefValue(value interface{}) (reflect.Value, error) {
	if value == nil {
		return reflect.Value{}, nil
	}
	
	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return reflect.Value{}, nil
		}
		v = v.Elem()
	}
	
	return v, nil
}

// checkValidValue 检查值是否有效，如果无效返回零值
func checkValidValue(v reflect.Value) bool {
	return v.IsValid()
}

// parseString 解析字符串，处理空字符串情况
func parseString(s string) string {
	return strings.TrimSpace(s)
}

// isEmptyString 检查字符串是否为空（去除空格后）
func isEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

// ConvertToIntFromString 将字符串转换为整数（向后兼容）
func ConvertToIntFromString(s string) (int, error) {
	return strconv.Atoi(s)
}

// ConvertToStringFromInt 将整数转换为字符串（向后兼容）
func ConvertToStringFromInt(i int) string {
	return strconv.Itoa(i)
}

// ConvertToBoolFromString 将字符串转换为布尔值（向后兼容）
func ConvertToBoolFromString(s string) (bool, error) {
	return ToBoolean(s)
}

// ConvertStringFromBool 将布尔值转换为字符串（向后兼容）
func ConvertStringFromBool(b bool) string {
	return ToString(b)
}