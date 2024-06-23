package backtracking

import "sort"

// Refer to the preferred method of combination sum. We can solve this problem
// recursively. For target n, candidate C, the problem can be represented as F(n, C). So we can
// solve this problem by the formula:
// F(n, C) = F(n-c1, C) + F(n-c2, C-c1) + ... + F(n-ck-1, C-c1-c2-...-ck-2) + F(n-ck, C-c1-...-ck-1)
// where C = {c1, c2, c3, c4, ..., ck}
// Remember to skip the duplicate values. If we have a candidate list [c11, c12, c13]. Try to avoid
// duplication by the process:
// 1. If we use c11, we can use c12 and c13 in the recursive call.
// 2. If we don't use c11, we cannot use c12 and c13. This way we can avoid duplication.
// Time: ???. Really depends on the candidates set. Best case would be only candidate and it will
//       take O(target/c) time or target is less than any numbers in the candidates. Worst case
//       happened when target is pretty large and numbers in candidates are really small.
// Space: O(len(candidate)).
func combinationSum2(candidates []int, target int) [][]int {
	if target == 0 {
		return [][]int{}
	}
	res := [][]int{}
	// Try to identify potential duplication.
	sort.Ints(candidates)
	combinationSumHelper(candidates, target, []int{}, &res, 0)
	return res
}

func combinationSumHelper(candidates []int, target int, list []int, res *[][]int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		des := make([]int, len(list))
		copy(des, list)
		*res = append(*res, des)
		return
	}
	for i := start; i < len(candidates); i++ {
		// avoid duplication
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		combinationSumHelper(candidates, target-candidates[i], append(list, candidates[i]), res, i+1)
	}
}
