package stackandlessmemo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Apply two pointer on BST but we use stack to track the position. We can better a better
// memory usage which is the height of the tree.
// Time: O(n)
// Space: O(h)
func findTarget(root *TreeNode, k int) bool {
	left := &stack{}
	right := &stack{}
	pushAll(left, root, true)
	pushAll(right, root, false)
	for left.Len() > 0 && right.Len() > 0 && left.Peek() != right.Peek() {
		l := left.Peek()
		r := right.Peek()
		sum := l.Val + r.Val
		if sum == k {
			return true
		} else if sum > k {
			pushAll(right, right.Pop().Left, false)
		} else {
			pushAll(left, left.Pop().Right, true)
		}
	}
	return false
}

func pushAll(s *stack, root *TreeNode, isLeft bool) {
	if isLeft {
		for root != nil {
			s.Push(root)
			root = root.Left
		}
	} else {
		for root != nil {
			s.Push(root)
			root = root.Right
		}
	}
}

type stack []*TreeNode

func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) Push(n *TreeNode) {
	*s = append(*s, n)
}

func (s *stack) Pop() *TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

func (s *stack) Peek() *TreeNode {
	return (*s)[len(*s)-1]
}
