package stringUtils

import (
	"testing"
)

func TestSplitWithOptions(t *testing.T) {
	// 测试用例1: 正常分割，不移除空条目
	s := "apple,banana,cherry"
	result := SplitWithOptions(s, ",", false)
	expected := []string{"apple", "banana", "cherry"}
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试用例2: 正常分割，移除空条目
	result = SplitWithOptions(s, ",", true)
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试用例3: 包含连续分隔符，不移除空条目
	s = "apple,,banana,,cherry"
	result = SplitWithOptions(s, ",", false)
	expected = []string{"apple", "", "banana", "", "cherry"}
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试用例4: 包含连续分隔符，移除空条目
	result = SplitWithOptions(s, ",", true)
	expected = []string{"apple", "banana", "cherry"}
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试用例5: 空字符串
	s = ""
	result = SplitWithOptions(s, ",", false)
	expected = []string{""}
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试用例6: 空字符串，移除空条目
	result = SplitWithOptions(s, ",", true)
	expected = []string{}
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试用例7: 仅包含分隔符
	s = ",,,"
	result = SplitWithOptions(s, ",", false)
	expected = []string{"", "", "", ""}
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试用例8: 仅包含分隔符，移除空条目
	result = SplitWithOptions(s, ",", true)
	expected = []string{}
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试用例9: 空分隔符
	s = "applebanana"
	result = SplitWithOptions(s, "", false)
	expected = []string{"applebanana"}
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试用例10: 空分隔符，移除空条目
	result = SplitWithOptions(s, "", true)
	expected = []string{"applebanana"}
	if !equalStringSlices(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// equalStringSlices 比较两个字符串切片是否相等
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
