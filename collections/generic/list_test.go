package generic

import (
	"testing"
)

func TestList_NewList(t *testing.T) {
	list := NewList[int]()
	
	if list.Count() != 0 {
		t.Errorf("Expected count 0, got %d", list.Count())
	}
}

func TestList_Add(t *testing.T) {
	list := NewList[int]()
	
	list.Add(1)
	list.Add(2)
	list.Add(3)
	
	if list.Count() != 3 {
		t.Errorf("Expected count 3, got %d", list.Count())
	}
	
	item, err := list.Get(0)
	if err != nil || item != 1 {
		t.Errorf("Expected item 1 at index 0, got %d", item)
	}
	
	item, err = list.Get(2)
	if err != nil || item != 3 {
		t.Errorf("Expected item 3 at index 2, got %d", item)
	}
}

func TestList_AddRange(t *testing.T) {
	list := NewList[int]()
	items := []int{1, 2, 3, 4, 5}
	
	list.AddRange(items)
	
	if list.Count() != 5 {
		t.Errorf("Expected count 5, got %d", list.Count())
	}
	
	for i, expected := range items {
		item, err := list.Get(i)
		if err != nil || item != expected {
			t.Errorf("Expected item %d at index %d, got %d", expected, i, item)
		}
	}
}

func TestList_Insert(t *testing.T) {
	list := NewList[string]()
	list.Add("apple")
	list.Add("cherry")
	
	err := list.Insert(1, "banana")
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	
	if list.Count() != 3 {
		t.Errorf("Expected count 3, got %d", list.Count())
	}
	
	expected := []string{"apple", "banana", "cherry"}
	for i, exp := range expected {
		item, err := list.Get(i)
		if err != nil || item != exp {
			t.Errorf("Expected item '%s' at index %d, got '%s'", exp, i, item)
		}
	}
}

func TestList_Remove(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 2, 4})
	
	// Remove first occurrence of 2
	removed := list.Remove(2)
	if !removed {
		t.Error("Expected Remove(2) to return true")
	}
	
	if list.Count() != 4 {
		t.Errorf("Expected count 4, got %d", list.Count())
	}
	
	// Verify the first 2 was removed, not the second
	expected := []int{1, 3, 2, 4}
	for i, exp := range expected {
		item, err := list.Get(i)
		if err != nil || item != exp {
			t.Errorf("Expected item %d at index %d, got %d", exp, i, item)
		}
	}
}

func TestList_RemoveAt(t *testing.T) {
	list := NewList[string]()
	list.AddRange([]string{"apple", "banana", "cherry"})
	
	err := list.RemoveAt(1)
	if err != nil {
		t.Errorf("RemoveAt failed: %v", err)
	}
	
	if list.Count() != 2 {
		t.Errorf("Expected count 2, got %d", list.Count())
	}
	
	expected := []string{"apple", "cherry"}
	for i, exp := range expected {
		item, err := list.Get(i)
		if err != nil || item != exp {
			t.Errorf("Expected item '%s' at index %d, got '%s'", exp, i, item)
		}
	}
}

func TestList_Contains(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	
	if !list.Contains(3) {
		t.Error("Expected Contains(3) to be true")
	}
	
	if list.Contains(10) {
		t.Error("Expected Contains(10) to be false")
	}
}

func TestList_IndexOf(t *testing.T) {
	list := NewList[string]()
	list.AddRange([]string{"apple", "banana", "cherry", "banana"})
	
	index := list.IndexOf("banana")
	if index != 1 {
		t.Errorf("Expected IndexOf('banana') to be 1, got %d", index)
	}
	
	index = list.IndexOf("grape")
	if index != -1 {
		t.Errorf("Expected IndexOf('grape') to be -1, got %d", index)
	}
}

func TestList_LastIndexOf(t *testing.T) {
	list := NewList[string]()
	list.AddRange([]string{"apple", "banana", "cherry", "banana"})
	
	index := list.LastIndexOf("banana")
	if index != 3 {
		t.Errorf("Expected LastIndexOf('banana') to be 3, got %d", index)
	}
	
	index = list.LastIndexOf("grape")
	if index != -1 {
		t.Errorf("Expected LastIndexOf('grape') to be -1, got %d", index)
	}
}

func TestList_Find(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	
	item, found := list.Find(func(x int) bool { return x > 3 })
	if !found || item != 4 {
		t.Errorf("Expected Find(x > 3) to return (4, true), got (%d, %t)", item, found)
	}
	
	item, found = list.Find(func(x int) bool { return x > 10 })
	if found {
		t.Errorf("Expected Find(x > 10) to return (0, false), got (%d, %t)", item, found)
	}
}

func TestList_FindAll(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5, 6})
	
	evenNumbers := list.FindAll(func(x int) bool { return x%2 == 0 })
	
	if evenNumbers.Count() != 3 {
		t.Errorf("Expected 3 even numbers, got %d", evenNumbers.Count())
	}
	
	expected := []int{2, 4, 6}
	for i, exp := range expected {
		item, err := evenNumbers.Get(i)
		if err != nil || item != exp {
			t.Errorf("Expected even number %d at index %d, got %d", exp, i, item)
		}
	}
}

func TestList_Sort(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{5, 2, 8, 1, 9, 3})
	
	list.Sort(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	
	expected := []int{1, 2, 3, 5, 8, 9}
	for i, exp := range expected {
		item, err := list.Get(i)
		if err != nil || item != exp {
			t.Errorf("Expected sorted item %d at index %d, got %d", exp, i, item)
		}
	}
}

func TestList_Reverse(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	
	list.Reverse()
	
	expected := []int{5, 4, 3, 2, 1}
	for i, exp := range expected {
		item, err := list.Get(i)
		if err != nil || item != exp {
			t.Errorf("Expected reversed item %d at index %d, got %d", exp, i, item)
		}
	}
}

func TestList_ForEach(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	
	sum := 0
	list.ForEach(func(item int) {
		sum += item
	})
	
	expectedSum := 15
	if sum != expectedSum {
		t.Errorf("Expected sum %d, got %d", expectedSum, sum)
	}
}

func TestList_ToArray(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	
	array := list.ToArray()
	
	if len(array) != 5 {
		t.Errorf("Expected array length 5, got %d", len(array))
	}
	
	for i, expected := range []int{1, 2, 3, 4, 5} {
		if array[i] != expected {
			t.Errorf("Expected array[%d] = %d, got %d", i, expected, array[i])
		}
	}
}

func TestList_Clear(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	
	list.Clear()
	
	if list.Count() != 0 {
		t.Errorf("Expected count 0 after Clear(), got %d", list.Count())
	}
}

func TestList_Exists(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	
	if !list.Exists(func(x int) bool { return x > 3 }) {
		t.Error("Expected Exists(x > 3) to be true")
	}
	
	if list.Exists(func(x int) bool { return x > 10 }) {
		t.Error("Expected Exists(x > 10) to be false")
	}
}

func TestList_TrueForAll(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{2, 4, 6, 8, 10})
	
	if !list.TrueForAll(func(x int) bool { return x%2 == 0 }) {
		t.Error("Expected TrueForAll(x % 2 == 0) to be true for all even numbers")
	}
	
	list.Add(3)
	if list.TrueForAll(func(x int) bool { return x%2 == 0 }) {
		t.Error("Expected TrueForAll(x % 2 == 0) to be false after adding odd number")
	}
}