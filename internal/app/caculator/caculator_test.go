// internal/calculator/calculator_test.go
package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5.0
	if result != expected {
		t.Errorf("Add(2, 3) = %v; want %v", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2.0
	if result != expected {
		t.Errorf("Subtract(5, 3) = %v; want %v", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	expected := 6.0
	if result != expected {
		t.Errorf("Multiply(2, 3) = %v; want %v", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(6, 3)
	expected := 2.0
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Divide(6, 3) = %v; want %v", result, expected)
	}

	// Test divide by zero
	_, err = Divide(6, 0)
	if err == nil {
		t.Error("expected error when dividing by zero")
	}
}
