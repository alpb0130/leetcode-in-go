package backtrackinglessmemo

// Construct the permutations by swapping element in original array. Remember to
// swap back after the recursive call.
// F(n) = F(n-1, 1) + F(n-1, 2) + ...+ F(n-1, n-1). F(n-1, k) means swap 0 with k.
// Remember to use a map for each call to avoid duplication. If a number has been swapped
// then don't swap the number equal it again but in another position. For example, the current
// array is [1, 1, 2, 2] and start index is 0, then we only need to swap the first 1 itself and
// the first 1, the first 1, swap 1_1 with 1_2 will cause duplication, swap 1_1 with 2_2 will
// cause duplication as well.
// One caveat: https://leetcode.com/problems/permutations-ii/discuss/18595/Why-this-backtracking-algorithm-doesn't-work
// Time: O(n!) ~ O(n*n!) - https://leetcode.com/problems/permutations/solution/
// Space: O(len(nums)). The recursive stack is of height len(nums) and for each call
//        we need len(nums) extra space to store tmp list.
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	permuteUniqueHelper(nums, 0, &res)
	return res
}
func permuteUniqueHelper(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		des := make([]int, start)
		copy(des, nums)
		*res = append(*res, des)
		return
	}
	used := map[int]bool{}
	for i := start; i < len(nums); i++ {
		if !used[nums[i]] {
			used[nums[i]] = true
			nums[i], nums[start] = nums[start], nums[i]
			permuteUniqueHelper(nums, start+1, res)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
}
