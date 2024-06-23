package curptr

type ListNode struct {
	Val  int
	Next *ListNode
}

// Use pre pointer to iterate
// Time: O(n^2)
// Space: O(1)
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := head
	for cur != nil {
		ptr := dummy
		next := cur.Next
		for ptr.Next != nil && ptr.Next.Val < cur.Val {
			ptr = ptr.Next
		}
		cur.Next = ptr.Next
		ptr.Next = cur
		cur = next
	}
	return dummy.Next
}
