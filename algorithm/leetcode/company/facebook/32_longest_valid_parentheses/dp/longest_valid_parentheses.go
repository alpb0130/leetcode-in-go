package dp

// The brute force method is enum all substring add check whether they are valid.
// But there are duplications during the process. Eg. For string "()()[substring]",
// when we enum all possible substring, we would check [substring] multiple times
// for prefix "()" and "()()". Then we can memorize the process.
// For reference: https://leetcode.com/problems/longest-valid-parentheses/solution/
// In general, we know all valid substring will end with ")". We use dp[i] to record
// the longest valid parentheses substring ending with s[i]. For all s[j] == '(', the
// dp[j] will be 0 because it's must be invalid. Then there is two case if we see '('.
// If s[i-1] = ')', we need consider string right before dp[i-1]
// If s[i-1] = '(', we only need to consider dp[i-2]
// Time: O(n)
// Space: O(n)
func longestValidParentheses(s string) int {
	// dp[0] is the base case. If s == "", it will be 0
	dp := make([]int, len(s)+1)
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i+1] = dp[i-1] + 2
			} else {
				if i-dp[i]-1 >= 0 && s[i-dp[i]-1] == '(' {
					dp[i+1] = dp[i] + dp[i-dp[i]-1] + 2
				}
			}
		}
		if res < dp[i+1] {
			res = dp[i+1]
		}
	}
	return res
}
