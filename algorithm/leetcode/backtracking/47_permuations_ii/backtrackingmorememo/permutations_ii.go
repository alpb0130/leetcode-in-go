package backtrackingmorememo

// Construct the permutations recursively. For a sublist [a0, a1,..., ak], the permutations of
// [a0, a1,..., ak, ak+1] is [ak+1, a0,..., ak], [a0, ak+1, a1,..., ak], ...,
// [a0, a1,..., ak, ak+1]. We need to do avoid duplication in the problem because there are
// probably duplicate numbers in nums. The way we can avoid that is for permutations of
// [a0, a1,..., ak, ak+1], insert ak+1 from back of the list [a0, a1,..., ak] until the current
// value is equal to ak+1. For example, if we insert 2 to [1, 2, 1], [1, 2, 2_new, 1] is equal
// to [1, 2_new, 2, 1].
// Time: O(n!) ~ O(n*n!)
// Space: O(len(nums)*len(nums)). The recursive stack is of height len(nums) and for each call
//        we need len(nums) extra space to store tmp list.
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	permuteUniqueHelper(nums, []int{}, &res)
	return res
}
func permuteUniqueHelper(nums []int, list []int, res *[][]int) {
	if len(nums) == 0 {
		des := make([]int, len(list))
		copy(des, list)
		*res = append(*res, des)
		return
	}
	for i := len(list); i >= 0; i-- {
		if i < len(list) && nums[0] == list[i] {
			break
		}
		des := make([]int, len(list[:i]))
		copy(des, list[:i])
		des = append(append(des[:i], nums[0]), list[i:]...)
		permuteUniqueHelper(nums[1:], des, res)
	}
}
