package backtracking

// For each recursive call, remove the first element from candidate set so that we don't
// generate duplicate results.
// Time: ???. Really depends on the candidates set. Best case would be only candidate and it will
//       take O(target/c) time. Worst case happened when target is pretty large and numbers in
//       candidates are really small.
// Space: O(len(candidates)) + O(max(results)) - we would have at most len(candidates)
// recursive call at the same time. The extra space we use is the list but every call shares the
// same slice space.
func combinationSum(candidates []int, target int) [][]int {
	if target == 0 {
		return [][]int{}
	}
	res := [][]int{}
	combinationSumHelper(candidates, target, []int{}, &res)
	return res
}

func combinationSumHelper(candidates []int, target int, list []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*res = append(*res, tmp)
		return
	}
	if len(candidates) == 0 {
		return
	}
	for target >= 0 {
		combinationSumHelper(candidates[1:], target, list, res)
		list = append(list, candidates[0])
		target -= candidates[0]
	}
	return
}
