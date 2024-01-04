package basic_number_filtering

import (
	"fmt"
)

func primeNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isPrime(v) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory3() {
	fmt.Println("\nBasic Number Filtering - Story3")

	res := primeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	expected := []int{2, 3, 5, 7, 11, 13}
	test(res, expected, "case 1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}")
}
