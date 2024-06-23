package bottomup

import "strings"

// Solve this problem recursively and in a bottom-up style.
// Time: O(3^N * 4^M) - N is the number of digits in the input that maps to 3 letters
//       and M is the number of digits in the input that maps to 4 letters
// Space: O(3^N * 4^M) - Different from top-down method. This algorithm need more space
//        because we return the temporary results to callee.
// Time and space complexity should be equal to the space of search tree.
var numLetterMap = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return letterCombinationsHelper(digits)
}

func letterCombinationsHelper(digits string) []string {
	if len(digits) == 0 {
		return []string{""}
	}
	res := []string{}
	for _, s := range letterCombinationsHelper(digits[1:]) {
		for _, str := range numLetterMap[digits[0]] {
			var builder strings.Builder
			builder.WriteString(str)
			builder.WriteString(s)
			res = append(res, builder.String())
		}
	}
	return res
}
