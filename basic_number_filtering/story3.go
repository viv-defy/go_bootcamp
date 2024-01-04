package basic_number_filtering

import (
	"fmt"
	"math"
)

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

func primeNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isPrime(v) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory3() {
	fmt.Println("Basic Number Filtering - Story3")
	res := primeNumbers(make([]int, 0))
	if len(res) == 0 {
		fmt.Println("PASS - case1: Empty list")
	} else {
		fmt.Println("FAIL - case1: Empty list")
	}

	res = primeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	expected := []int{2, 3, 5, 7, 11, 13}
	var failed bool
	for i, v := range expected {
		if res[i] != v {
			failed = true
			fmt.Println("FAIL - case2: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}")
			break
		}
	}
	if !failed {
		fmt.Println("PASS - case2: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}")
	}
}
