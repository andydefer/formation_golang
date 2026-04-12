// Package mathutil provides basic mathematical utility functions.
package mathutil

import (
	"errors"
	"fmt"
)

// Addition returns the sum of two integers.
func Addition(a, b int) int {
	return a + b
}

// Multiplication returns the product of two integers.
func Multiplication(a, b int) int {
	return a * b
}

var ErrorDiv = errors.New("la division par zero est impossible")

// Division returns the rest of two float64
// It panics if b is equal to zero
func Division(a, b float64) (float64, error) { // ← retourne error, pas DivisionError
	if b == 0 {
		return 0, ErrorDivision{
			Dividende: a,
			Diviseur:  b,
			Err:       ErrorDiv,
		}
	}
	return a / b, nil
}

// Factorial returns n! (n factorial).
// It panics if n is negative because factorial is undefined for negative numbers.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial undefined for negative number: %d", n)
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// IsEven reports whether n is an even number.
func IsEven(n int) bool {
	return n%2 == 0
}

type ErrorDivision struct {
	Dividende float64
	Diviseur  float64
	Err       error // ← je renomme pour éviter confusion
}

// Error ErrorDivision implémente error
func (e ErrorDivision) Error() string {
	return fmt.Sprintf("division impossible: %.2f / %.2f - %v",
		e.Dividende, e.Diviseur, e.Err)
}
