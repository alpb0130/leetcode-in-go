package longestpath

// Use BFS twice to find the longest path and get the middle points
// Time: O(n) ~ O(V+E) but it's a tree-like graph so V=E=n
// Space: O(n)
func findMinHeightTrees(n int, edges [][]int) []int {
	graph := make([][]int, n)
	isMarked := make([]bool, n)
	edgeFrom := make([]int, n)
	dist := make([]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// Find the longest path using BFS
	furthestV := bfs(graph, isMarked, edgeFrom, dist, 0)
	// Rerun BFS but start with furthest vertex
	for i := 0; i < n; i++ {
		isMarked[i] = false
	}
	longestPathEnd := bfs(graph, isMarked, edgeFrom, dist, furthestV)
	// walk throught the longest path and get the vertex in the middle
	res := []int{}
	longPathLength := dist[longestPathEnd]
	for i := 0; i <= longPathLength; i++ {
		if i == longPathLength/2 || i == (longPathLength+1)/2 {
			res = append(res, longestPathEnd)
		}
		longestPathEnd = edgeFrom[longestPathEnd]
	}
	return res
}

func bfs(graph [][]int, isMarked []bool, edgeFrom []int, dist []int, start int) int {
	q := &queue{start}
	isMarked[start] = true
	edgeFrom[start] = -1
	dist[start] = 0
	furthestV := 0
	for q.Len() > 0 {
		v := q.Poll()
		furthestV = v
		for _, to := range graph[v] {
			if !isMarked[to] {
				dist[to] = dist[v] + 1
				edgeFrom[to] = v
				q.Offer(to)
				isMarked[to] = true
			}
		}
	}
	return furthestV
}

type queue []int

func (q *queue) Offer(i int) {
	*q = append(*q, i)
}

func (q *queue) Poll() int {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}

func (q *queue) Len() int {
	return len(*q)
}
