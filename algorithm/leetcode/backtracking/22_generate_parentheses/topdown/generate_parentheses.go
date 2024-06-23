package topdown

import "strings"

// Solve this problem recursively. Instead of check all possible combinations and validate
// the string, we can prune if the sub candidate is not valid and backtrack. More specifically,
// keep build the string only if the number of left bracket are larger than right bracket.
// Time: O(n*Catalan(n)) = O(4^n / n^(1/2)) - there are totally Catalan(n) possible combinations
//       each need n time to complete
// Space: O(n*n) - height of the search tree is n the for every call, it would be 1, 2, 3,..., n
//        space.
func generateParenthesis(n int) []string {
	res := []string{}
	if n != 0 {
		generateParenthesisHelper("", n, n, &res)
	}
	return res
}

func generateParenthesisHelper(str string, leftNum, rightNum int, res *[]string) {
	if leftNum == 0 && rightNum == 0 {
		*res = append(*res, str)
	}
	if leftNum > 0 {
		var s strings.Builder
		s.WriteString(str)
		s.WriteString("(")
		generateParenthesisHelper(s.String(), leftNum-1, rightNum, res)
	}
	if rightNum > 0 && rightNum > leftNum {
		var s strings.Builder
		s.WriteString(str)
		s.WriteString(")")
		generateParenthesisHelper(s.String(), leftNum, rightNum-1, res)
	}
}
