package basic_number_filtering

import "fmt"

func even(n int) bool {
	return n%2 == 0
}

func evenNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if even(v) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory1() {
	fmt.Println("Basic Number Filtering - Story1")
	res := evenNumbers(make([]int, 0))
	if len(res) == 0 {
		fmt.Println("PASS - case1: Empty list")
	} else {
		fmt.Println("FAIL - case1: Empty list")
	}

	res = evenNumbers([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	i := 0
	for ; i < 5; i++ {
		if res[i] != i*2 {
			fmt.Println("FAIL - case2: {0, 1, 2, 3, 4, 5, 6, 7, 8}")
			break
		}
	}
	if i == 5 {
		fmt.Println("PASS - case2: {0, 1, 2, 3, 4, 5, 6, 7, 8}")
	}
}
