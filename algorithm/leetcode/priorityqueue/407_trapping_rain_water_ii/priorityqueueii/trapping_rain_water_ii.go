package priorityqueueii

import (
	"container/heap"
	"math"
)

// The same as another method but use a different way to set the current min boundary height.
// Time: O(m*n*log(m+n))
// Space: O(m+n)
func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	queue := &pq{}
	// init - add outermost bar into heap
	m := len(heightMap)
	n := len(heightMap[0])
	isVisited := make([][]bool, m)
	for i := 0; i < m; i++ {
		isVisited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		heap.Push(queue, bar{i, 0, heightMap[i][0]})
		isVisited[i][0] = true
		heap.Push(queue, bar{i, n - 1, heightMap[i][n-1]})
		isVisited[i][n-1] = true
	}
	for i := 1; i < n-1; i++ {
		heap.Push(queue, bar{0, i, heightMap[0][i]})
		isVisited[0][i] = true
		heap.Push(queue, bar{m - 1, i, heightMap[m-1][i]})
		isVisited[m-1][i] = true
	}
	res := 0
	diff := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	min := queue.Peek().(bar).h
	for queue.Len() > 0 {
		b := heap.Pop(queue).(bar)
		if b.h > min {
			min = b.h
		}
		for _, d := range diff {
			row := b.x + d[0]
			col := b.y + d[1]
			if row >= 0 && row < m && col >= 0 && col < n && !isVisited[row][col] {
				isVisited[row][col] = true
				res += maxInt(0, min-heightMap[row][col])
				heap.Push(queue, bar{row, col, heightMap[row][col]})
			}
		}
	}
	return res
}

type bar struct {
	x int
	y int
	h int
}

type pq []bar

func (q pq) Len() int {
	return len(q)
}
func (q pq) Less(i, j int) bool {
	return q[i].h < q[j].h
}
func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *pq) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}
func (q *pq) Push(i interface{}) {
	*q = append(*q, i.(bar))
}
func (q *pq) Peek() interface{} {
	return (*q)[0]
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
