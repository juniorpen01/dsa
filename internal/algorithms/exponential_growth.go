package algorithms

func ExponentialGrowth(n, factor, days int) []int {
	res := make([]int, 0, days+1)
	for range days + 1 {
		res = append(res, n)
		n *= factor
	}
	return res
}
