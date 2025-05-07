package test_generated

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Addition failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Subtraction failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(7, 3)
	expected := 21
	if result != expected {
		t.Errorf("Multiplication failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(15, 3)
	if err != nil {
		t.Errorf("Error dividing: %v", err)
	}
	expected := 5
	if result != expected {
		t.Errorf("Division failed. Expected: %d, Got: %d", expected, result)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Division by zero should return an error")
	}
}

func TestPower(t *testing.T) {
	result := Power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Power calculation failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := SquareRoot(16)
	if err != nil {
		t.Errorf("Error calculating square root: %v", err)
	}
	expected := 4
	if result != expected {
		t.Errorf("Square root calculation failed. Expected: %d, Got: %d", expected, result)
	}

	_, err = SquareRoot(-1)
	if err == nil {
		t.Error("Calculating square root of negative number should return an error")
	}
}

func TestAbs(t *testing.T) {
	result := Abs(-5)
	expected := 5
	if result != expected {
		t.Errorf("Absolute value calculation failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestMax(t *testing.T) {
	result := Max(10, 7)
	expected := 10
	if result != expected {
		t.Errorf("Max calculation failed. Expected: %d, Got: %d", expected, result)
	}
}

