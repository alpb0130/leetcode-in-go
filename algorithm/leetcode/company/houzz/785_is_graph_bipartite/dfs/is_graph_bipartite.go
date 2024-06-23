package dfs

// Use DFS and mark each vertex with different labels.
// The graph is not guaranteed to be a single connected component. So you need
// to process all vertices in the graph. Even there are multiple connected component,
// the graph is still be able to be bipartite.
// Time: O(V+E)
// Space: O(V+D) - D is the max degree a vertex has
func isBipartite(graph [][]int) bool {
	labels := make([]int, len(graph))
	for from := range graph {
		if labels[from] == 0 {
			if !dfs(graph, from, &labels, 1) {
				return false
			}
		}
	}
	return true
}

func dfs(graph [][]int, from int, labels *[]int, label int) bool {
	(*labels)[from] = label
	for _, to := range graph[from] {
		if (*labels)[to] != 0 && (*labels)[to] == label {
			return false
		}
		if (*labels)[to] == 0 {
			if !dfs(graph, to, labels, -label) {
				return false
			}
		}
	}
	return true
}
