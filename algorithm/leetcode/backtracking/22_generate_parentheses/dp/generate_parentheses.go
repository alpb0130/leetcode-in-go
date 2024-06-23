package dp

import "strings"

// Solve this problem using dp. For F(n), it can be construct from F(0), F(1), F(2),..., F(n-1) by
// F(n) = (-F(0)-)-F(n-1), (-F(1)-)-F(n-2),..., (-F(n-1)-)-F(0). This should construct all
// solutions without repetition.
// Time: O(n*Catalan(n)) = O(4^n / n^(1/2)) - there are totally Catalan(n) possible combinations
//       each need n time to complete
// Space: O(n*n) - height of the search tree is n the for every call, it would be 1, 2, 3,..., n
//        space.
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, s1 := range dp[j] {
				for _, s2 := range dp[i-j-1] {
					var s strings.Builder
					s.WriteString("(")
					s.WriteString(s1)
					s.WriteString(")")
					s.WriteString(s2)
					dp[i] = append(dp[i], s.String())
				}
			}

		}
	}
	return dp[n]
}
