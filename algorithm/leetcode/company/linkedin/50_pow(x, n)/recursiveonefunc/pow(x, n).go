package recursiveonefunc

// Divide and conquer but only use one function.
// Time: O(logn)
// Space: O(logn)
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
	prod := myPow(x, n/2)
	prod *= prod
	if n%2 == 1 {
		prod *= x
	}
	return prod
}
