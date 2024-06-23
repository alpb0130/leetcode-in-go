package slidingwindowbetter

import "math"

// Use a map to record whether a char has appeared and its index. If not, add the char to map
// and update long length. Otherwise, remove the seen char from the map and update the left bound.
// One thing needed to be careful is we actually skip a bunch of chars which are before the
// repeated char but not remove them from the map. So next time we check a seen char, we need
// to tell whether it's a true repeated or not. Check that by compare their index with current
// left index. If current left larger than them, it means they are all fake repeat and we don't
// update left. Otherwise we update left. After we update left, we can update the substring length.
// Time: O(n)
// Space: O(1)
func lengthOfLongestSubstring(s string) int {
	res := 0
	charIdxMap := map[byte]int{}
	// left start with -1 to deal with the case where we cannot enter the if case in for loop
	// For example, it will not work for "a" if we set left to 0.
	left, right := -1, 0
	for right < len(s) {
		if idx, ok := charIdxMap[s[right]]; ok {
			// because when we remove a seen char from the map, we skip a bunch
			// of chars before it without removing the index in map, we need to
			// update left also consider whether the seen char's index is smaller
			// than current left. If it is, we keep the current left.
			left = maxInt(left, idx)
		}
		res = maxInt(res, right-left)
		charIdxMap[s[right]] = right
		right++
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
