package linearonepass

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	uglys := []int{1}
	indexs := make([]int, len(primes))
	min := 1
	for i := 1; i < n; i++ {
		curMin := math.MaxInt32
		for j := 0; j < len(primes); j++ {
			if uglys[indexs[j]]*primes[j] == min {
				indexs[j]++
			}
			if uglys[indexs[j]]*primes[j] < curMin {
				curMin = primes[j] * uglys[indexs[j]]
			}
		}
		min = curMin
		uglys = append(uglys, min)
	}
	return uglys[n-1]
}
