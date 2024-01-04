package basic_number_filtering

import "fmt"

func isEvenAndMultipleOf5(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isEven(v) && isMultipleOf(5)(v) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory5() {
	fmt.Println("\nBasic Number Filtering - Story5")

	res := isEvenAndMultipleOf5([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	expected := []int{10, 20}
	test(res, expected, "case 1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
}
