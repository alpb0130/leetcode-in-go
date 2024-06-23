package backtrackinglessmemo

// Construct the permutations by swapping element in original array. Remember to
// swap back after the recursive call.
// F(n) = F(n-1, 1) + F(n-1, 2) + ...+ F(n-1, n-1). F(n-1, k) means swap 1 with k.
// Time: O(n!) ~ O(n*n!) - https://leetcode.com/problems/permutations/solution/
// Space: O(len(nums)). The recursive stack is of height len(nums) and for each call
//        we need len(nums) extra space to store tmp list.
func permute(nums []int) [][]int {
	res := [][]int{}
	permuteHelper(nums, 0, &res)
	return res
}

func permuteHelper(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		des := make([]int, len(nums))
		copy(des, nums)
		*res = append(*res, des)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		permuteHelper(nums, start+1, res)
		nums[i], nums[start] = nums[start], nums[i]
	}
}
