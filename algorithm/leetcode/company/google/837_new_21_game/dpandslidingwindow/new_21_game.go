package dpandslidingwindow

// Start from 0 point with possibility 1.0. Then at every point, the possibility x
// is equal to f(x) = (1/W)*(f(x-1)+f(x-2)+...+f(x-W)). We can improve that by using
// sliding window. If i is larger or equal to, we contribute the possibility to the final
// results. Also, if i >= K, we don't contribute the possibility to accumulative sum
// Time: O(N)
// Space: O(N)
func new21Game(N int, K int, W int) float64 {
	if K-1+W <= N || K == 0 {
		return 1.0
	}
	dp := make([]float64, N+1)
	dp[0] = 1.0
	res := 0.0
	ratio := float64(1) / float64(W)
	sum := 1.0
	for i := 1; i <= N; i++ {
		dp[i] += ratio * sum
		if i >= K {
			res += dp[i]
		} else {
			sum += dp[i]
		}
		if i-W >= 0 {
			sum -= dp[i-W]
		}
	}
	return res
}
