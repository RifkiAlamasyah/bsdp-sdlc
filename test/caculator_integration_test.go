// tests/calculator_integration_test.go
package tests

import (
	calculator "bsdp/internal/app/caculator"
	"testing"
)

func TestCalculatorIntegration(t *testing.T) {
	// Use multiple calculator functions in a sequence
	result := calculator.Add(1, calculator.Multiply(2, 3)) // 1 + (2 * 3) = 7
	expected := 7.0
	if result != expected {
		t.Errorf("Add(1, Multiply(2, 3)) = %v; want %v", result, expected)
	}

	// Test a more complex sequence
	result, err := calculator.Divide(calculator.Add(10, calculator.Subtract(5, 3)), 2) // (10 + (5 - 3)) / 2 = 6
	expected = 6.0
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Divide(Add(10, Subtract(5, 3)), 2) = %v; want %v", result, expected)
	}
}
