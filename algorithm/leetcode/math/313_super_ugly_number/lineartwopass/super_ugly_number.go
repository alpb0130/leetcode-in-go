package lineartwopass

import "math"

// Time: O(nk) - k is the number of primes
// Space: O(k) - priority queue
func nthSuperUglyNumber(n int, primes []int) int {
	uglys := []int{1}
	indexs := make([]int, len(primes))
	for i := 1; i < n; i++ {
		min := math.MaxInt32
		for j := 0; j < len(primes); j++ {
			if uglys[indexs[j]]*primes[j] < min {
				min = uglys[indexs[j]] * primes[j]
			}
		}
		for j := 0; j < len(primes); j++ {
			if uglys[indexs[j]]*primes[j] == min {
				indexs[j]++
			}
		}
		uglys = append(uglys, min)
	}
	return uglys[n-1]
}
