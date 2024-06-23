package dp

// DP.
// A slightly better solution: http://yucoding.blogspot.com/2013/02/leetcode-question-123-wildcard-matching.html
// Time: O(m*n)
// Space: O(m*n)
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 1; i < len(dp[0]); i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			// if pattern byte is equal to string byte or pattern byte is '?'
			// we match string and pattern and check dp[i-1][j-1]
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			}
			// if pattern byte is '*', it can
			// 1. don't match the string byte - we check dp[i][j-1]
			// 2. only match the string byte - we check dp[i-1][j-1]
			// 3. '*' also match the previous string byte - we check dp[i-1][j]
			if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j-1] || dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
