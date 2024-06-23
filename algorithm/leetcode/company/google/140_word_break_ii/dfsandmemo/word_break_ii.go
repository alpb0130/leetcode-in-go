package dfsandmemo

import "strings"

// Time: O(n^3)
// Space: O(n^2)
func wordBreak(s string, wordDict []string) []string {
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}
	return dfs(s, wordMap, map[string][]string{})
}

func dfs(s string, wordMap map[string]bool, memo map[string][]string) []string {
	if len(memo[s]) != 0 {
		return memo[s]
	}
	if s == "" {
		return []string{""}
	}
	res := []string{}
	for word := range wordMap {
		if strings.HasPrefix(s, word) {
			strs := dfs(s[len(word):], wordMap, memo)
			for _, str := range strs {
				if str == "" {
					res = append(res, word)
				} else {
					res = append(res, word+" "+str)
				}

			}
		}
	}
	memo[s] = res
	return res
}
