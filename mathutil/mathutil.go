// Package mathutil provides basic mathematical utility functions.
package mathutil

import "fmt"

// Addition returns the sum of two integers.
func Addition(a, b int) int {
	return a + b
}

// Multiplication returns the product of two integers.
func Multiplication(a, b int) int {
	return a * b
}

// Factorial returns n! (n factorial).
// It panics if n is negative because factorial is undefined for negative numbers.
// mathutil/mathutil.go
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
