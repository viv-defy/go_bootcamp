package basic_number_filtering

import (
	"fmt"
)

func All(nums []int, conditions []func(int) bool) []int {
	res := make([]int, 0)
	for _, v := range nums {
		var doesNotSatisfy bool
		for _, f := range conditions {
			if !f(v) {
				doesNotSatisfy = true
				break
			}
		}
		if !doesNotSatisfy {
			res = append(res, v)
		}
	}
	return res
}

func TestStory7() {
	fmt.Println("\nBasic Number Filtering - Story7")

	conditions := []func(int) bool{isOdd, isGreaterThan(5), isMultipleOf(3)}
	res := All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, conditions)
	expected := []int{9, 15}
	test(res, expected, "case 1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}; conditions: {odd, greaterThan(5), multipleOf(3)}")

	conditions = []func(int) bool{isEven, isLessThan(15), isMultipleOf(3)}
	res = All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, conditions)
	expected = []int{6, 12}
	test(res, expected, "case 1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}; conditions: {even, lessThan(15), multipleOf(3)}")
}
