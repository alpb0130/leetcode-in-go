package iterative

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(m+n)
// Space: O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	ptr1, ptr2 := l1, l2
	dummy := ListNode{}
	prev := &dummy
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val <= ptr2.Val {
			prev.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			prev.Next = ptr2
			ptr2 = ptr2.Next
		}
		prev = prev.Next
	}
	if ptr1 != nil {
		prev.Next = ptr1
	} else {
		prev.Next = ptr2
	}
	return dummy.Next
}
