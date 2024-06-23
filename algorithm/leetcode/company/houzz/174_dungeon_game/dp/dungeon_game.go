package dp

import "math"

// DP. The problem requires us to find the min start blood needed for the game.
// We can just do dp from the right bottom to left top. Imagine if we want to
// survive the game till the end, we need at least 1 blood. If we already know
// the result for right bottom part, when we solve dp[i][j], we need to do
// Take min blood need for the current position which is:
// min(dp[i+1][j], dp[i][j+1])-dungeon[i][j]. If dungeon[i][j] is negative,
// it means we need more blood. Otherwise, it mean we need less blood but the
// blood should be at least 1.
// Time: O(m*n)
// Space: O(m*n)
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[m][n-1], dp[m-1][n] = 1, 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = maxInt(1, minInt(dp[i][j+1], dp[i+1][j])-dungeon[i][j])
		}
	}
	return dp[0][0]
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
