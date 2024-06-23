package iterative

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(max(m, n)) - m, n is the lenght of l1 and l2
// Space: O(1)? O(max(m, n))?
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	ptr := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		ptr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		ptr = ptr.Next
	}
	if carry != 0 {
		ptr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
