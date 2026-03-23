package utils

import (
	"strings"
)

// timeFormat 存储时间格式映射
var timeFormat map[string]string

func init() {
	timeFormat = map[string]string{
		"yyyy": "2006",
		"MM":   "01",
		"dd":   "02",
		"HH":   "15",
		"mm":   "04",
		"ss":   "05",
		"fff":  "000",
	}
}

// ConvertLayout 将C#日期时间格式转换为Go日期时间格式
// 参数:
//   cslayout: C#日期时间格式字符串
// 返回值:
//   string: Go日期时间格式字符串
func ConvertLayout(cslayout string) string {
	cslayoutNew := strings.ReplaceAll(cslayout, "h", "H")
	for k, v := range timeFormat {
		if strings.Contains(cslayoutNew, k) {
			cslayoutNew = strings.ReplaceAll(cslayout, k, v)
		}
	}
	return cslayoutNew
}
