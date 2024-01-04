package basic_number_filtering

func EvenNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isEven(v) {
			res = append(res, v)
		}
	}
	return res
}
