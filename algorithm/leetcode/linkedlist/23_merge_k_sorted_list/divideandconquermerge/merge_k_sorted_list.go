package divideandconquermerge

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
	if len(lists) == 0 {
		return nil
	}
	l := len(lists)
	for l > 1 {
		for i := 0; i < l; i += 2 {
			if i+1 < l {
				lists[i/2] = mergeTwoList(lists[i], lists[i+1])
			} else {
				lists[i/2] = lists[i]
			}
		}
		l = (l + 1) / 2
	}
	return lists[0]

}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	ptr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			ptr.Next = list2
			list2 = list2.Next
		} else {
			ptr.Next = list1
			list1 = list1.Next
		}
		ptr = ptr.Next
	}
	if list1 != nil {
		ptr.Next = list1
	} else if list2 != nil {
		ptr.Next = list2
	}
	return dummy.Next
}
