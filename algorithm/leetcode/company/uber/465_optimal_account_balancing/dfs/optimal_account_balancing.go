package dfs

import (
	"math"
	"sort"
)

// The problem requires to find the min exchange solution so that nobody owns money.
// Sort and greedy won't work for the case [-5, -4, -1,  4, 3, 2, 1].
// We always can go the place where we get all final debts for every person. Our task
// is find the best way to do exchange. We can using dfs ans search. Basic idea is
// start with every position and try to transfer all it's debt to someone after him.
// We can try each option and move the the next position.
// Time: O(n!)
// Space: O(n)
func minTransfers(transactions [][]int) int {
	if len(transactions) <= 1 {
		return len(transactions)
	}
	debts := map[int]int{}
	for _, tran := range transactions {
		debts[tran[0]] += tran[2]
		debts[tran[1]] -= tran[2]
	}
	nonZero := []int{}
	for _, debt := range debts {
		if debt != 0 {
			nonZero = append(nonZero, debt)
		}
	}
	// sort so that we can skip
	sort.Ints(nonZero)
	// dfs
	return dfs(nonZero, 0)
}

func dfs(debts []int, idx int) int {
	for idx < len(debts) && debts[idx] == 0 {
		idx++
	}
	if idx == len(debts) {
		return 0
	}
	res := math.MaxInt32
	for i := idx + 1; i < len(debts); i++ {
		// Optimization. If the debt of two near by person is the skip,
		// no need to check this exchange way again
		if i > idx+1 && debts[i] == debts[i-1] {
			continue
		}
		// Only two person in different status can solve the debts
		if debts[i]*debts[idx] < 0 {
			debts[i] += debts[idx]
			res = minInt(res, dfs(debts, idx+1)+1)
			debts[i] -= debts[idx]
		}

	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
