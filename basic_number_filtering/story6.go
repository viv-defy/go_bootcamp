package basic_number_filtering

import "fmt"

func isOddAndMultipleOf3GreaterThan10(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if v > 10 && odd(v) && isMultipleOf(v, 3) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory6() {
	fmt.Println("\nBasic Number Filtering - Story6")

	res := isOddAndMultipleOf3GreaterThan10([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	expected := []int{15}
	test(res, expected, "case 1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
}
