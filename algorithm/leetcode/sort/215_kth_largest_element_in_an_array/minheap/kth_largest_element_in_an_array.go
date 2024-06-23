package minheap

import "container/heap"

// Use max heap. Push all element into heap and then pop until kth element
// Time: O(nlogk)
// Space: O(n)
func findKthLargest(nums []int, k int) int {
	min := &Nums{}
	heap.Init(min)
	for _, num := range nums {
		heap.Push(min, num)
		if len(*min) > k {
			heap.Pop(min)
		}
	}
	return heap.Pop(min).(int)
}

type Nums []int

func (n Nums) Len() int           { return len(n) }
func (n Nums) Less(i, j int) bool { return n[i] < n[j] }
func (n Nums) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n *Nums) Pop() interface{} {
	old := *n
	l := len(old)
	i := old[l-1]
	*n = old[:l-1]
	return i
}

func (n *Nums) Push(i interface{}) {
	*n = append(*n, i.(int))
}
