package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// The challenge here is to define a recursive inline function (closure)
// Time: O(n)
// Space: O(h)
func recoverTree(root *TreeNode) {
	var pre, first, second *TreeNode
	var inOrderTraverse func(root *TreeNode)
	inOrderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrderTraverse(root.Left)
		if pre != nil && root.Val <= pre.Val {
			if first == nil {
				first = pre
			}
			second = root
		}
		pre = root
		inOrderTraverse(root.Right)
	}
	inOrderTraverse(root)
	first.Val, second.Val = second.Val, first.Val
}
