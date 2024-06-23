package dp

// DP. F(n) = F(n-c1)+F(n-c2)+...+F(n-ck).
// Followup: what if there is negative numbers in nums?? Do we add some limitations to the problem?
// One case we can imagine is if nums are [1, -1] and target is 1, there would be infinite possible
// solutions: 1, -1, 1, -1, ..., 1. In this case, we need to either require each element can be
// used only once or limit the length of result list. With those limitation, the problem can be
// solve using iii or iii.
// Time: O(n*target)
// Space: O(target)
func combinationSum4(nums []int, target int) int {
	if len(nums) == 0 || target == 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
