package basic_number_filtering

func OddNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isOdd(v) {
			res = append(res, v)
		}
	}
	return res
}
