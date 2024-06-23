package iterative

type ListNode struct {
	Val  int
	Next *ListNode
}

// Iterative
// Time: O(n)
// Space: O(1)
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var prev *ListNode
	ptr := head
	for ptr != nil {
		next := ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = next
	}
	return prev
}
