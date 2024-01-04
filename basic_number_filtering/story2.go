package basic_number_filtering

import (
	"fmt"
)

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
	fmt.Println("\nBasic Number Filtering - Story2")

	res := oddNumbers([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	expected := []int{1, 3, 5, 7}
	test(res, expected, "case 1: {0, 1, 2, 3, 4, 5, 6, 7, 8}")
}
