package dp

import "math"

// DP. F(i, j) represents the max coin we can gain by burst all balloons in [i, j]
// with any order. In order to get F(i, j). We can split this problem to two sub-problem.
// F(i, j) = max(F(i, k-1)+F(k+1, j)+nums[i-1]*nums[k]*nums[j+1]]) which means we select
// balloon k as the only balloon we haven't burst yet so if we burst balloon k, max # of
// coins we will get if we burst balloon k last is is max # of coins we get by bursting
// balloon [i, k-1] and [k+1, j]. Then plus nums[k]*nums[i-1]*num[j+1].
// Reference: https://www.youtube.com/watch?v=z3hu2Be92UA&t=461s
// Time: O(n^3)
// Space: O(n^2)
func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nums = append([]int{1}, append(nums, 1)...)
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
		if i > 0 && i < len(nums)-1 {
			dp[i][i] = nums[i-1] * nums[i] * nums[i+1]
		}
	}

	for i := 1; i < len(nums)-1; i++ {
		for j := i - 1; j >= 1; j-- {
			for k := j; k <= i; k++ {
				dp[j][i] = maxInt(dp[j][i], nums[j-1]*nums[k]*nums[i+1]+dp[j][k-1]+dp[k+1][i])
			}
		}
	}

	return dp[1][len(nums)-2]
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
