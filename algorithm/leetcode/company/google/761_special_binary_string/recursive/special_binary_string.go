package recursive

import (
	"sort"
	"strings"
)

// Recursive. The problems requires to get the max special binary string by swapping
// two adjacent sub special binary strings any times. Special string is # of '1' are equal
// to number of '0' and in any prefix, # '1' > # '0'.
// Two things can think about:
// 1. Like valid parentheses (we want to find the max valid parentheses string by swapping)
// 2. If 1, up by 1, if 0, down by one. So we can get a maintain like graph.
//    1100 would be like:
//         /\
//        /  \
// We will know all special binary string will start with 0 y and end with 0 y and any point
// will not cross y=0. If we have:
//     /\
//  /\/  \
// We know a better string is:
//   /\
//  /  \/\
// So as we can see. What we want to do is always put the sub special string with larger height
// first.
// What if there are sub special string inside of a special binary string, like:
//       /\
//    /\/  \
//   /      \
// we cannot swap [0:3] with [3:] because 0, 3 is not special. We need to cut down the string to
// something like:
//       /\
//    /\/  \
// And then solve this problem.
// The basic idea is from top to down and we can solve the sub problems which are only consist
// special binary string.
// Reference: https://leetcode.com/problems/special-binary-string/solution/
// Time: O(n^2)
// Space: O(n^2)
func makeLargestSpecial(S string) string {
	if len(S) == 0 {
		return S
	}
	count := 0
	i := 0
	subStrings := []string{}
	for j := 0; j < len(S); j++ {
		if S[j] == '1' {
			count++
		} else {
			count--
		}
		if count == 0 {
			subStrings = append(subStrings, "1"+makeLargestSpecial(S[i+1:j])+"0")
			i = j + 1
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(subStrings)))
	return strings.Join(subStrings, "")
}
