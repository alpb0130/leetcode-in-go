package bfs

// Use graph and BFS. Every equation can be represented by two vertices and an edge. The value
// can be represent the weight. You need to take care of some edge case. BFS to search for the
// shortest path from numerator to denominator and calculate the value.
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
	// bfs to find a path and calculate value
	return bfs(graph, query[0], query[1])
}

func bfs(graph map[string]map[string]float64, start, end string) float64 {
	q := &queue{start}
	isMarked := map[string]bool{start: true}
	curVal := map[string]float64{start: 1.0}
	for q.Len() != 0 {
		from := q.Poll()
		for to, val := range graph[from] {
			if !isMarked[to] {
				if to == end {
					return curVal[from] * val
				}
				curVal[to] = curVal[from] * val
				isMarked[to] = true
				q.Offer(to)
			}
		}
	}
	return -1.0
}

type queue []string

func (q *queue) Offer(i string) {
	*q = append(*q, i)
}

func (q *queue) Poll() string {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}

func (q *queue) Len() int {
	return len(*q)
}
