package twomap

import "strings"

// Use two map to record the index mapping for pattern letter and string separately
// Time: O(n)
// Space: O(n)
func wordPattern(pattern string, str string) bool {
	patternMap := map[byte]int{}
	strMap := map[string]int{}
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if patternMap[pattern[i]] != strMap[strs[i]] {
			return false
		}
		// Be careful! It i+1 because we don't want to use 0.
		// We use 0 as a special init value
		patternMap[pattern[i]] = i + 1
		strMap[strs[i]] = i + 1
	}
	return true
}
