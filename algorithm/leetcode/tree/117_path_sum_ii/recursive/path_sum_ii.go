package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n*logn) - we need to copy every path to result list which is O(logn) and the worst case
// is each leaves will cause a result path and the time complexity will be O(nlogn)
// Space: O(h)
// Reference: https://stackoverflow.com/questions/24601111/whats-time-complexity-of-this-algorithm-for-finding-all-path-sum
func pathSum(root *TreeNode, sum int) [][]int {
	res := &[][]int{}
	if root == nil {
		return *res
	}
	pathSumHelper(root, sum, []int{}, res)
	return *res
}

func pathSumHelper(root *TreeNode, sum int, path []int, res *[][]int) {
	sum -= root.Val
	path = append(path, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*res = append(*res, pathCopy)
	}
	if root.Left != nil {
		pathSumHelper(root.Left, sum, path, res)
	}
	if root.Right != nil {
		pathSumHelper(root.Right, sum, path, res)
	}

}
