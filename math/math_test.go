package math

import (
	"math"
	"testing"
)

func TestConstants(t *testing.T) {
	if E != math.E {
		t.Errorf("Expected E to be %f, got %f", math.E, E)
	}
	
	if PI != math.Pi {
		t.Errorf("Expected PI to be %f, got %f", math.Pi, PI)
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{-5.5, 5.5},
		{5.5, 5.5},
		{0, 0},
		{-0, 0},
	}
	
	for _, test := range tests {
		result := Abs(test.input)
		if result != test.expected {
			t.Errorf("Abs(%f) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

func TestAbsInt32(t *testing.T) {
	tests := []struct {
		input    int32
		expected int32
	}{
		{-5, 5},
		{5, 5},
		{0, 0},
	}
	
	for _, test := range tests {
		result := AbsInt32(test.input)
		if result != test.expected {
			t.Errorf("AbsInt32(%d) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{5.5, 3.3, 5.5},
		{3.3, 5.5, 5.5},
		{5.5, 5.5, 5.5},
	}
	
	for _, test := range tests {
		result := Max(test.a, test.b)
		if result != test.expected {
			t.Errorf("Max(%f, %f) = %f, expected %f", test.a, test.b, result, test.expected)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{5.5, 3.3, 3.3},
		{3.3, 5.5, 3.3},
		{5.5, 5.5, 5.5},
	}
	
	for _, test := range tests {
		result := Min(test.a, test.b)
		if result != test.expected {
			t.Errorf("Min(%f, %f) = %f, expected %f", test.a, test.b, result, test.expected)
		}
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{5.4, 5},
		{5.5, 6},
		{5.6, 6},
		{-5.4, -5},
		{-5.5, -6},
		{-5.6, -6},
	}
	
	for _, test := range tests {
		result := Round(test.input)
		if result != test.expected {
			t.Errorf("Round(%f) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

func TestCeiling(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{5.1, 6},
		{5.9, 6},
		{5.0, 5},
		{-5.1, -5},
		{-5.9, -5},
	}
	
	for _, test := range tests {
		result := Ceiling(test.input)
		if result != test.expected {
			t.Errorf("Ceiling(%f) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{5.1, 5},
		{5.9, 5},
		{5.0, 5},
		{-5.1, -6},
		{-5.9, -6},
	}
	
	for _, test := range tests {
		result := Floor(test.input)
		if result != test.expected {
			t.Errorf("Floor(%f) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		x, y     float64
		expected float64
	}{
		{2, 3, 8},
		{5, 2, 25},
		{10, 0, 1},
	}
	
	for _, test := range tests {
		result := Pow(test.x, test.y)
		if result != test.expected {
			t.Errorf("Pow(%f, %f) = %f, expected %f", test.x, test.y, result, test.expected)
		}
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{4, 2},
		{9, 3},
		{16, 4},
		{0, 0},
	}
	
	for _, test := range tests {
		result := Sqrt(test.input)
		if result != test.expected {
			t.Errorf("Sqrt(%f) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

func TestTrigonometric(t *testing.T) {
	// Test Sin, Cos, Tan with known values
	tolerance := 1e-10
	
	// Sin(0) = 0
	if result := Sin(0); math.Abs(result-0) > tolerance {
		t.Errorf("Sin(0) = %f, expected 0", result)
	}
	
	// Cos(0) = 1
	if result := Cos(0); math.Abs(result-1) > tolerance {
		t.Errorf("Cos(0) = %f, expected 1", result)
	}
	
	// Tan(0) = 0
	if result := Tan(0); math.Abs(result-0) > tolerance {
		t.Errorf("Tan(0) = %f, expected 0", result)
	}
}

func TestSign(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{5.5, 1},
		{-5.5, -1},
		{0, 0},
	}
	
	for _, test := range tests {
		result := Sign(test.input)
		if result != test.expected {
			t.Errorf("Sign(%f) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestIsNaN(t *testing.T) {
	if !IsNaN(math.NaN()) {
		t.Error("Expected IsNaN(NaN) to be true")
	}
	
	if IsNaN(5.5) {
		t.Error("Expected IsNaN(5.5) to be false")
	}
}

func TestIsInfinity(t *testing.T) {
	if !IsInfinity(math.Inf(1)) {
		t.Error("Expected IsInfinity(+Inf) to be true")
	}
	
	if !IsInfinity(math.Inf(-1)) {
		t.Error("Expected IsInfinity(-Inf) to be true")
	}
	
	if IsInfinity(5.5) {
		t.Error("Expected IsInfinity(5.5) to be false")
	}
}