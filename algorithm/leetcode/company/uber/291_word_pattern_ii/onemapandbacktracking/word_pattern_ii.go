package onemapandbacktracking

// DFS and backtracking. At the same time, we use a map to map a letter to a substring.
// If a letter and a string are matched, we can deal with the remaining sub pattern and sub string
// Otherwise, we need to dfs and enum all substring from remaining string. We also need a substring
// map to function as a set so that we can make sure a substring doesn't match with two different
// letters.
// Time: O(m*n) - m is length of pattern and n is length of string. For every dfs search from root
//       to leaf, we actually always need to process the who string but it might be distributed in
//       m calls (length of pattern). So the total is m*n. m is # of all possible root to leaf path.
//       The actually time would be less than that.
// Space: O(m+n) - the map
func wordPatternMatch(pattern string, str string) bool {
	return matchHelper(pattern, str, map[byte]string{}, map[string]bool{})
}

func matchHelper(pattern string, str string, patternMap map[byte]string, matchedStr map[string]bool) bool {
	if len(pattern) == 0 && len(str) == 0 {
		return true
	}
	if len(pattern) == 0 || len(str) == 0 {
		return false
	}
	l := pattern[0]
	if subStr, ok := patternMap[l]; ok {
		if len(str) < len(subStr) {
			return false
		}
		for i := 0; i < len(subStr); i++ {
			if subStr[i] != str[i] {
				return false
			}
		}
		return matchHelper(pattern[1:], str[len(subStr):], patternMap, matchedStr)
	} else {
		for i := 1; i <= len(str); i++ {
			if matchedStr[str[:i]] {
				continue
			}
			patternMap[l] = str[:i]
			matchedStr[str[:i]] = true
			if matchHelper(pattern[1:], str[i:], patternMap, matchedStr) {
				return true
			}
			matchedStr[str[:i]] = false
			delete(patternMap, l)
		}
	}
	return false
}
