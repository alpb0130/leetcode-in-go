package inplace

type ListNode struct {
	Val  int
	Next *ListNode
}

// Two pointer to find the second part of the list and reverse second part of the list.
// Then do comparison.
// Time: O(n)
// Space: O(1)
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	slow = reverseList(slow)
	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true

}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	newHead := reverseList(next)
	next.Next = head
	head.Next = nil
	return newHead
}
