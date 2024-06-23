package slidingwindow

import "math"

// Sliding window. Every time we see there are more than 2 distinct letters in the window,
// we just move the left pointer and reduce the window size to make sure there are
// at most 2 characters
// Time: O(n)
// Space: O(1)
func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) <= 2 {
		return len(s)
	}
	i, j := 0, 0
	windowMap := map[byte]int{}
	res := 0
	for j < len(s) {
		windowMap[s[j]]++
		for len(windowMap) > 2 && i <= j {
			windowMap[s[i]]--
			if windowMap[s[i]] == 0 {
				delete(windowMap, s[i])
			}
			i++
		}
		res = maxInt(res, j-i+1)
		j++
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
