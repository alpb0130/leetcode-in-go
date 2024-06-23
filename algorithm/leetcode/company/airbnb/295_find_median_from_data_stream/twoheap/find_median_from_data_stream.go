package twoheap

import "container/heap"

type MedianFinder struct {
	min *minHeap
	max *maxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	min := &minHeap{}
	max := &maxHeap{}
	heap.Init(min)
	heap.Init(max)
	return MedianFinder{
		min: min,
		max: max,
	}
}
// Two steps:
// 1 - add number to the correct heap and make sure the max heap contains the small half and
//     the min heap contains the large half. The way to do that is utilize the heap sorting.
//     Add to max heap first and push the max of max heap to min heap.
// 2 - rebalance heap. Make sure len of max heap is no less than len of min heap. Equal or larger.
// **** can be improve by use both max heap. we can add -num into another queue for ranking.
// Time: O(logn)
// Space: O(n)
func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.max, num)
	heap.Push(this.min, heap.Pop(this.max).(int))
	if this.max.Len() < this.min.Len() {
		heap.Push(this.max, heap.Pop(this.min).(int))
	}
}

// Time: O(1)
// Space: O(n)
func (this *MedianFinder) FindMedian() float64 {
	if this.max.Len() > this.min.Len() {
		return float64(this.max.Peek().(int))
	} else {
		return float64(this.max.Peek().(int)+this.min.Peek().(int)) / 2
	}
}

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}
func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *minHeap) Pop() interface{} {
	x := (*m)[len(*m)-1]
	*m = (*m)[0 : len(*m)-1]
	return x
}
func (m *minHeap) Push(i interface{}) {
	*m = append(*m, i.(int))
}
func (m *minHeap) Peek() interface{} {
	return (*m)[0]
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}
func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *maxHeap) Pop() interface{} {
	x := (*m)[len(*m)-1]
	*m = (*m)[0 : len(*m)-1]
	return x
}
func (m *maxHeap) Push(i interface{}) {
	*m = append(*m, i.(int))
}
func (m *maxHeap) Peek() interface{} {
	return (*m)[0]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
