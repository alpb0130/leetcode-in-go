package priorityqueue

import "container/heap"

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

// Use priority queue to sort
// Time: O(n*logk) - n is the number of node and k is the number of linked list
// Space: O(k)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := ListNode{}
	ptr := &dummy
	// init heap
	h := &Nodes{}
	heap.Init(h)
	// push the first node of each list to
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}
	for len(*h) > 0 {
		next := heap.Pop(h).(*ListNode)
		ptr.Next = next
		ptr = ptr.Next
		if next.Next != nil {
			heap.Push(h, next.Next)
		}
	}
	return dummy.Next
}

type Nodes []*ListNode

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Less(i, j int) bool { return n[i].Val < n[j].Val }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n *Nodes) Pop() interface{} {
	old := *n
	l := len(old)
	i := old[l-1]
	*n = old[:l-1]
	return i
}
func (n *Nodes) Push(i interface{}) {
	*n = append(*n, i.(*ListNode))
}
