// Generated tests for PR 1
// Generated on: 2025-05-06 11:05:11

package main

import (
    "testing"
)


// Test: unnamed_test
Here are the tests for the functions in the "math" package:

```go
package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Addition failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 5)
	expected := 5
	if result != expected {
		t.Errorf("Subtraction failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 5)
	expected := 15
	if result != expected {
		t.Errorf("Multiplication failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	expected := 5
	if err != nil || result != expected {
		t.Errorf("Division failed. Expected: %d, Got: %d", expected, result)
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
	result, err := SquareRoot(9)
	expected := 3
	if err != nil || result != expected {
		t.Errorf("Square root calculation failed. Expected: %d, Got: %d", expected, result)
	}
}

func TestAbs(t *testing.T) {
	result := Abs(-5)
	expected := 5
	if result != expected {
		t.Errorf("Absolute value calculation failed. Expected: %d, Got: %d", expected, result)
	}
}
```
