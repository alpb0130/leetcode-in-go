package iterativereversepreorder1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Apply preorder traversal but push left node first. What post order want us to do is get
// left, right and then root. We can get the right-left preorder traversal which is root,
// right and then left and reverse the results. We can always prepend the node val to the
// result list to implement the reverse.
// Time: O(n)
// Space: O(h)
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	s := &stack{root}
	for s.Len() != 0 {
		node := s.Pop()
		res = append([]int{node.Val}, res...)
		if node.Left != nil {
			s.Push(node.Left)
		}
		if node.Right != nil {
			s.Push(node.Right)
		}
	}
	return res
}

type stack []*TreeNode

func (s *stack) Push(n *TreeNode) {
	*s = append(*s, n)
}

func (s *stack) Pop() *TreeNode {
	x := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s *stack) Len() int {
	return len(*s)
}
