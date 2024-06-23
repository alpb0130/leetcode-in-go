package dpfromstarttoend

// Time: O(n^2)
// Space: O(n^2)
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	max := 0
	var left, right int
	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			if i == j {
				dp[i][j] = true
			} else {
				dp[i][j] = s[i] == s[j] && (dp[i-1][j+1] || i-1 < j+1)
			}
			if dp[i][j] && i-j+1 > max {
				max = i - j + 1
				left, right = j, i
			}
		}
	}
	if max == 0 {
		return ""
	} else {
		return s[left : right+1]
	}
}
