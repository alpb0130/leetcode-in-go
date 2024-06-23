package recursive

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive.
// Time: O(n)
// Space: O(n)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	}
	head.Next = removeElements(head.Next, val)
	return head
}
