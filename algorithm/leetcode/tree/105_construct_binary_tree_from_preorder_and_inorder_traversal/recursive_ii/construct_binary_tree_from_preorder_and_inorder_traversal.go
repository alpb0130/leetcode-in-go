package recursive_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// A global map to reduce the index search time
// Master theorem: T(n) = a * T(n/b) + f(n)
// T(n) = n^(log_b^a) + f (compare to n^(log_b^a))
// Here, f(n) = n. T(n) = nlogn.
// Time: O(n*logn) - master theorem.
// Space: O(h)
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := map[int]int{}
	for i, value := range inorder {
		inorderMap[value] = i
	}
	return buildTreeHelper(preorder, 0, len(preorder)-1, 0, len(inorder)-1, inorderMap)
}

func buildTreeHelper(preorder []int, prestart, preend, instart, inend int, inorderMap map[int]int) *TreeNode {
	if preend < prestart || inend < instart {
		return nil
	}
	root := &TreeNode{Val: preorder[prestart]}
	i := inorderMap[preorder[prestart]]
	leftLen := i - instart
	root.Left = buildTreeHelper(preorder, prestart+1, prestart+leftLen, instart, i-1, inorderMap)
	root.Right = buildTreeHelper(preorder, prestart+leftLen+1, preend, i+1, inend, inorderMap)
	return root
}
