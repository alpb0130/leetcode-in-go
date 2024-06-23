package morristraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	ptr *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var ptr *TreeNode
	cur := root
	for cur != nil && cur.Left != nil {
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
			cur = cur.Right
		}
	}
	ptr = cur
	return BSTIterator{ptr}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	res := this.ptr
	cur := this.ptr.Right
	for cur != nil && cur.Left != nil {
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
			cur = cur.Right
		}
	}
	this.ptr = cur
	return res.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.ptr != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
