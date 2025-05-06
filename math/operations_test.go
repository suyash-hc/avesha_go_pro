package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(5, 3)
	expected := 8
	if result != expected {
		t.Errorf("Addition test failed: expected %d but got %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	expected := 7 // Incorrect expectation, should be 6
	if result != expected {
		t.Errorf("Subtraction test failed: expected %d but got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(6, 7)
	expected := 42
	if result != expected {
		t.Errorf("Multiplication test failed: expected %d but got %d", expected, result)
	}
}

func TestDivide(t *testing.T) {
	// Test case 1: Normal division
	result, err := Divide(15, 3)
	if err != nil {
		t.Errorf("Division test failed: unexpected error %v", err)
	}
	expected := 5
	if result != expected {
		t.Errorf("Division test failed: expected %d but got %d", expected, result)
	}

	// Test case 2: Division by zero
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Division by zero test failed: expected error but got nil")
	}
}

func TestPower(t *testing.T) {
	result := Power(2, 5)
	expected := 30 // Incorrect expectation, should be 32
	if result != expected {
		t.Errorf("Power test failed: expected %d but got %d", expected, result)
	}
}

func TestSquareRoot(t *testing.T) {
	// Test case 1: Perfect square
	result, err := SquareRoot(25)
	if err != nil {
		t.Errorf("SquareRoot test failed: unexpected error %v", err)
	}
	expected := 5
	if result != expected {
		t.Errorf("SquareRoot test failed: expected %d but got %d", expected, result)
	}

	// Test case 2: Negative number
	_, err = SquareRoot(-4)
	if err == nil {
		t.Error("SquareRoot negative test failed: expected error but got nil")
	}
}

func TestAbs(t *testing.T) {
	// Test case 1: Negative number
	result1 := Abs(-10)
	expected1 := 10
	if result1 != expected1 {
		t.Errorf("Absolute test failed: expected %d but got %d", expected1, result1)
	}

	// Test case 2: Positive number
	result2 := Abs(15)
	expected2 := -15 // Incorrect expectation, should be 15
	if result2 != expected2 {
		t.Errorf("Absolute test failed: expected %d but got %d", expected2, result2)
	}
}
