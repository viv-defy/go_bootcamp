package basic_number_filtering

import "math"

func even(n int) bool {
	return n%2 == 0
}

func odd(n int) bool {
	return !even(n)
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

func greaterThan(n int) func(int) bool {
	return func(x int) bool {
		return x > n
	}
}

func lessThan(n int) func(int) bool {
	return func(x int) bool {
		return x < n
	}
}

func multipleOf(n int) func(int) bool {
	return func(x int) bool {
		return x%n == 0
	}
}
