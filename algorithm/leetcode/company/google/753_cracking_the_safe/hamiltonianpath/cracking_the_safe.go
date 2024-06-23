package hamiltonianpath

import (
	"math"
	"strings"
)

// Basically we know there are k^n different passwords and inorder to get the shortest string that
// contains all passwords, we need to concatenate all passwords together and try to make sure
// two nearby password overlapping with each other in length of n-1. So ideally, the length of
// shortest results are k^n+(n-1). We can do dfs and backtracking to solve the problem.
// Start with "0000" (assume n is 4) and we use the last three digit as the first 3 digits of next
// passwords and we do dfs on the last digits. Our purpose is to find the first string that can
// visit all passwords exactly once.
// Reference: https://www.youtube.com/watch?v=kRdlLahVZDc
// Time: O(k^(k^n)) best case: O(k^n)
// Space: O(k^n*(n+k^n-1)) - n+k^n-1 is length of string
func crackSafe(n int, k int) string {
	if n == 0 || k == 0 {
		return ""
	}
	numPass := int(math.Pow(float64(k), float64(n)))
	res := ""
	var start strings.Builder
	for i := 0; i < n; i++ {
		start.WriteString("0")
	}
	isVisited := map[string]bool{start.String(): true}
	// DFS
	dfs(start.String(), n, k, numPass, 1, isVisited, &res)
	return res
}

func dfs(cur string, n, k, numPass, curNumPass int, isVisited map[string]bool, res *string) bool {
	if numPass == curNumPass {
		*res = cur
		return true
	}
	prefix := cur[len(cur)-n+1:]
	for i := 0; i < k; i++ {
		number := string('0' + i)
		if !isVisited[prefix+number] {
			isVisited[prefix+number] = true
			if dfs(cur+number, n, k, numPass, curNumPass+1, isVisited, res) {
				return true
			}
			isVisited[prefix+number] = false
		}
	}
	return false
}
