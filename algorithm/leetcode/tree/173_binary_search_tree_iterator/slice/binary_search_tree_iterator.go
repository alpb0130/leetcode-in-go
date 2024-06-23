package slice

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


type BSTIterator struct {
	nodes []*TreeNode
	index int
}

// Preorder traverse the tree and put ordered result in a slice
// Time: O(n)
// Space: O(n)
func Constructor(root *TreeNode) BSTIterator {
	nodes := []*TreeNode{}
	cur := root
	for cur != nil {
		if cur.Left != nil {
			left := cur.Left
			for left.Right != nil {
				left = left.Right
			}
			left.Right = cur
			left = cur.Left
			cur.Left = nil
			cur = left
		} else {
			nodes = append(nodes, cur)
			cur = cur.Right
		}
	}
	return BSTIterator{nodes, 0}
}

// Time: O(1)
// Space: O(1)
/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	res := this.nodes[this.index]
	this.index++
	return res.Val
}

// Time: O(1)
// Space: O(1)
/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) > this.index
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
