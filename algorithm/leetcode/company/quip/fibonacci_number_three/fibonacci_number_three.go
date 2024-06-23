package fibonacci_number_three

func Fibonacci(N int) int {
	if N <= 3 {
		return 1
	}
	dp := make([]int, N+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < N; i++ {
		dp[i] = dp[i-1]+dp[i-2]+dp[i-3]
	}
	return dp[N-1]
}
