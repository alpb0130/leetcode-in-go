package greedyandpriorityqueue

import "container/heap"

// Greedy method with priority queue - instead of sorting, we use priority queue to sort
// Time: O(times) - sorting takes O(1) time because the max len of the list is 26...
// Space: O(1) - map with const size
func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}
	taskMap := map[byte]int{}
	for _, task := range tasks {
		taskMap[task] += 1
	}
	frequency := IntHeap{}
	for _, freq := range taskMap {
		frequency = append(frequency, freq)
	}
	freqPtr := &frequency
	heap.Init(freqPtr)
	time := 0
	for {
		if freqPtr.Len() == 0 {
			break
		}
		temp := []int{}
		for i := 0; i < n+1; i++ {
			if len(temp) == 0 && freqPtr.Len() == 0 {
				break
			}
			if freqPtr.Len() > 0 {
				max := heap.Pop(freqPtr).(int) - 1
				if max > 0 {
					temp = append(temp, max)
				}
			}
			time++
		}
		for _, freq := range temp {
			heap.Push(freqPtr, freq)
		}
	}
	return time
}

// Int max heap
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
