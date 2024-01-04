package basic_number_filtering

import "fmt"

func isMultipleOf(n, m int) bool {
	return n%m == 0
}

func isEvenAndMultipleOf5(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if even(v) && isMultipleOf(v, 5) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory5() {
	fmt.Println("Basic Number Filtering - Story5")
	res := isEvenAndMultipleOf5(make([]int, 0))
	if len(res) == 0 {
		fmt.Println("PASS - case1: Empty list")
	} else {
		fmt.Println("FAIL - case1: Empty list")
	}

	res = isEvenAndMultipleOf5([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	expected := []int{10, 20}
	var failed bool
	for i, v := range expected {
		if res[i] != v {
			failed = true
			fmt.Println("FAIL - case2: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
			break
		}
	}
	if !failed {
		fmt.Println("PASS - case2: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
	}
}
