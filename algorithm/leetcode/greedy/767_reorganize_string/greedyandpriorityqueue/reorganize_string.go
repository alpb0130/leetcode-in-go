package greedyandpriorityqueue

import (
	"bytes"
	"container/heap"
)

type Counter struct {
	char rune
	cnt  int
}

// Time: O(n) - should be O(n+n*logA) more specifically. n is the length of string. A is the # of
//				unique char. Given that A is a constant, the algorithm could be linear.
// Space: O(n)
func reorganizeString(S string) string {
	counter := map[rune]int{}
	max := 0
	for _, s := range S {
		counter[s]++
		if counter[s] > max {
			max = counter[s]
		}
	}
	if max > (len(S)+1)/2 {
		return ""
	}
	// build heap
	q := &pq{}
	heap.Init(q)
	for char, cnt := range counter {
		heap.Push(q, Counter{char, cnt})
	}
	// greedy method to rearrange string
	var str bytes.Buffer
	for !q.isEmpty() {
		popList := []Counter{}
		for i := 0; i < 2 && !q.isEmpty(); i++ {
			c := heap.Pop(q).(Counter)
			str.WriteString(string(c.char))
			if c.cnt-1 > 0 {
				popList = append(popList, Counter{c.char, c.cnt - 1})
			}
		}
		for _, c := range popList {
			heap.Push(q, c)
		}
	}
	return str.String()
}

type pq []Counter

func (q pq) Len() int           { return len(q) }
func (q pq) Less(i, j int) bool { return q[i].cnt > q[j].cnt }
func (q pq) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *pq) Pop() interface{} {
	old := *q
	result := old[len(old)-1]
	*q = old[:len(old)-1]
	return result
}

func (q *pq) Push(i interface{}) {
	*q = append(*q, i.(Counter))
}

func (q *pq) isEmpty() bool {
	return len(*q) == 0
}
