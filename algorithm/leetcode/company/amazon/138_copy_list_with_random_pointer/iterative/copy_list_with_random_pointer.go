package iterative

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

	nodeMap := map[*ListNode]*ListNode{}
	head := node
	newHead := getOrCloneNode(head, nodeMap)
	newNode := newHead
	for node != nil {
		newNode.Next = getOrCloneNode(node.Next, nodeMap)
		newNode.Random = getOrCloneNode(node.Random, nodeMap)
		node = node.Next
		newNode = newNode.Next
	}
	return newHead

}

func getOrCloneNode(node *ListNode, nodeMap map[*ListNode]*ListNode) *ListNode {
	if node == nil {
		return nil
	}
	newNode, ok := nodeMap[node]
	if !ok {
		newNode = &ListNode{node.Val, true, nil, nil}
		nodeMap[node] = newNode
	}
	return newNode
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
