package basic_number_filtering

import "fmt"

func isOddAndMultipleOf3(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if v > 10 && odd(v) && isMultipleOf(v, 3) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory6() {
	fmt.Println("Basic Number Filtering - Story6")
	res := isOddAndMultipleOf3(make([]int, 0))
	if len(res) == 0 {
		fmt.Println("PASS - case1: Empty list")
	} else {
		fmt.Println("FAIL - case1: Empty list")
	}

	res = isOddAndMultipleOf3([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	expected := []int{15}
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
