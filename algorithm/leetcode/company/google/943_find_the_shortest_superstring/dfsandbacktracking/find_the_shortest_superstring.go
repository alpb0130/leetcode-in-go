package dfsandbacktracking

import "math"

// DFS and backtracking. Basically we try all string combinations but do some
// preprocessing and pruning.
// 1. Get the overlapping matrix to record overlapping length for all string pairs
// 2. DFS all string combinations and update the best path and best length (only record path)
//    to avoid string combination in dfs.
// 3. Use path to recover string
// Time: O(n!)
// Space: o(n^2) - overlapping map
func shortestSuperstring(A []string) string {
	if len(A) == 0 {
		return ""
	}
	if len(A) == 1 {
		return A[0]
	}
	// precompute overlapping
	overlapping := make([][]int, len(A))
	for i := range overlapping {
		overlapping[i] = make([]int, len(A))
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if i != j {
				for k := 1; k < minInt(len(A[i]), len(A[j])); k++ {
					if A[i][len(A[i])-k:] == A[j][:k] {
						overlapping[i][j] = k
					}
				}
			}
		}
	}
	// dfs to get the min super string
	bestPath := make([]int, len(A))
	bestLen := math.MaxInt32
	dfs(A, overlapping, 0, []int{}, make([]bool, len(A)), &bestPath, &bestLen)

	res := A[bestPath[0]]
	for i := 1; i < len(bestPath); i++ {
		prev := bestPath[i-1]
		cur := bestPath[i]
		res += A[bestPath[i]][overlapping[prev][cur]:]
	}
	return res
}

func dfs(A []string, overlapping [][]int, curLen int, path []int, isVisited []bool, bestPath *[]int, bestLen *int) {
	if curLen >= *bestLen {
		return
	}
	if len(path) == len(A) {
		*bestLen = curLen
		copy(*bestPath, path)
		return
	}
	for i := 0; i < len(A); i++ {
		if isVisited[i] {
			continue
		}
		isVisited[i] = true
		var incLen int
		if len(path) == 0 {
			incLen = len(A[i])
		} else {
			last := path[len(path)-1]
			incLen = len(A[i]) - overlapping[last][i]
		}
		dfs(A, overlapping, curLen+incLen, append(path, i), isVisited, bestPath, bestLen)
		isVisited[i] = false
	}
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
