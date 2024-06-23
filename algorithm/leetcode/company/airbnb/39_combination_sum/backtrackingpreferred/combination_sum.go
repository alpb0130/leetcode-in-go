package backtrackingpreferred

// Similar to another backtracking method but slightly different. We can solve this problem
// recursively. For target n, candidate C, the problem can be represented as F(n, C). So we can
// solve this problem by the formula:
// F(n, C) = F(n-c1, C) + F(n-c2, C-c1) + ... + F(n-ck-1, C-c1-c2-...-ck-2) + F(n-ck, C-c1-...-ck-1)
// where C = {c1, c2, c3, c4, ..., ck}
// Time: ???. Really depends on the candidates set. Best case would be only candidate and it will
//       take O(target/c) time or target is less than any numbers in the candidates. Worst case
//       happened when target is pretty large and numbers in candidates are really small and length
//       of candidates is pretty big.
// Space: O(target/min(candidates))
func combinationSum(candidates []int, target int) [][]int {
	if target == 0 {
		return [][]int{}
	}
	res := [][]int{}
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
		combinationSumHelper(candidates, target-candidates[i], append(list, candidates[i]), res, i)
	}
}
