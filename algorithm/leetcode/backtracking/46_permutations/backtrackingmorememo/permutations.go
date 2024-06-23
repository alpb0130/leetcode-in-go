package backtrackingmorememo

// Construct the permutations recursively. For a sublist [a0, a1,..., ak], the permutations of
// [a0, a1,..., ak, ak+1] is [ak+1, a0,..., ak], [a0, ak+1, a1,..., ak], ...,
// [a0, a1,..., ak, ak+1].
// Time: O(n!) ~ O(n*n!)
// Space: O(len(nums)*len(nums)). The recursive stack is of height len(nums) and for each call
//        we need len(nums) extra space to store tmp list.
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	permuteHelper(nums, []int{}, &res)
	return res
}

func permuteHelper(nums []int, list []int, res *[][]int) {
	if len(nums) == 0 {
		des := make([]int, len(list))
		copy(des, list)
		*res = append(*res, des)
		return
	}
	for i := 0; i <= len(list); i++ {
		des := make([]int, len(list[:i]))
		copy(des, list[:i])
		des = append(append(des[:i], nums[0]), list[i:]...)
		permuteHelper(nums[1:], des, res)
	}
}
