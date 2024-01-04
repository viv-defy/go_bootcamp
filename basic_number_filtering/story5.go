package basic_number_filtering

func IsEvenAndMultipleOf5(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isEven(v) && isMultipleOf(5)(v) {
			res = append(res, v)
		}
	}
	return res
}
