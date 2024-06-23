package math1

import "math"

// Time: O(sqrt(n))
// Space: O(1)
func consecutiveNumbersSum(N int) int {
	max := int(math.Sqrt(float64(2 * N)))
	res := 0
	for k := 1; k <= max; k++ {
		x := N - k*(k-1)/2
		if x%k != 0 || x/k < 1 || x/k > N {
			continue
		}
		res++
	}
	return res
}

// n, n+1, n+k-1 (start with n and k numbers)
// N = (k(2n+k-1))/2 = nk+k(k-1)/2
// 2N = k*(2n+k-1) = 2nk+k(k-1)
// we need to find all integer k for a given N which n is also a positive number
// because 2n+k-1 is > k so we can limit our search space to sqrt(2N)
