package test_generated

import (
	"test-avesha-agent/math"
	"testing"
)

func TestAdd(t *testing.T) {
	result := math.Add(5, 3)
	expected := 8
	if result != expected {
		t.Errorf("Addition test failed: expected %d but got %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := math.Subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Subtraction test failed: expected %d but got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := math.Multiply(6, 7)
	expected := 42
	if result != expected {
		t.Errorf("Multiplication test failed: expected %d but got %d", expected, result)
	}
}

func TestDivide(t *testing.T) {
	result, err := math.Divide(15, 3)
	if err != nil {
		t.Errorf("Division test failed: unexpected error %v", err)
	}
	expected := 5
	if result != expected {
		t.Errorf("Division test failed: expected %d but got %d", expected, result)
	}
}

func TestPower(t *testing.T) {
	result := math.Power(2, 5)
	expected := 32
	if result != expected {
		t.Errorf("Power test failed: expected %d but got %d", expected, result)
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := math.SquareRoot(25)
	if err != nil {
		t.Errorf("SquareRoot test failed: unexpected error %v", err)
	}
	expected := 5
	if result != expected {
		t.Errorf("SquareRoot test failed: expected %d but got %d", expected, result)
	}
}

func TestAbs(t *testing.T) {
	result := math.Abs(-10)
	expected := 10
	if result != expected {
		t.Errorf("Absolute test failed: expected %d but got %d", expected, result)
	}
}

