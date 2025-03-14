package ginkgo

import "errors"

// Calculator is a struct that provides basic calculation functions
type Calculator struct{}

// Add adds two numbers
func (c *Calculator) Add(a, b int) int {
	return a + b
}

// Subtract subtracts the second number from the first
func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two numbers
func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

// Divide divides the first number by the second
// Returns an error if trying to divide by zero
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// ErrDivideByZero is an error returned when trying to divide by zero
var ErrDivideByZero = errors.New("cannot divide by zero")
