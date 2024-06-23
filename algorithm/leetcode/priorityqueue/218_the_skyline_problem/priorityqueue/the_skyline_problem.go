package priorityqueue

import (
	"container/heap"
	"sort"
)

// Time: O(nlogn + n*n) - n is the number of buildings. The first part is sorting. The second
// part is iterate over critical point and update current max height. Because we don't have
// red-black tree, the remove operation need O(n) time.
// Space: O(n)
func getSkyline(buildings [][]int) [][]int {
	results := [][]int{}
	ps := points{}
	// Sort critical points (both start and end points together by the x value)
	// If x is the same, break the tie by point type and height. There are three cases:
	// 1. If both point are start points, put higher point first so that when lower point is added,
	//    it's not highest point and no need to add to results
	// 2. If both points are end points, put lower point first so that when lower point is added,
	// 	  it's not highest point and no need to add to results
	// 3. If two points have different types, always put the start point first.
	for _, b := range buildings {
		ps = append(ps, point{b[0], b[2], true})
		ps = append(ps, point{b[1], b[2], false})
	}
	sort.Sort(ps)

	pq := &heights{0}
	heap.Init(pq)
	for _, p := range ps {
		// If it's start point, push to height pool and if it's higher than cur max height,
		// add to results.
		if p.isLi {
			if p.h > pq.Top().(int) {
				results = append(results, []int{p.x, p.h})
			}
			heap.Push(pq, p.h)
		} else {
			// If it's end point, remove the point from height pool and if it higher than
			// cur max height, add to results
			i := 0
			for {
				if pq.get(i).(int) == p.h || i == pq.Len() {
					break
				}
				i++
			}
			heap.Remove(pq, i)
			if p.h > pq.Top().(int) {
				results = append(results, []int{p.x, pq.Top().(int)})
			}
		}
	}
	return results
}

type point struct {
	x    int
	h    int
	isLi bool
}

type points []point

func (p points) Len() int { return len(p) }
func (p points) Less(i, j int) bool {
	if p[i].x == p[j].x {
		if p[i].isLi != p[j].isLi {
			return p[i].isLi
		} else {
			return (p[i].h < p[j].h) != p[i].isLi
		}
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
