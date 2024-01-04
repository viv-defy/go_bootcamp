package basic_number_filtering

import (
	"math"
)

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isGreaterThan(n int) func(int) bool {
	return func(x int) bool {
		return x > n
	}
}

func isLessThan(n int) func(int) bool {
	return func(x int) bool {
		return x < n
	}
}

func isMultipleOf(n int) func(int) bool {
	return func(x int) bool {
		return x%n == 0
	}
}
