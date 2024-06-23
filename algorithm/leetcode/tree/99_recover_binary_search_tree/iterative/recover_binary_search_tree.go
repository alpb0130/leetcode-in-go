package iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Inorder traversal. The only thing you need to take care of is how to find two swapped node.
// Look at the comment below.
// Time: O(n)
// Space: O(h)
func recoverTree(root *TreeNode) {
	cur := root
	s := &stack{}
	var pre, first, second *TreeNode
	for cur != nil || s.Len() > 0 {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		cur = s.Pop()
		// Very tricky here! This is how we find the swapped node. If two node are swapped,
		// there will be always at least one pair af nearby nodes that first.Val > second.Val.
		// The first time this happen, the first.Val is always the larger swapped node. But the
		// second.Val is probably be the smaller swapped node because two cases:
		// 1. 1, 2, 3, 4 => 1, 3, 2, 4 (one pair of unordered nodes: (3, 2))
		// 2. 1, 2, 3, 4 => 1, 4, 3, 2 (two pair of unordered nodes: (4, 3) and (3, 2) )
		if pre != nil && cur.Val <= pre.Val {
			if first == nil {
				first = pre
			}
			// The second unordered node will be update in case 2
			second = cur
		}
		pre = cur
		cur = cur.Right
	}
	first.Val, second.Val = second.Val, first.Val
}

type stack []*TreeNode

func (s *stack) Pop() *TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return n
}
func (s *stack) Push(n *TreeNode) {
	*s = append(*s, n)
}
func (s *stack) Len() int {
	return len(*s)
}
