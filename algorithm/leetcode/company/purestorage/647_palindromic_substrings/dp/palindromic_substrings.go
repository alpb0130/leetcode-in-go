package dp

// The brute force solution is to check all substrings (O(n^s) substrings) and for check it takes
// linear time. So the time complexity is O(n^3). But we can see many substrings are checked more
// than once. Say, for "abba", when we check "abba", we check "bb" once but we would check "bb" one
// more time when we check substring "bb". So we can come up with dynamic programming.
// Take care of some edge cases.
// Time: O(n^2)
// Space: O(n^2)
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		// Base case. If only one letter, it should be palindrome. Don't forget to update res!!!
		dp[i][i] = true
		res++
	}
	for i := 0; i < len(dp); i++ {
		for j := i - 1; j >= 0; j-- {
			// If dp[j+1][i-1] is true of i, j cross path
			dp[j][i] = s[j] == s[i] && (j+1 > i-1 || dp[j+1][i-1])
			if dp[j][i] {
				res++
			}
		}
	}
	return res
}
