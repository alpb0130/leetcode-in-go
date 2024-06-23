package separatechaining

import (
	"hash/fnv"
	"strconv"
)

// A better solutioin: https://leetcode.com/problems/design-hashmap/discuss/205285/reproduce-hash()-and-resize()-in-Java-8-source-code
// including resizing and generate better hashing
// Use separate chaining to design the hash map
type MyHashMap struct {
	nodeList []*Node
}

type Node struct {
	key   int
	value int
	next  *Node
}

// Space: O(n)
/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		nodeList: make([]*Node, 10000),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	idx := this.getIndex(key)
	prev := this.findPrevNode(idx, key)
	if prev.next == nil {
		prev.next = &Node{
			key: key,
		}
	}
	prev.next.value = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	idx := this.getIndex(key)
	prev := this.findPrevNode(idx, key)
	if prev.next != nil {
		return prev.next.value
	} else {
		return -1
	}
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	idx := this.getIndex(key)
	prev := this.findPrevNode(idx, key)
	if prev.next != nil {
		prev.next = prev.next.next
	}
}

// Remember how to generate hash code
func (this *MyHashMap) getIndex(key int) int {
	h := fnv.New32()
	keyStr := strconv.Itoa(key)
	h.Write([]byte(keyStr))
	return int(h.Sum32()) % len(this.nodeList)
}

func (this *MyHashMap) findPrevNode(idx int, key int) *Node {
	list := this.nodeList[idx]
	if list == nil {
		n := Node{}
		this.nodeList[idx] = &n
		return &n
	}
	ptr := list
	for ptr.next != nil && ptr.next.key != key {
		ptr = ptr.next
	}
	return ptr
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
