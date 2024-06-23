package dp

import "math"

// The problems requires us to find the least # of perfect square number required to sum to n.
// For example, if n=12, it could be 4+4+4 which is 3.
// One way we can do is try func: f(n) = min(f(n-1^2), f(n-2^2), f(n-3^2), ...)
// But this has many duplicate computation. We can use dp to compute from 1, 2, ..., n
// Base case: f(1) = 1
// From 2 to n: f(n) = min(f(n-1)+f(1), f(n-2)+f(2), ..., f(1)+f(n-1))
// We see the symmetric pattern and we can improve to:
// From 2 to n: f(n) = min(f(n-1)+f(1), f(n-2)+f(2), ..., f(n/2)+f(n-n/2))
// But if we keep thinking about this problem, we would find do we need to
// check f(3)+f(n-3)? Will this be as good as f(4)+f(n-4)? The answer is no.
// So we can keep improve the dp:
// From 2 to n: f(n) = min(f(n-1^2)+f(1), f(n-2^2)+f(2^2), ..., f(n-a^2)+f(a^2))
// a is sqrt(n).
// Time: O(nlogn)
// Space: O(n)
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n
	}
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		a := int(math.Sqrt(float64(i)))
		if a*a == i {
			dp[i] = 1
		} else {
			for j := 1; j <= a; j++ {
				dp[i] = minInt(dp[i], 1+dp[i-j*j])
			}
		}
	}
	return dp[n]
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
