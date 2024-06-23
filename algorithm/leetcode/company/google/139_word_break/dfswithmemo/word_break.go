package dfswithmemo

// DFS + memo
// Time: O(n^2)
// Space: O(n)
func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}
	// dfs string
	return dfs(s, wordMap, map[string]bool{})
}

func dfs(word string, wordMap map[string]bool, dp map[string]bool) bool {
	if valid, ok := dp[word]; ok {
		return valid
	}
	if len(word) == 0 {
		return true
	}
	for i := 0; i < len(word); i++ {
		if !wordMap[word[:i+1]] {
			continue
		}
		if dfs(word[i+1:], wordMap, dp) {
			return true
		}
	}
	dp[word] = false
	return false
}
