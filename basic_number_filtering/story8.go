package basic_number_filtering

import (
	"fmt"
)

func Any(nums []int, conditions []func(int) bool) []int {
	res := make([]int, 0)
	for _, v := range nums {
		var satisfies bool
		for _, f := range conditions {
			if f(v) {
				satisfies = true
				break
			}
		}
		if satisfies {
			res = append(res, v)
		}
	}
	return res
}

func TestStory8() {
	fmt.Println("Basic Number Filtering - Story8")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	conditions := []func(int) bool{isPrime, greaterThan(15), multipleOf(5)}
	res := Any(nums, conditions)
	expected := []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}
	var notSatisfied bool
	for i, v := range expected {
		if res[i] != v {
			notSatisfied = true
			break
		}
	}
	if notSatisfied || len(res) != len(expected) {
		fmt.Println("FAIL - case1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
		fmt.Println("\t conditions: {isPrime, greaterThan(15), multipleOf(5)}")
	} else {
		fmt.Println("PASS - case1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
		fmt.Println("\t conditions: {odd, greaterThan(5), multipleOf(3)}")
	}

	conditions = []func(int) bool{lessThan(6), multipleOf(3)}
	res = Any(nums, conditions)
	expected = []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18}
	notSatisfied = false
	for i, v := range expected {
		if res[i] != v {
			notSatisfied = true
			break
		}
	}
	if notSatisfied || len(res) != len(expected) {
		fmt.Println("FAIL - case2: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
		fmt.Println("\t conditions: {lessThan(6), multipleOf(3)}")
	} else {
		fmt.Println("PASS - case2: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}")
		fmt.Println("\t conditions: {lessThan(6), multipleOf(3)}")
	}
}
