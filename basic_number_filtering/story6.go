package basic_number_filtering

func IsOddAndMultipleOf3GreaterThan10(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if v > 10 && isOdd(v) && isMultipleOf(3)(v) {
			res = append(res, v)
		}
	}
	return res
}
