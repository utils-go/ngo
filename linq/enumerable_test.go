package linq

import (
	"testing"
)

func TestFrom(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	enum := From(items)
	
	if enum.Count() != 5 {
		t.Errorf("Expected count 5, got %d", enum.Count())
	}
	
	result := enum.ToSlice()
	for i, v := range result {
		if v != items[i] {
			t.Errorf("Expected %d at index %d, got %d", items[i], i, v)
		}
	}
}

func TestRange(t *testing.T) {
	enum := Range(5, 3) // Start at 5, count 3: [5, 6, 7]
	
	if enum.Count() != 3 {
		t.Errorf("Expected count 3, got %d", enum.Count())
	}
	
	result := enum.ToSlice()
	expected := []int{5, 6, 7}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestWhere(t *testing.T) {
	enum := From([]int{1, 2, 3, 4, 5, 6})
	
	// Filter even numbers
	evens := enum.Where(func(x int) bool { return x%2 == 0 })
	
	if evens.Count() != 3 {
		t.Errorf("Expected 3 even numbers, got %d", evens.Count())
	}
	
	result := evens.ToSlice()
	expected := []int{2, 4, 6}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestSelect(t *testing.T) {
	enum := From([]int{1, 2, 3})
	
	// Square each number
	squares := Select(enum, func(x int) int { return x * x })
	
	if squares.Count() != 3 {
		t.Errorf("Expected count 3, got %d", squares.Count())
	}
	
	result := squares.ToSlice()
	expected := []int{1, 4, 9}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestFirst(t *testing.T) {
	enum := From([]int{1, 2, 3, 4, 5})
	
	first, exists := enum.First()
	if !exists || first != 1 {
		t.Errorf("Expected first element 1, got %d (exists: %t)", first, exists)
	}
	
	// Test empty enumerable
	empty := From([]int{})
	_, exists = empty.First()
	if exists {
		t.Error("Expected First() to return false for empty enumerable")
	}
}

func TestFirstWhere(t *testing.T) {
	enum := From([]int{1, 2, 3, 4, 5})
	
	// First even number
	first, exists := enum.FirstWhere(func(x int) bool { return x%2 == 0 })
	if !exists || first != 2 {
		t.Errorf("Expected first even number 2, got %d (exists: %t)", first, exists)
	}
	
	// No number > 10
	_, exists = enum.FirstWhere(func(x int) bool { return x > 10 })
	if exists {
		t.Error("Expected FirstWhere to return false for condition not met")
	}
}

func TestAny(t *testing.T) {
	enum := From([]int{1, 2, 3, 4, 5})
	
	hasEven := enum.Any(func(x int) bool { return x%2 == 0 })
	if !hasEven {
		t.Error("Expected Any to return true for even numbers")
	}
	
	hasLarge := enum.Any(func(x int) bool { return x > 10 })
	if hasLarge {
		t.Error("Expected Any to return false for numbers > 10")
	}
}

func TestAll(t *testing.T) {
	enum := From([]int{2, 4, 6, 8})
	
	allEven := enum.All(func(x int) bool { return x%2 == 0 })
	if !allEven {
		t.Error("Expected All to return true for all even numbers")
	}
	
	enum2 := From([]int{1, 2, 3, 4})
	allEven2 := enum2.All(func(x int) bool { return x%2 == 0 })
	if allEven2 {
		t.Error("Expected All to return false when not all numbers are even")
	}
}

func TestContains(t *testing.T) {
	enum := From([]int{1, 2, 3, 4, 5})
	
	if !enum.Contains(3) {
		t.Error("Expected Contains(3) to return true")
	}
	
	if enum.Contains(10) {
		t.Error("Expected Contains(10) to return false")
	}
}

func TestDistinct(t *testing.T) {
	enum := From([]int{1, 2, 2, 3, 3, 3, 4})
	
	distinct := enum.Distinct()
	
	if distinct.Count() != 4 {
		t.Errorf("Expected 4 distinct elements, got %d", distinct.Count())
	}
	
	result := distinct.ToSlice()
	expected := []int{1, 2, 3, 4}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestSkip(t *testing.T) {
	enum := From([]int{1, 2, 3, 4, 5})
	
	skipped := enum.Skip(2)
	
	if skipped.Count() != 3 {
		t.Errorf("Expected 3 elements after skipping 2, got %d", skipped.Count())
	}
	
	result := skipped.ToSlice()
	expected := []int{3, 4, 5}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestTake(t *testing.T) {
	enum := From([]int{1, 2, 3, 4, 5})
	
	taken := enum.Take(3)
	
	if taken.Count() != 3 {
		t.Errorf("Expected 3 elements, got %d", taken.Count())
	}
	
	result := taken.ToSlice()
	expected := []int{1, 2, 3}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestReverse(t *testing.T) {
	enum := From([]int{1, 2, 3, 4, 5})
	
	reversed := enum.Reverse()
	
	result := reversed.ToSlice()
	expected := []int{5, 4, 3, 2, 1}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestOrderBy(t *testing.T) {
	enum := From([]int{5, 2, 8, 1, 9})
	
	// Order by value (ascending)
	ordered := OrderBy(enum, func(x int) int { return x }, func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	
	result := ordered.ToSlice()
	expected := []int{1, 2, 5, 8, 9}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestConcat(t *testing.T) {
	enum1 := From([]int{1, 2, 3})
	enum2 := From([]int{4, 5, 6})
	
	concatenated := enum1.Concat(enum2)
	
	if concatenated.Count() != 6 {
		t.Errorf("Expected 6 elements, got %d", concatenated.Count())
	}
	
	result := concatenated.ToSlice()
	expected := []int{1, 2, 3, 4, 5, 6}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestUnion(t *testing.T) {
	enum1 := From([]int{1, 2, 3})
	enum2 := From([]int{3, 4, 5})
	
	union := enum1.Union(enum2)
	
	if union.Count() != 5 {
		t.Errorf("Expected 5 unique elements, got %d", union.Count())
	}
	
	result := union.ToSlice()
	expected := []int{1, 2, 3, 4, 5}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestIntersect(t *testing.T) {
	enum1 := From([]int{1, 2, 3, 4})
	enum2 := From([]int{3, 4, 5, 6})
	
	intersection := enum1.Intersect(enum2)
	
	if intersection.Count() != 2 {
		t.Errorf("Expected 2 common elements, got %d", intersection.Count())
	}
	
	result := intersection.ToSlice()
	expected := []int{3, 4}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestAggregate(t *testing.T) {
	enum := From([]int{1, 2, 3, 4, 5})
	
	// Sum all numbers
	sum := Aggregate(enum, 0, func(acc, x int) int { return acc + x })
	
	expected := 15
	if sum != expected {
		t.Errorf("Expected sum %d, got %d", expected, sum)
	}
}