package linear

import "math"

// Be careful about overflow
// Time: O(n)
// Space: O(1)
func reverse(x int) int {
	res := 0
	for x != 0 {
		if res > math.MaxInt32/10 ||
			(res == math.MaxInt32/10 && math.MaxInt32%10 < x%10) ||
			res < math.MinInt32/10 ||
			(res == math.MinInt32/10 && math.MinInt32%10 > x%10) {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}
	return res
}
