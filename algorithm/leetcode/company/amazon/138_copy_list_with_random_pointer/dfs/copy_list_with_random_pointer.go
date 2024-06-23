package dfs

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{val, nil, nil}
}

func CopyList(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	return copyHelper(node, map[*ListNode]*ListNode{})
}

func copyHelper(node *ListNode, nodeMap map[*ListNode]*ListNode) *ListNode {
	if node == nil {
		return nil
	}
	newNode, ok := nodeMap[node]
	if ok {
		return newNode
	}
	newNode = &ListNode{node.Val, nil, nil}
	nodeMap[node] = newNode
	newNode.Next = copyHelper(node.Next, nodeMap)
	newNode.Random = copyHelper(node.Random, nodeMap)
	return newNode
}

func Encode(node *ListNode) string {
	var str strings.Builder
	for node != nil {
		str.WriteString("[ Val: " + strconv.Itoa(node.Val) + ", ")
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
