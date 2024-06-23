package slidingwindow

// Use a map to record whether a char has appeared. If not, add the char to map and update
// long length. Otherwise, remove start char until current char is not in the map (substring)
// and then update the longest length. This requires total 2n move.
// Time: O(n)
// Space: O(1)
func lengthOfLongestSubstring(s string) int {
	res := 0
	charMap := map[byte]bool{}
	left, right := 0, 0
	for right < len(s) {
		if charMap[s[right]] {
			for charMap[s[right]] && left < right {
				charMap[s[left]] = false
				left++
			}
		}
		charMap[s[right]] = true
		if right-left+1 > res {
			res = right - left + 1
		}
		right++
	}
	return res
}
