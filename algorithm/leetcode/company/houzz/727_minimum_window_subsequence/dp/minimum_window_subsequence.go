package dp

import "math"

// DP. Imagine we know all details of the solution to the sub-problem S[0:i], T[0:j]. What can we
// do for S[0:i+1], T[0:j+1]? If we know S[i+1] == T[j+1], we can get the answer from
// S[0:i], T[0:j]. If S[i+1] != T[j+1], we can get the answer from S[0:i], T[0:j+1]. This leads us
// to thinking about dynamic programming. What do we need to know about the the solution for
// the sub-problem? In order to get the substring length, we need to know the start matching index.
// Then we can get the substring length by doing i-dp[i][j]+. i is the end index of current S sub
// string.
// Reference: https://leetcode.com/problems/minimum-window-subsequence/discuss/109362/Java-Super-Easy-DP-Solution-(O(mn))
// Time: O(s*t)
// Space: O(s*t)
func minWindow(S string, T string) string {
	if len(S) == 0 || len(T) == 0 {
		return ""
	}
	dp := make([][]int, len(T))
	for i := range dp {
		dp[i] = make([]int, len(S))
		for j := 0; j < len(S); j++ {
			dp[i][j] = -1
		}
	}
	for i := 0; i < len(S); i++ {
		if S[i] == T[0] {
			dp[0][i] = i
		} else if i >= 1 {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < len(T); i++ {
		for j := 1; j < len(S); j++ {
			if S[j] == T[i] {
				dp[i][j] = maxInt(dp[i-1][j-1], dp[i][j-1])
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	var start, end int
	min := len(S) + 1
	for i := 0; i < len(S); i++ {

		if dp[len(T)-1][i] != -1 && i-dp[len(T)-1][i]+1 < min {
			min = i - dp[len(T)-1][i] + 1
			start, end = dp[len(T)-1][i], i
		}
	}
	if min == len(S)+1 {
		return ""
	}
	return S[start : end+1]
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
