package priorityqueue

import "container/heap"

// Similar to to ugly number ii but use a priority queue to get current min value
// instead of iterate over the k list which would cause linear time.
// Time: O(nlogk) - k is the number of primes
// Space: O(k) - priority queue
func nthSuperUglyNumber(n int, primes []int) int {
	uglys := []int{1}
	minHeap := &numbers{}
	heap.Init(minHeap)
	for _, prime := range primes {
		heap.Push(minHeap, number{prime, 1, 0})
	}
	for i := 1; i < n; i++ {
		min := minHeap.Min().(number)
		uglys = append(uglys, min.value())
		for {
			curMin := minHeap.Min().(number)
			if curMin.value() != min.value() {
				break
			}
			curMin = heap.Pop(minHeap).(number)
			curMin.index++
			curMin.ugly = uglys[curMin.index]
			heap.Push(minHeap, curMin)
		}
	}
	return uglys[n-1]
}

type number struct {
	prime int
	ugly  int
	index int
}

func (n number) value() int {
	return n.prime * n.ugly
}

type numbers []number

func (n numbers) Len() int {
	return len(n)
}
func (n numbers) Less(i, j int) bool {
	return n[i].prime*n[i].ugly < n[j].prime*n[j].ugly
}
func (n numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n *numbers) Pop() interface{} {
	x := (*n)[len(*n)-1]
	(*n) = (*n)[0 : len(*n)-1]
	return x
}
func (n *numbers) Push(i interface{}) {
	(*n) = append(*n, i.(number))
}
func (n *numbers) Min() interface{} {
	return (*n)[0]
}
