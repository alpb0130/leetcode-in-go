package dprecursive

// Solve the problem recursively. We need to support "." and "*". The hard part is to support "*".
// Three cases:
// 1. The second char of the pattern string is not "*":
// 	  look at the first char of s and first char of pattern and check whether they are matching
//    with each other. Then call isMatch(s[1:], p[1:]) recursively
// 2. The second char of the pattern string is "*":
//    2.1 If first char of s match with first char of pattern, call isMatch(s[1:], p)
//    2.2 If not, call isMatch(s, p[2:])
// Instead of directly calling the function recursively, we use a cache to store all computed results
// from existing call stack. Those results can be used by other call.
// Reference: https://leetcode.com/problems/regular-expression-matching/solution/
// Time: O(len(s) * len(p))
// Space: O(len(s) * len(p))
func isMatch(s string, p string) bool {
	dp := map[int]map[int]bool{}
	return isMatchHelper(s, p, 0, 0, dp)
}

func isMatchHelper(s string, p string, i, j int, dp map[int]map[int]bool) bool {
	if _, ok := dp[i][j]; ok {
		return dp[i][j]
	}
	if _, ok := dp[i]; !ok {
		dp[i] = map[int]bool{}
	}
	if j == len(p) {
		dp[i][j] = len(s) == i
		return dp[i][j]
	}
	matchFirst := i < len(s) && (s[i] == p[j] || p[j] == '.')
	if j+1 < len(p) && p[j+1] == '*' {
		// The result from first recursive call can be used by the second recursive call
		dp[i][j] = (matchFirst && isMatchHelper(s, p, i+1, j, dp)) ||
			isMatchHelper(s, p, i, j+2, dp)
	} else {
		dp[i][j] = matchFirst && isMatchHelper(s, p, i+1, j+1, dp)
	}
	return dp[i][j]
}
