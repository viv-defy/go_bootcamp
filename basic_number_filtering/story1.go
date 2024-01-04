package basic_number_filtering

import (
	"fmt"
)

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
	fmt.Println("\nBasic Number Filtering - Story1")

	res := evenNumbers([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	expected := []int{0, 2, 4, 6, 8}
	test(res, expected, "case 1: {0, 1, 2, 3, 4, 5, 6, 7, 8}")
}
