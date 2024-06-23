package divideandconquermergerecursive

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

func mergeKLists(lists []*ListNode) *ListNode {
	return partition(lists, 0, len(lists)-1)
}

func partition(lists []*ListNode, i, j int) *ListNode {
	if i > j {
		return nil
	}
	if i == j {
		return lists[i]
	}
	mid := i + (j-i)/2
	return merge(partition(lists, i, mid), partition(lists, mid+1, j))
}

func merge(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = merge(list1.Next, list2)
		return list1
	} else {
		list2.Next = merge(list2.Next, list1)
		return list2
	}
}
