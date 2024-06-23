package priorityqueue

import (
	"container/heap"
	"math"
)

// Use priority queue. Extend the two pointer method of I to 3D. Basically we want to put all
// boundaries into the priority queue and every time we pop out the bar with smallest height.
// Then we process its neighbors, if the height of a neighbor is smaller than current min height
// among boundary, we can trap some water. And after that we push neighbors into the boundary
// queue. Think about how to deal with teh current min!!!! The min height should be incremental!
// One way is every time we push the neighbor but modify its real height if it's low than current
// min height.
// Another way is we only update the min height when we pop a bar and it's higher than current min.
// We can only update the min height only if we are sure that the new boundary is higher than
// current height.
// In this method, we take way #1.
// Time: O(m*n*log(m+n)) - m*n is the size of the map.
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
	for queue.Len() > 0 {
		b := heap.Pop(queue).(bar)
		for _, d := range diff {
			row := b.x + d[0]
			col := b.y + d[1]
			if row >= 0 && row < m && col >= 0 && col < n && !isVisited[row][col] {
				isVisited[row][col] = true
				res += maxInt(0, b.h-heightMap[row][col])
				heap.Push(queue, bar{row, col, maxInt(b.h, heightMap[row][col])})
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

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
