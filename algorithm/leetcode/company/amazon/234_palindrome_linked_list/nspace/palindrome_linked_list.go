package nspace

type ListNode struct {
	Val  int
	Next *ListNode
}

// Convert the linked list to a number array and use two pointer.
// Time: O(n)
// Space: O(n)
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	nodes := []int{}
	node := head
	for node != nil {
		nodes = append(nodes, node.Val)
		node = node.Next
	}
	i, j := 0, len(nodes)-1
	for i < j {
		if nodes[i] != nodes[j] {
			return false
		}
		i++
		j--
	}
	return true
}
