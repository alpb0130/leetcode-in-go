package _25_valid_palindrome

import (
	"strings"
	"unicode"
)

// Time: O(n)
// Space: O(1)
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
		} else if !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
		} else {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}
