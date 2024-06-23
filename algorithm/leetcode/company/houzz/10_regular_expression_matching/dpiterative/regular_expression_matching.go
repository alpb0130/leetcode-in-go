package dpiterative

// Prefix matching till the who string and pattern.
// Reference: https://leetcode.com/problems/regular-expression-matching/discuss/5651/Easy-DP-Java-Solution-with-detailed-Explanation
// Time: O(len(s) * len(p))
// Space: O(len(s) * len(p))
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i < len(p); i++ {
		if p[i] == '*' && i-1 >= 0 {
			dp[0][i+1] = dp[0][i-1]
		}
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' && j-1 >= 0 && (p[j-1] == s[i] || p[j-1] == '.') {
				dp[i+1][j+1] = dp[i+1][j] || dp[i+1][j-1]
			} else if p[j] == '*' {
				dp[i+1][j+1] = dp[i+1][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
