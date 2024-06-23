package recursive

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive.
// Time: O(n)
// Space: O(n)
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	tail := head.Next
	tail.Next = head
	head.Next = nil
	return newHead
}
