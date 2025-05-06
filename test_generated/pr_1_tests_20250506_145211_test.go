// Generated tests for PR 1
// Generated on: 2025-05-06 14:52:11


// Test: unnamed_test
The generated tests are:

```go
package main_test

import (
	"testing"
	"test-avesha-agent/math"
)

func TestAdd(t *testing.T) {
	result := math.Add(10, 5)
	if result != 15 {
		t.Errorf("Addition test failed. Expected: 15, Got: %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := math.Subtract(10, 5)
	if result != 5 {
		t.Errorf("Subtraction test failed. Expected: 5, Got: %d", result)
	}
}

func TestMultiply(t *testing.T) {
	result := math.Multiply(10, 5)
	if result != 50 {
		t.Errorf("Multiplication test failed. Expected: 50, Got: %d", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := math.Divide(10, 5)
	if err != nil {
		t.Errorf("Division test failed with error: %s", err.Error())
	}
	if result != 2 {
		t.Errorf("Division test failed. Expected: 2, Got: %d", result)
	}
}

func TestPower(t *testing.T) {
	result := math.Power(2, 3)
	if result != 8 {
		t.Errorf("Power test failed. Expected: 8, Got: %d", result)
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := math.SquareRoot(16)
	if err != nil {
		t.Errorf("Square root test failed with error: %s", err.Error())
	}
	if result != 4 {
		t.Errorf("Square root test failed. Expected: 4, Got: %d", result)
	}
}

func TestAbs(t *testing.T) {
	result := math.Abs(-5)
	if result != 5 {
		t.Errorf("Absolute value test failed. Expected: 5, Got: %d", result)
	}
}
```
