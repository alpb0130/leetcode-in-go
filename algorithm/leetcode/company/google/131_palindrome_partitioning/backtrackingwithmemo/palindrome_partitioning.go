package backtrackingwithmemo

// DFS with memorization.
// Time: O(2^n)
// Space: O(n*2^n)
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	dp := map[string][][]string{"": [][]string{{}}}
	// dfs
	return dfs(s, dp)
}

func dfs(s string, dp map[string][][]string) [][]string {
	if dp[s] != nil {
		return dp[s]
	}
	res := [][]string{}
	for i := 0; i < len(s); i++ {
		if isPalindrome(s[0 : i+1]) {
			tmp := dfs(s[i+1:], dp)
			for _, list := range tmp {
				res = append(res, append([]string{s[0 : i+1]}, list...))
			}
		}
	}
	dp[s] = res
	return res
}

func isPalindrome(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
