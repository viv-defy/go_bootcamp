package basic_number_filtering

import "fmt"

func oddAndPrimeNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isPrime(v) && odd(v) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory4() {
	fmt.Println("Basic Number Filtering - Story4")
	res := oddAndPrimeNumbers(make([]int, 0))
	if len(res) == 0 {
		fmt.Println("PASS - case1: Empty list")
	} else {
		fmt.Println("FAIL - case1: Empty list")
	}

	res = oddAndPrimeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8})
	expected := []int{3, 5, 7}
	var failed bool
	for i, v := range expected {
		if res[i] != v {
			failed = true
			fmt.Println("FAIL - case2: {1, 2, 3, 4, 5, 6, 7, 8}")
			break
		}
	}
	if !failed {
		fmt.Println("PASS - case2: {1, 2, 3, 4, 5, 6, 7, 8}")
	}
}
