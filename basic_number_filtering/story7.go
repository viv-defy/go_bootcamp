package basic_number_filtering

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
