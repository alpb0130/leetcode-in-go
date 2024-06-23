package twoforloop

import "strings"

// Time: O(n^m)
// Space: O(n^m)
func validWordSquare(words []string) bool {
	if len(words) == 0 {
		return true
	}
	for i := 0; i < len(words); i++ {
		word := words[i]
		var colWord strings.Builder
		for j := 0; j < len(words) && i < len(words[j]); j++ {
			colWord.WriteString(string(words[j][i]))
		}
		if word != colWord.String() {
			return false
		}
	}
	return true
}
