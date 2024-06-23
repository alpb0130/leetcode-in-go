package dijkstra

import (
	"container/heap"
	"math"
)

// Dijkstra+Priority Queue. General Dijkstra will have O(V^2) time, because for each node,
// we need to go over all node to find the node with smallest dist and is not visited yet.
// We can update the process to find smallest dist using priority queue. This will give us
// O(VlogV) time.
// Dijkstra is very similar to BFS because BFS is solving all edges weight are 1 but Dijkstra
// is solving edge with different weight.
// Time: O(m*nlog(m*n))
// Space: O(m*n)
func shortestDistance(maze [][]int, start []int, destination []int) int {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return -1
	}
	distance := make([][]int, len(maze))
	for i := range distance {
		distance[i] = make([]int, len(maze[0]))
		for j := range distance[i] {
			distance[i][j] = math.MaxInt32
		}
	}
	distance[start[0]][start[1]] = 0
	// dijkstra and update the distance
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dijkstra(maze, start, distance, dir)

	if distance[destination[0]][destination[1]] == math.MaxInt32 {
		return -1
	}
	return distance[destination[0]][destination[1]]
}

func dijkstra(maze [][]int, start []int, dist [][]int, dir [][]int) {
	isVisited := make([][]bool, len(maze))
	for i := range isVisited {
		isVisited[i] = make([]bool, len(maze[0]))
	}
	pq := &priorityQueue{Pos{start[0], start[1], 0}}
	for pq.Len() > 0 {
		pos := heap.Pop(pq).(Pos)
		i, j, distance := pos.i, pos.j, pos.dist
		isVisited[i][j] = true
		if dist[i][j] < distance {
			continue
		}
		for _, d := range dir {
			k, l := i, j
			step := 0
			for k+d[0] >= 0 && k+d[0] < len(maze) && l+d[1] >= 0 && l+d[1] < len(maze[0]) && maze[k+d[0]][l+d[1]] != 1 {
				k += d[0]
				l += d[1]
				step++
			}
			if (k != i || l != j) && !isVisited[k][l] {
				if distance+step < dist[k][l] {
					dist[k][l] = distance + step
					heap.Push(pq, Pos{k, l, distance + step})
				}
			}
		}
	}

}

type Pos struct {
	i    int
	j    int
	dist int
}

type priorityQueue []Pos

func (q priorityQueue) Len() int {
	return len(q)
}
func (q priorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q priorityQueue) Less(i, j int) bool {
	return q[i].dist < q[j].dist
}
func (q *priorityQueue) Push(i interface{}) {
	*q = append(*q, i.(Pos))
}
func (q *priorityQueue) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}
