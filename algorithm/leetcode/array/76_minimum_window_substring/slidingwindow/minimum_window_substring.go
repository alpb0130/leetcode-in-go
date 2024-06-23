package slidingwindow

// Sliding window. Use two pointers. The right pointer is to expand the substring and make it
// eligible. The left pointer is to contact the substring until it's not eligible.
// The tricky part here is how to valid the substring in O(1) time. We can easily do that in
// O(S) by go through the char map of s but a more efficient way is to keep track of that
// as we process the string. Please refer to the solution below.
// Reference: https://leetcode.com/problems/minimum-window-substring/solution/
// Time: O(S+T) - S is the length of s and T is the length of T
// Space: O(# unique chars in s and t)
func minWindow(s string, t string) string {
	charMap := map[byte]int{}
	for _, c := range t {
		charMap[byte(c)]++
	}
	charNum := len(charMap)
	min := len(s) + 1
	l, r := 0, 0
	windowMap := map[byte]int{}
	charCnt := 0
	start, end := 0, 0
	for end < len(s) {
		windowMap[s[end]]++
		if charMap[s[end]] != 0 && windowMap[s[end]] == charMap[s[end]] {
			charCnt++
		}
		for start <= end && charCnt == charNum {
			if end-start+1 < min {
				min = end - start + 1
				l = start
				r = end
			}
			windowMap[s[start]]--
			if charMap[s[start]] != 0 && windowMap[s[start]] < charMap[s[start]] {
				charCnt--
			}
			start++
		}
		end++
	}
	if min == len(s)+1 {
		return ""
	}
	return s[l : r+1]
}
