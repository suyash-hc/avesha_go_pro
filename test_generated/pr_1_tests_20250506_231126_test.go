package test_generated

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 5, 10, 15},
		{"Negative numbers", -5, -10, -15},
		{"Zero and positive number", 0, 10, 10},
		{"Zero and negative number", 0, -10, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 10, 5, 5},
		{"Negative numbers", -10, -5, -5},
		{"Zero and positive number", 10, 0, 10},
		{"Zero and negative number", -10, 0, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 5, 10, 50},
		{"Negative numbers", -5, -10, 50},
		{"Zero and positive number", 0, 10, 0},
		{"Zero and negative number", 0, -10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           int
		b           int
		expected    int
		expectedErr error
	}{
		{"Valid division", 10, 2, 5, nil},
		{"Division by zero", 10, 0, 0, errors.New("division by zero")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if result != tt.expected || err.Error() != tt.expectedErr.Error() {
				t.Errorf("Divide(%d, %d) = %d, %v; want %d, %v", tt.a, tt.b, result, err, tt.expected, tt.expectedErr)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 2, 3, 8},
		{"Negative exponent", 2, -3, 0},
		{"Zero exponent", 2, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Power(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Power(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		name        string
		a           int
		expected    int
		expectedErr error
	}{
		{"Perfect square", 16, 4, nil},
		{"Non-perfect square", 10, 3, nil},
		{"Negative number", -10, 0, errors.New("cannot calculate square root of negative number")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SquareRoot(tt.a)
			if result != tt.expected || err.Error() != tt.expectedErr.Error() {
				t.Errorf("SquareRoot(%d) = %d, %v; want %d, %v", tt.a, result, err, tt.expected, tt.expectedErr)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		expected int
	}{
		{"Positive number", 10, 10},
		{"Negative number", -10, 10},
		{"Zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.a)
			if result != tt.expected {
				t.Errorf("Abs(%d) = %d; want %d", tt.a, result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"First number larger", 10, 5, 10},
		{"Second number larger", 5, 10, 10},
		{"Equal numbers", 5, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

