package dp

import (
	"math"
	"sort"
)

// Typical DP solution.
// A better solution could be from the solution here:
// https://leetcode.com/problems/longest-increasing-subsequence/
// Time: O(n^2)
// Space: O(n)
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Sort(Envelopes(envelopes))
	dp := make([]int, len(envelopes))
	for i := range dp {
		dp[i] = 1
	}
	res := 0
	for i, e := range envelopes {
		for j := 0; j < i; j++ {
			if envelopes[j][0] < e[0] && envelopes[j][1] < e[1] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

type Envelopes [][]int

func (e Envelopes) Len() int {
	return len(e)
}
func (e Envelopes) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e Envelopes) Less(i, j int) bool {
	if e[i][0] != e[j][0] {
		return e[i][0] < e[j][0]
	}
	return e[i][1] < e[j][1]
}
