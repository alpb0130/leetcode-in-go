package leftandrightscan

import "math"

// Scan from left to right and from right to left. We can get some idea from the
// problem longest valid parentheses sub sequence.
// When we scan from left to right, use a variable to note down how many '(' we have seen.
// If '(', left++, else left--. If left == 0, we know current substring must be a valid string
// and we can update current max length. In order to do that, we need a var to record the last
// invalid char (start with -1) and the valid length would be idx - idx_last_invalid_char.
// Once we see ')' and left is already be 0, we know currently the substring is not valid
// (like "())"). So we update left to be 0 and idx_last_invalid_char is current idx.
// But from left to right, we can only solve strings like "())". We can not solve "(()"
// correctly because left will never be 0 and we cannot update the result. So we need to
// use the same method to scan from right to left again.
// Time: O(n)
// Space: O(1)
func longestValidParentheses(s string) int {
	res := 0
	left := 0
	preIdx := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			if left > 0 {
				left--
				if left == 0 {
					res = maxInt(res, i-preIdx)
				}
			} else {
				preIdx = i
				left = 0
			}
		}
	}
	right := 0
	preIdx = len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			if right > 0 {
				right--
				if right == 0 {
					res = maxInt(res, preIdx-i)
				}
			} else {
				preIdx = i
				right = 0
			}
		}
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
