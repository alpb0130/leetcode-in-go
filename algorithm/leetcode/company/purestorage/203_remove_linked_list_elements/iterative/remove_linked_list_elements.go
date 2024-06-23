package iterative

type ListNode struct {
	Val  int
	Next *ListNode
}

// Use dummy head
// Time: O(n)
// Space: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	ptr := dummy
	for ptr.Next != nil {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
		} else {
			ptr = ptr.Next
		}
	}
	return dummy.Next
}
