package dp

// DP. f(n) = f(n+i)+isWord(n:n+i)
// Time: O(n^2)
// Space: O(n)
func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true
	for i := len(s); i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if dp[j+1] && wordMap[s[i:j+1]] {
				dp[i] = true
			}
		}
	}
	return dp[0]
}
