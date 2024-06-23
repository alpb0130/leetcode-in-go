package twostack

import (
	"math"

	"algorithm/leetcode/company/linkedin/models"
)

// The problem as us to find k closest nodes to target.
// The general method is get the sorted order array and find the pivotal place. Start from
// this place, use two pointer to get the first k closest node. Like [1,2,3,4,5,6] and 3.7,
// the pivotal place is 2 and 3, we can then use two pointer to go through the array.
// But the problem give us a BST and it's balanced. If it not a tree and we are given a sorted
// array we could solve it in o(logn+k) time using binary search and two pointer. Given the
// balanced BST, we can still use the same method. Instead of doing  binary search, we need two
// stack to store all the successors (larger than target) and predecessors (smaller than target).
// Find the two pivotal nodes and backtracking over the stack (two pointer process). The time
// is O(logn+k+logn). First logn is to initialize predecessor and successor. k+logn is the process
// to find k closest nodes. The loop will be executed k times and each time we pop out a value but
// finally there will be maximum 2logn node in stacks. And the stack start with logn values. So
// as we pop out k values, we also introduce logn push on the stack.
// Time: O(k+logn)
// Space: O(logn)
func closestKValues(root *models.TreeNode, target float64, k int) []int {
	res := []int{}
	if root == nil {
		return res
	}
	succ := &stack{}
	pred := &stack{}
	node := root
	// init succ and pred stack
	for node != nil {
		if float64(node.Val) >= target {
			succ.Push(node)
			node = node.Left
		} else {
			pred.Push(node)
			node = node.Right
		}
	}
	for len(res) < k {
		if succ.Len() > 0 && pred.Len() > 0 {
			successor := succ.Peek()
			predecessor := pred.Peek()
			succDiff := math.Abs(float64(successor.Val) - target)
			predDiff := math.Abs(float64(predecessor.Val) - target)
			if succDiff <= predDiff {
				res = append(res, successor.Val)
				getNextSucc(succ.Pop().Right, succ)
			} else {
				res = append(res, predecessor.Val)
				getNextPred(pred.Pop().Left, pred)
			}
		} else if succ.Len() > 0 {
			successor := succ.Pop()
			res = append(res, successor.Val)
			getNextSucc(successor.Right, succ)
		} else if pred.Len() > 0 {
			predecessor := pred.Pop()
			res = append(res, predecessor.Val)
			getNextPred(predecessor.Left, pred)
		} else {
			break
		}
	}
	return res
}

func getNextSucc(root *models.TreeNode, s *stack) {
	for root != nil {
		s.Push(root)
		root = root.Left
	}
}

func getNextPred(root *models.TreeNode, s *stack) {
	for root != nil {
		s.Push(root)
		root = root.Right
	}
}

type stack []*models.TreeNode

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Pop() *models.TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}
func (s *stack) Push(n *models.TreeNode) {
	*s = append(*s, n)
}
func (s *stack) Peek() *models.TreeNode {
	return (*s)[len(*s)-1]
}
