package basic_number_filtering

import "fmt"

func oddNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if odd(v) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory2() {
	fmt.Println("Basic Number Filtering - Story2")
	res := oddNumbers(make([]int, 0))
	if len(res) == 0 {
		fmt.Println("PASS - case1: Empty list")
	} else {
		fmt.Println("FAIL - case1: Empty list")
	}

	res = oddNumbers([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	i := 0
	for ; i < 4; i++ {
		if res[i] != i*2+1 {
			fmt.Println("FAIL - case2: {0, 1, 2, 3, 4, 5, 6, 7, 8}")
			break
		}
	}
	if i == 4 {
		fmt.Println("PASS - case2: {0, 1, 2, 3, 4, 5, 6, 7, 8}")
	}
}
