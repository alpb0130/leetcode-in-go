package backtracking

import "strings"

// Get the first half of the palindrome string and the possible odd letter.
// Then permute over the first half string and construct the final result.
// Time: O(n+n*n!) - most of the time is on permutation which should be O(n*n!)
// Space: O(n) - n is the length of the string. The counter map is O(1) space.
//        Then most of the space is from recursive call to get all permutations.
//        The height of recursive tree is n/2 (O(n/2))
func generatePalindromes(s string) []string {
	counter := map[rune]int{}
	for _, c := range s {
		counter[c]++
	}
	halfStr := []rune{}
	oddChar := ""
	oddCnt := 0
	for c, cnt := range counter {
		for cnt > 1 {
			halfStr = append(halfStr, c)
			cnt -= 2
		}
		if cnt == 1 {
			oddChar = string(c)
			oddCnt++
		}
	}
	if oddCnt > 1 {
		return []string{}
	}
	res := []string{}
	permutationHelper(halfStr, oddChar, 0, &res)
	return res
}

func permutationHelper(str []rune, oddChar string, start int, res *[]string) {
	if start == len(str) {
		var s strings.Builder
		for _, c := range str {
			s.WriteString(string(c))
		}
		s.WriteString(oddChar)
		s.WriteString(reverseStr(str))
		*res = append(*res, s.String())
		return
	}
	used := map[rune]bool{}
	for i := start; i < len(str); i++ {
		if !used[str[i]] {
			used[str[i]] = true
			str[i], str[start] = str[start], str[i]
			permutationHelper(str, oddChar, start+1, res)
			str[i], str[start] = str[start], str[i]
		}
	}
}

func reverseStr(str []rune) string {
	var s strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		s.WriteString(string(str[i]))
	}
	return s.String()
}
