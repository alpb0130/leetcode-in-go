package topdown

import "strings"

// Solve this problem recursively and in a top-down style.
// Time: O(3^N * 4^M) - N is the number of digits in the input that maps to 3 letters
//       and M is the number of digits in the input that maps to 4 letters
// Space: O(l^2) - the results set should not count as space complexity. For recursive solution,
//        the space should related to the tree height. l here is the length of digits and for each
//        call, we probably need 1, 2,..., l extra space. So the space complexity is l^2
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
	res := []string{}
	letterCombinationsHelper(digits, "", &res)
	return res
}

func letterCombinationsHelper(digits string, str string, res *[]string) {
	if len(digits) == 0 {
		(*res) = append(*res, str)
		return
	}
	for _, letter := range numLetterMap[digits[0]] {
		var s strings.Builder
		s.WriteString(str)
		s.WriteString(letter)
		letterCombinationsHelper(digits[1:], s.String(), res)
	}
}
