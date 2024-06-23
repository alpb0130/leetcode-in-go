package recursive

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(max(m, n)) - m, n is the lenght of l1 and l2
// Space: O(max(m, n))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addHelper(l1, l2, 0)
}

func addHelper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	if l1 == nil && l2 == nil {
		return &ListNode{Val: carry}
	}
	sum := 0
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}
	sum += carry
	next := addHelper(l1, l2, sum/10)
	node := &ListNode{Val: sum % 10, Next: next}
	return node
}
