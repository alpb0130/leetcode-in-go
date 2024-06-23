package backtracking

// Time: O(1) - the size of the problem is limited by 9 because duplication is not
// allowed.
// Space: O(k) - the stack size is also limit by k.
func combinationSum3(k int, n int) [][]int {
	if n < 6 {
		return [][]int{}
	}
	res := [][]int{}
	combinationSumHelper(k, n, 1, []int{}, &res)
	return res
}

func combinationSumHelper(k, n, start int, list []int, res *[][]int) {
	if n == 0 && k == 0 {
		des := make([]int, len(list))
		copy(des, list)
		*res = append(*res, des)
		return
	}
	if k == 0 || n < 0 || start > 9 {
		return
	}
	for i := start; i < 10; i++ {
		combinationSumHelper(k-1, n-i, i+1, append(list, i), res)
	}

}
