package recursivetwofunc

// Divide and conquer.
// Time: O(logn)
// Space: O(logn)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if x == 0 {
		return 0
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return powHelper(x, n)
}

func powHelper(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	res := 1.0
	if n%2 == 1 {
		res *= x
		n--
	}
	prod := powHelper(x, n/2)
	res *= prod * prod
	return res
}
