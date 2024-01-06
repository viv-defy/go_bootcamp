package basic_number_filtering

func EvenNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isEven(v) {
			res = append(res, v)
		}
	}
	return res
}

func OddNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isOdd(v) {
			res = append(res, v)
		}
	}
	return res
}

func PrimeNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isPrime(v) {
			res = append(res, v)
		}
	}
	return res
}

func OddAndPrimeNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isPrime(v) && isOdd(v) {
			res = append(res, v)
		}
	}
	return res
}

func IsEvenAndMultipleOf5(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isEven(v) && isMultipleOf(5)(v) {
			res = append(res, v)
		}
	}
	return res
}

func IsOddAndMultipleOf3GreaterThan10(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if v > 10 && isOdd(v) && isMultipleOf(3)(v) {
			res = append(res, v)
		}
	}
	return res
}

func All(nums []int, conditions []func(int) bool) []int {
	res := make([]int, 0)
	for _, v := range nums {
		var doesNotSatisfy bool
		for _, f := range conditions {
			if !f(v) {
				doesNotSatisfy = true
				break
			}
		}
		if !doesNotSatisfy {
			res = append(res, v)
		}
	}
	return res
}

func Any(nums []int, conditions []func(int) bool) []int {
	res := make([]int, 0)
	for _, v := range nums {
		var satisfies bool
		for _, f := range conditions {
			if f(v) {
				satisfies = true
				break
			}
		}
		if satisfies {
			res = append(res, v)
		}
	}
	return res
}
