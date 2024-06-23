package twomapandbacktracking

// DFS and backtracking. At the same time, we use two map to map a letter/substring to a index.
// If a letter and a string has the same dst index, it means they are matched (or both index is 0
// and we know they haven't matched with any index). If a letter is not matched with current
// sub string, we skip the down stream check. Otherwise, we match them and narrow down our string
// search space.
// Time: O(m*n) - m is length of pattern and n is length of string. For every dfs search from root
//       to leaf, we actually always need to process the who string but it might be distributed in
//       m calls (length of pattern). So the total is m*n. m is # of all possible root to leaf path.
//       The actually time would be less than that.
// Space: O(m+n) - the map
func wordPatternMatch(pattern string, str string) bool {
	return matchHelper(pattern, str, 0, 0, map[byte]int{}, map[string]int{})
}

func matchHelper(pattern string, str string, pIdx, sIdx int, patternMap map[byte]int, matchedStr map[string]int) bool {
	if len(pattern) == pIdx && len(str) == sIdx {
		return true
	}
	if len(pattern) == pIdx || len(str) == sIdx {
		return false
	}
	l := pattern[pIdx]
	for i := 1; i <= len(str[sIdx:]); i++ {
		old := patternMap[l]
		if patternMap[l] == matchedStr[str[sIdx:sIdx+i]] {
			patternMap[l] = pIdx + 1
			matchedStr[str[sIdx:sIdx+i]] = pIdx + 1
			if matchHelper(pattern, str, pIdx+1, sIdx+i, patternMap, matchedStr) {
				return true
			}
			patternMap[l] = old
			matchedStr[str[sIdx:sIdx+i]] = old
		}
	}
	return false
}
