package bfsii

// Use BFS without modifying the original graph and mark each vertex with different labels.
// The graph is not guaranteed to be a single connected component. So you need
// to process all vertices in the graph. Even there are multiple connected component,
// the graph is still be able to be bipartite.
// Time: O(V+E)
// Space: O(V+D) - D is the max degree a vertex has
func isBipartite(graph [][]int) bool {
	labels := make([]int, len(graph))
	for from := range graph {
		if labels[from] == 0 {
			q := &queue{from}
			labels[from] = 1
			for q.Len() > 0 {
				v := q.Pop()
				for _, to := range graph[v] {
					if labels[to] != 0 && labels[to] == labels[v] {
						return false
					}
					if labels[to] == 0 {
						q.Push(to)
						labels[to] = -labels[v]
					}
				}
			}
		}
	}
	return true
}

type queue []int

func (q *queue) Pop() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}
func (q *queue) Push(v int) {
	*q = append(*q, v)
}
func (q *queue) Len() int {
	return len(*q)
}
