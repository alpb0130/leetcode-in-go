package morris_traverse

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Inorder morris traverse.
// A better idea is if we can store the num of node of a tree in the node, we can do that in
// O(logn) time because we can do binary search.
// Time: O(k)
// Space: O(1)
func kthSmallest(root *TreeNode, k int) int {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			pred := cur.Left
			for pred.Right != nil && pred.Right != cur {
				pred = pred.Right
			}
			if pred.Right == nil {
				pred.Right = cur
				cur = cur.Left
				continue
			} else {
				pred.Right = nil
			}
		}
		k--
		if k == 0 {
			return cur.Val
		}

		cur = cur.Right
	}
	return -1
}
