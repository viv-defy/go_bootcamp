package basic_number_filtering

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
