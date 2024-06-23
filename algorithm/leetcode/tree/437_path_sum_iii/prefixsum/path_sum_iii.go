package prefixsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Utilize prefix sum. For problem related to range sum, always try to solve it using prefix sum.
// The benefit of prefix sum is that you don't need to compute the sum again for all range end with
// current value. You can always get a range sum by large_range_sum - small_range_sum. Generate DFS.
// Time: O(n)
// Space: O(h)
func pathSum(root *TreeNode, sum int) int {
	return pathSumHelper(root, 0, sum, map[int]int{0: 1})
}

func pathSumHelper(root *TreeNode, curSum, targetSum int, sumMap map[int]int) int {
	if root == nil {
		return 0
	}
	curSum += root.Val
	res := sumMap[curSum-targetSum]
	sumMap[curSum]++
	res += pathSumHelper(root.Left, curSum, targetSum, sumMap)
	res += pathSumHelper(root.Right, curSum, targetSum, sumMap)
	// Remove current prefix sum when backtracking
	sumMap[curSum]--
	return res
}
