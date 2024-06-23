package priorityqueue

import "container/heap"

// Time: O(nlogk) - using priority queue to sort.
// Space: O(k)
func kClosest(points [][]int, K int) [][]int {
	pq := &priorityQueue{}
	for _, point := range points {
		heap.Push(pq, point)
		if pq.Len() > K {
			heap.Pop(pq)
		}
	}
	return [][]int(*pq)
}

func calcDistance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

type priorityQueue [][]int

func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Less(i, j int) bool {
	return calcDistance(pq[i]) > calcDistance(pq[j])
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *priorityQueue) Pop() interface{} {
	i := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return i
}
func (pq *priorityQueue) Push(i interface{}) {
	*pq = append(*pq, i.([]int))
}
