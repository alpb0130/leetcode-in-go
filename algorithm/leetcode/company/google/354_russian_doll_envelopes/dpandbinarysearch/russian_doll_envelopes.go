package dpandbinarysearch

import (
	"math"
	"sort"
)

// Comes from here: https://leetcode.com/problems/longest-increasing-subsequence/
// Time: O(nlogn)
// Space: O(n)
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Sort(Envelopes(envelopes))
	dp := []int{}
	for _, e := range envelopes {
		// binary search dp (in order) so we get the right place
		left, right := 0, len(dp)-1
		for left <= right {
			mid := (left + right) / 2
			if dp[mid] >= e[1] && (mid == 0 || dp[mid-1] < e[1]) {
				left = mid
				break
			} else if dp[mid] >= e[1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if left >= len(dp) {
			dp = append(dp, e[1])
		} else {
			dp[left] = e[1]
		}

	}
	return len(dp)
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
	// Be careful
	return e[i][1] > e[j][1]
}
