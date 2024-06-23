package constantspace

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val    int
	IsNew  bool
	Next   *ListNode
	Random *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{val, false, nil, nil}
}

func CopyList(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	// create new node
	head := node
	for node != nil {
		next := node.Next
		node.Next = &ListNode{node.Val, true, nil, nil}
		node.Next.Next = next
		node = node.Next.Next
	}

	node = head
	// handle random node
	for node != nil {
		new := node.Next
		if node.Random != nil {
			new.Random = node.Random.Next
		}
		node = node.Next.Next
	}

	// split list
	node = head
	newHead := head.Next
	for node != nil {
		new := node.Next
		oldNext := new.Next
		node.Next = oldNext
		if oldNext != nil {
			new.Next = oldNext.Next
		}
		node = node.Next
		new = new.Next
	}
	return newHead
}

func Encode(node *ListNode) string {
	var str strings.Builder
	for node != nil {
		str.WriteString("[ Val: " + strconv.Itoa(node.Val) + ", ")
		str.WriteString("IsNew: " + strconv.FormatBool(node.IsNew) + ", ")
		str.WriteString("Next: ")
		if node.Next == nil {
			str.WriteString("nil")
		} else {
			str.WriteString(strconv.Itoa(node.Next.Val))
		}
		str.WriteString(", ")
		str.WriteString("Random: ")
		if node.Random == nil {
			str.WriteString("nil")
		} else {
			str.WriteString(strconv.Itoa(node.Random.Val))
		}
		str.WriteString("")
		str.WriteString("], ")
		node = node.Next
	}
	return str.String()
}
