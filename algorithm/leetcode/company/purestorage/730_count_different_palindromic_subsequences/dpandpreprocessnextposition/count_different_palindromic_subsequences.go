package dpandpreprocessnextposition

// DP. But we use O(n) time to preprocess and get the position of leftmost and rightmost
// repeated letters for a letter in a specific position so that we don't need to do that
// in for loop.
// Time: O(n^2)
// Space: O(n^2)
func countPalindromicSubsequences(S string) int {
	if len(S) == 0 {
		return 0
	}
	dp := make([][]int, len(S))
	for i := range dp {
		dp[i] = make([]int, len(S))
		dp[i][i] = 1
	}
	charMap := map[byte]int{}
	rightNext := make([]int, len(S))
	for i := len(S) - 1; i >= 0; i-- {
		if idx, ok := charMap[S[i]]; ok {
			rightNext[i] = idx
		} else {
			rightNext[i] = len(S)
		}
		charMap[S[i]] = i
	}

	charMap = map[byte]int{}
	leftNext := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if idx, ok := charMap[S[i]]; ok {
			leftNext[i] = idx
		} else {
			leftNext[i] = -1
		}
		charMap[S[i]] = i
	}

	for i := 0; i < len(S); i++ {
		for j := i - 1; j >= 0; j-- {
			if S[i] != S[j] {
				dp[j][i] = dp[j][i-1] + dp[j+1][i] - dp[j+1][i-1]
			} else {
				left := rightNext[j]
				right := leftNext[i]
				if left > right {
					dp[j][i] = 2*dp[j+1][i-1] + 2
				}
				if left == right {
					dp[j][i] = 2*dp[j+1][i-1] + 1
				}
				if left < right {
					dp[j][i] = 2*dp[j+1][i-1] - dp[left+1][right-1]
				}
			}
			if dp[j][i] < 0 {
				dp[j][i] += 1000000007
			}
			dp[j][i] %= 1000000007
		}
	}
	return dp[0][len(S)-1]
}
