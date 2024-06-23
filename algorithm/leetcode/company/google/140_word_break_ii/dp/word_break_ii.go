package dp

import "strings"

// 2D DP but we also want to store the tmp string list from last couple of steps
// Time: O(n^3) - Loop over the string nested and we also need to loop over the string list
// from last step
// Space: O(n^3)
func wordBreak(s string, wordDict []string) []string {
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([][]string, len(s)+1)
	dp[0] = []string{""}
	for i := 0; i < len(s); i++ {
		dp[i+1] = []string{}
		for j := 0; j <= i; j++ {
			if len(dp[j]) != 0 && wordMap[s[j:i+1]] {
				for _, str := range dp[j] {
					var tmp strings.Builder
					tmp.WriteString(str)
					if str != "" {
						tmp.WriteString(" ")
					}
					tmp.WriteString(s[j : i+1])
					dp[i+1] = append(dp[i+1], tmp.String())
				}
			}
		}
	}

	return dp[len(s)]
}
