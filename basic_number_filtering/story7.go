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
	fmt.Println("Basic Number Filtering - Story7")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	conditions := []func(int) bool{odd, greaterThan(5), multipleOf(3)}
	res := All(nums, conditions)
	expected := []int{9, 15}
	var notSatisfied bool
	for i, v := range expected {
		if res[i] != v {
			notSatisfied = true
			fmt.Println("FAIL - case1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
			fmt.Println("\t conditions: {odd, greaterThan(5), multipleOf(3)}")
			break
		}
	}
	if !notSatisfied {
		fmt.Println("PASS - case1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
		fmt.Println("\t conditions: {odd, greaterThan(5), multipleOf(3)}")
	}

	conditions = []func(int) bool{even, lessThan(15), multipleOf(3)}
	res = All(nums, conditions)
	expected = []int{6, 12}
	notSatisfied = false
	for i, v := range expected {
		if res[i] != v {
			notSatisfied = true
			fmt.Println("FAIL - case2: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
			fmt.Println("\t conditions: {even, lessThan(15), multipleOf(3)}")
			break
		}
	}
	if !notSatisfied {
		fmt.Println("PASS - case2: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
		fmt.Println("\t conditions: {even, lessThan(15), multipleOf(3)}")
	}
}
