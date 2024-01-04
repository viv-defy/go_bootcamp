package basic_number_filtering

func PrimeNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isPrime(v) {
			res = append(res, v)
		}
	}
	return res
}
