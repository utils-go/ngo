package converter

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ToInt32 converts the specified value to a 32-bit signed integer
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

// ToBoolean converts the specified value to an equivalent Boolean value
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

// ToString converts the specified value to its equivalent string representation
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

// ToBase64String converts an array of 8-bit unsigned integers to base-64 string
func ToBase64String(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// FromBase64String converts base-64 string to byte array
func FromBase64String(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// Legacy methods for backward compatibility
func ConvertToIntFromString(s string) (int, error) {
	return strconv.Atoi(s)
}

func ConvertToStringFromInt(i int) string {
	return strconv.Itoa(i)
}

func ConvertToBoolFromString(s string) (bool, error) {
	return ToBoolean(s)
}

func ConvertStringFromBool(b bool) string {
	return ToString(b)
}