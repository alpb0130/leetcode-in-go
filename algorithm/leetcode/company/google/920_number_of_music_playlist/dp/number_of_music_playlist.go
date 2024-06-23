package dp

import (
	"fmt"
	"math"
)

// 1. dp[0][0] is 0 because we can imagine we have 1 way to create a empty list
// 2. dp[i][j] == 0 where j > i because we cannot put j songs in a list of size i because we don't
//    have enough space
// 3. dp[i][0] == 0 because we cannot create a list with size i but no songs in it.
// Reference: https://leetcode.com/problems/number-of-music-playlists/solution/
//            https://leetcode.com/problems/number-of-music-playlists/discuss/180338/DP-solution-that-is-Easy-to-understand
// Time: O(N*L)
// Space: O(N*L) - can be improved to be O(N)
func numMusicPlaylists(N int, L int, K int) int {
	if N == 0 || L == 0 {
		return 0
	}
	dp := make([][]int, L+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	dp[0][0] = 1
	for i := 1; i < L+1; i++ {
		for j := 1; j < N+1; j++ {
			dp[i][j] = dp[i-1][j-1]*(N-j+1) + dp[i-1][j]*maxInt(0, j-K)
			dp[i][j] %= 1000000007
		}
	}
	fmt.Println(dp)
	return dp[L][N]
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
