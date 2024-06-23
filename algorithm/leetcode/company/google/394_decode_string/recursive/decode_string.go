package recursive

import (
	"strconv"
	"strings"
	"unicode"
)

// Solve this problem recursively.
// If we see digit, update number
// If we see string, update overall string in this call
// If we see left bracket, recursive call and write the result to overall string
// If we see right bracket, just return.
// Be careful about the index
// Time: O(n)
// Space: O(k) - k is depth of bracket
func decodeString(s string) string {
	i := 0
	return decodeHelper(s, &i)
}

func decodeHelper(s string, idx *int) string {
	if len(s) == (*idx) {
		return s
	}
	var res strings.Builder
	num := 0
	for (*idx) < len(s) {
		if unicode.IsNumber(rune(s[(*idx)])) {
			digit, _ := strconv.Atoi(string(s[(*idx)]))
			num = num*10 + digit
		} else if unicode.IsLetter(rune(s[(*idx)])) {
			res.WriteString(string(s[(*idx)]))
		} else if s[(*idx)] == '[' {
			(*idx)++
			subStr := decodeHelper(s, idx)
			for i := 0; i < num; i++ {
				res.WriteString(subStr)
			}
			num = 0
		} else {
			break
		}
		(*idx)++
	}
	return res.String()
}
