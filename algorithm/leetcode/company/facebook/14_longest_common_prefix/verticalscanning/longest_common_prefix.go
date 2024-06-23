package verticalscanning

// Compare prefix letter position by position with all string. Different from
// horizontal scanning, vertical scanning will compare a position with every
// string once and avoid the worst case where shortest string is in the end.
// And also we can only use O(1) space
// Time: O(m*n) - m is number of strings and n is the shortest length of strings
// Space: O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
