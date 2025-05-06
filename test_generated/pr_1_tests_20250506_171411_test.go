package test_generated

import (
	"test-avesha-agent/math"
	"testing"
)

func TestAdd(t *testing.T) {
	result := math.Add(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Addition test failed: expected %d but got %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := math.Subtract(10, 3)
	expected := 7
	if result != expected {
		t.Errorf("Subtraction test failed: expected %d but got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := math.Multiply(4, 6)
	expected := 24
	if result != expected {
		t.Errorf("Multiplication test failed: expected %d but got %d", expected, result)
	}
}

func TestDivide(t *testing.T) {
	result, err := math.Divide(10, 2)
	expected := 5
	if result != expected || err != nil {
		t.Errorf("Division test failed: expected %d but got %d with error %v", expected, result, err)
	}
}

func TestPower(t *testing.T) {
	result := math.Power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Power test failed: expected %d but got %d", expected, result)
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := math.SquareRoot(16)
	expected := 4
	if result != expected || err != nil {
		t.Errorf("SquareRoot test failed: expected %d but got %d with error %v", expected, result, err)
	}
}

func TestAbs(t *testing.T) {
	result := math.Abs(-5)
	expected := 5
	if result != expected {
		t.Errorf("Absolute test failed: expected %d but got %d", expected, result)
	}
}

