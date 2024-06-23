package priorityqueue

import "container/heap"

// Use min heap to keep top k frequent element. Need to reverse the results in the final for loop.
// Time: O(nlongk)
// Space: O(max(x, k))
func topKFrequent(nums []int, k int) []int {
	results := make([]int, k)
	if k == 0 {
		return results
	}
	cnts := map[int]int{}
	for _, num := range nums {
		cnts[num]++
	}
	c := &counters{}
	heap.Init(c)
	for num, cnt := range cnts {
		heap.Push(c, counter{num, cnt})
		if c.Len() > k {
			heap.Pop(c)
		}
	}
	for i := k; i > 0; i-- {
		results[i-1] = heap.Pop(c).(counter).num
	}

	return results
}

type counter struct {
	num int
	cnt int
}
type counters []counter

func (c counters) Len() int           { return len(c) }
func (c counters) Less(i, j int) bool { return c[i].cnt < c[j].cnt }
func (c counters) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c *counters) Pop() interface{} {
	x := (*c)[len(*c)-1]
	*c = (*c)[0 : len(*c)-1]
	return x
}
func (c *counters) Push(i interface{}) {
	*c = append(*c, i.(counter))
}
