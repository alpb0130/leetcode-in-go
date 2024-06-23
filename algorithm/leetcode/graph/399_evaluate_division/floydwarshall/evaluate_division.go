package floydwarshall

// Use Floyd-Warshall algorithm but not to find the short path for all vertex pair.
// We only want to go through all vertex pair in a reasonable way and get the value.
// Repeat computing is not a problem because value will be the same from time to time.
// Time: O(n+V^3)
// Space: O(V^2)
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string]map[string]float64{}
	for i, equation := range equations {
		if _, ok := graph[equation[0]]; !ok {
			graph[equation[0]] = map[string]float64{}
		}
		graph[equation[0]][equation[0]] = 1.0
		graph[equation[0]][equation[1]] = values[i]

		if _, ok := graph[equation[1]]; !ok {
			graph[equation[1]] = map[string]float64{}
		}
		graph[equation[1]][equation[1]] = 1.0
		graph[equation[1]][equation[0]] = 1.0 / values[i]
	}

	for mid := range graph {
		// Only compute values for vertex which are directly connect with mid
		// All other connections will be take care later
		for from := range graph[mid] {
			for to := range graph[mid] {
				graph[from][to] = graph[from][mid] * graph[mid][to]
			}
		}
	}

	res := []float64{}
	for _, query := range queries {
		val, ok := graph[query[0]][query[1]]
		if !ok {
			val = -1
		}
		res = append(res, val)
	}
	return res
}
