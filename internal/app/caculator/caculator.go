// internal/calculator/calculator.go
package calculator

import "errors"

// Add two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract two numbers
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide two numbers, returns an error if divided by zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
