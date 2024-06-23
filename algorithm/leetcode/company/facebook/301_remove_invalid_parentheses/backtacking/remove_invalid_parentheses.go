package backtacking

// The problem requires to generate all valid parentheses string by removing min invalid
// parentheses.
// First, we get how many parentheses are invalid and this is the number of brackets we should
// remove.
// Then based on that, we apply back tracking to get all solution. Do not to put duplicated
// results to the final result. We could either dedupe during backtracking or use a map
// to store the final results which is much easier.
// Time: O(n+2^n) - n is the length of string. O(n) is to get how many invalid brackets.
//       O(2^n) is to get the answer. Basically each elements should have two choice: either
//       int the final sting not. So it should be 2^n. But because we apply some optimizations,
//       in reality, it should be less than 2^n
// Space: O(n) - dfs call.
func removeInvalidParentheses(s string) []string {
	left, right := countMisplaced(s)
	if left == 0 && right == 0 {
		return []string{s}
	}
	if left+right == len(s) {
		return []string{""}
	}
	res := []string{}
	constructRes(s, []byte{}, 0, left, right, 0, 0, &res)
	return res
}

func constructRes(s string, bytes []byte, start int, left int, right int, newLeft int, newRight int, res *[]string) {
	if newRight > 0 {
		return
	}
	if start == len(s) {
		if newLeft == 0 && newRight == 0 && left == 0 && right == 0 {
			dst := make([]byte, len(bytes))
			copy(dst, bytes)
			*res = append(*res, string(dst))
		}
		return
	}
	for i := start; i < len(s); i++ {
		if i == start || s[i] != s[i-1] {
			if s[i] == '(' && left > 0 {
				constructRes(s, bytes, i+1, left-1, right, newLeft, newRight, res)
			}
			if s[i] == ')' && right > 0 {
				constructRes(s, bytes, i+1, left, right-1, newLeft, newRight, res)
			}
		}
		if s[i] == '(' {
			newLeft++
		}
		if s[i] == ')' {
			if newLeft > 0 {
				newLeft--
			} else {
				newRight++
			}
		}
		bytes = append(bytes, s[i])
	}
	constructRes(s, bytes, len(s), left, right, newLeft, newRight, res)
}

func countMisplaced(s string) (int, int) {
	left, right := 0, 0
	for _, c := range s {
		if c == '(' {
			left++
		}
		if c == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	return left, right
}
