package conditioncheck

import "strings"

// Reference: https://leetcode.com/problems/swap-adjacent-in-lr-string/solution/
// Time: O(n)
// Space: O(n)
func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}
	if strings.Replace(start, "X", "", -1) != strings.Replace(end, "X", "", -1) {
		return false
	}
	// check left
	j := 0
	for i := 0; i < len(start); i++ {
		if start[i] == 'L' {
			for end[j] != 'L' && j < len(end) {
				j++
			}
			if i < j {
				return false
			}
			j++
		}
	}
	// check right
	j = 0
	for i := 0; i < len(end); i++ {
		if start[i] == 'R' {
			for end[j] != 'R' && j < len(end) {
				j++
			}
			if i > j {
				return false
			}
			j++
		}
	}
	return true
}
