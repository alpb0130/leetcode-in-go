package hasmapanddoublelinkedlist

type LRUCache struct {
	keyNodeMap map[int]*Node
	head       *Node
	tail       *Node
	capacity   int
}

// Use double linked list and hash map. Hash map maps to key-value node and record the length.
// double linked list would allow us to do all operation in O(1)
// All operation is O(1)
// Space: O(n)
func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		keyNodeMap: map[int]*Node{},
		head:       head,
		tail:       tail,
		capacity:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.keyNodeMap[key] == nil {
		return -1
	}
	node := this.keyNodeMap[key]
	this.detach(node)
	this.addFirst(node)
	return this.keyNodeMap[key].value
}

func (this *LRUCache) Put(key int, value int) {
	// insert/change
	node := this.keyNodeMap[key]
	if node == nil {
		node = &Node{
			key:   key,
			value: value,
		}
		this.keyNodeMap[key] = node
	} else {
		node.value = value
		this.detach(node)
	}
	this.addFirst(node)

	// kickoff is necessary
	if len(this.keyNodeMap) > this.capacity {
		removed := this.detach(this.tail.prev)
		delete(this.keyNodeMap, removed.key)
	}
}

func (this *LRUCache) detach(node *Node) *Node {
	next := node.next
	prev := node.prev
	prev.next = next
	next.prev = prev
	node.prev = nil
	node.next = nil
	return node
}

func (this *LRUCache) addFirst(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next = node
	node.next.prev = node
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
