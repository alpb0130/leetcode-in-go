package iterative

// Divide and conquer. A little bit tricky.
// This is a bottom-up method. If we
// Time: O(logn)
// Space: O(1)
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	res := 1.0
	prod := x
	for n > 0 {
		if n%2 == 1 {
			res *= prod
		}
		prod = prod * prod
		n /= 2
	}
	return res
}
