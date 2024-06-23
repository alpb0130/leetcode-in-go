package priorityqueueclean

import (
	"container/heap"
	"sort"
)

// Time: O(nlogn + n*n)
// Space: O(n)
// Simplify the point sorting by make the height of all start points to negative number.
// This also helps to differentiate start and endpoint.
// Also simplify the iterate process. Add or remove height first regardless the point type.
// If the prev height is different from cur height, add to results.
func getSkyline(buildings [][]int) [][]int {
	results := [][]int{}
	ps := points{}
	for _, b := range buildings {
		ps = append(ps, point{b[0], -b[2]})
		ps = append(ps, point{b[1], b[2]})
	}
	sort.Sort(ps)

	pq := &heights{0}
	heap.Init(pq)
	curMax := 0
	for _, p := range ps {
		if p.h < 0 {
			heap.Push(pq, -p.h)
		} else {
			i := 0
			for {
				if pq.get(i).(int) == p.h || i == pq.Len() {
					break
				}
				i++
			}
			heap.Remove(pq, i)
		}
		max := pq.Top().(int)
		if curMax != max {
			results = append(results, []int{p.x, max})
			curMax = max
		}
	}
	return results
}

type point struct {
	x int
	h int
}

type points []point

func (p points) Len() int { return len(p) }
func (p points) Less(i, j int) bool {
	if p[i].x == p[j].x {
		return p[i].h < p[j].h
	} else {
		return p[i].x < p[j].x
	}
}
func (p points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type heights []int

func (h heights) Len() int           { return len(h) }
func (h heights) Less(i, j int) bool { return h[i] > h[j] }
func (h heights) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *heights) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}
func (h *heights) Push(i interface{}) {
	*h = append(*h, i.(int))
}
func (h *heights) Top() interface{} {
	return (*h)[0]
}
func (h *heights) get(i int) interface{} {
	return (*h)[i]
}
