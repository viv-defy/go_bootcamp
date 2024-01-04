package basic_number_filtering

func OddAndPrimeNumbers(nums []int) []int {
	res := make([]int, 0)
	for _, v := range nums {
		if isPrime(v) && isOdd(v) {
			res = append(res, v)
		}
	}
	return res
}
