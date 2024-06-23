package sorted_circular_linked_list

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Append(val int) {
	if l.Head == nil {
		l.Head = &Node{val, nil}
		l.Head.Next = l.Head
		l.Tail = l.Head
		return
	}
	head := l.Head
	tail := l.Tail
	ptr := l.Head
	if l.Head.Val > val {
		l.Head = &Node{val, nil}
		l.Head.Next = head
		tail.Next = l.Head
		return
	}
	for ptr.Next != head && ptr.Next.Val <= val {
		ptr = ptr.Next
	}
	node := &Node{val, nil}
	node.Next = ptr.Next
	ptr.Next = node
	if l.Tail == ptr {
		l.Tail = node
	}
}

func (l *LinkedList) Show() {
	if l.Head == nil {
		return
	}
	ptr := l.Head
	start := true
	for start || ptr != l.Head {
		fmt.Printf("%d ",ptr.Val)
		ptr = ptr.Next
		start = false
	}
	fmt.Println()
}
