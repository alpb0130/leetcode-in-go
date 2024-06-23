package math

import "math"

// Time: O(sqrt(n))
// Space: O(1)
func consecutiveNumbersSum(N int) int {
	max := int(math.Sqrt(float64(2 * N)))
	res := 0
	for k := 1; k <= max; k++ {
		if (2*N)%k != 0 {
			continue
		}
		x := (2*N)/k - k + 1
		if x%2 != 0 || x/2 < 1 || x/2 > N {
			continue
		}
		res++
	}
	return res
}

// n, n+1, n+k-1 (start with n and k numbers)
// N = (k(2n+k-1))/2
// 2N = k*(2n+k-1)
// we need to find all integer k for a given N which n is also a positive number
// because 2n+k-1 is > k so we can limit our search space to sqrt(2N)
