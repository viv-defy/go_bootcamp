package basic_number_filtering

import "fmt"

func oddAndPrimeNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isPrime(v) && isOdd(v) {
			res = append(res, v)
		}
	}
	return res
}

func TestStory4() {
	fmt.Println("\nBasic Number Filtering - Story4")

	res := oddAndPrimeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8})
	expected := []int{3, 5, 7}
	test(res, expected, "case 1: {1, 2, 3, 4, 5, 6, 7, 8}")
}
