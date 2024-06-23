package dfs

// Use graph and DFS. Every equation can be represented by two vertices and an edge. The value
// can be represent the weight. You need to take care of some edge case. DFS to search for the
// path from numerator to denominator and calculate the value.
// Time: O(n*(V+E)) - n is number of queries.
// Space: O(V+E) - V is number of variables. E is number of equations.
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string]map[string]float64{}
	for i, equation := range equations {
		if _, ok := graph[equation[0]]; !ok {
			graph[equation[0]] = map[string]float64{}
		}
		graph[equation[0]][equation[1]] = values[i]
		if _, ok := graph[equation[1]]; !ok {
			graph[equation[1]] = map[string]float64{}
		}
		graph[equation[1]][equation[0]] = 1.0 / values[i]
	}
	// calculate value for each query
	res := []float64{}
	for _, query := range queries {
		res = append(res, calc(query, graph))
	}
	return res
}

func calc(query []string, graph map[string]map[string]float64) float64 {
	if graph[query[0]] == nil || graph[query[1]] == nil {
		return -1
	}
	if query[0] == query[1] {
		return 1
	}
	// dfs to find a path and calculate value
	return dfs(graph, map[string]bool{}, query[0], query[1], 1.0)
}

func dfs(graph map[string]map[string]float64, isMarked map[string]bool, start, end string, curVal float64) float64 {
	isMarked[start] = true
	if start == end {
		return curVal
	}
	for to, val := range graph[start] {
		if !isMarked[to] {
			v := dfs(graph, isMarked, to, end, curVal*val)
			if v != -1.0 {
				return v
			}
		}
	}
	return -1.0
}
