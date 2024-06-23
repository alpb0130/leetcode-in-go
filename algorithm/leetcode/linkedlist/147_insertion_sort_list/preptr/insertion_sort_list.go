package preptr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// Use pre pointer to iterate
// Time: O(n^2)
// Space: O(1)
func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	if dummyHead.Next == nil {
		return dummyHead.Next
	}
	prePtr := dummyHead.Next
	for prePtr.Next != nil {
		ptr := dummyHead
		for ptr != prePtr && ptr.Next.Val < prePtr.Next.Val {
			ptr = ptr.Next
		}
		if ptr != prePtr {
			smallPtr := prePtr.Next
			prePtr.Next = smallPtr.Next
			smallPtr.Next = ptr.Next
			ptr.Next = smallPtr
		} else {
			prePtr = prePtr.Next
		}
	}
	return dummyHead.Next
}
