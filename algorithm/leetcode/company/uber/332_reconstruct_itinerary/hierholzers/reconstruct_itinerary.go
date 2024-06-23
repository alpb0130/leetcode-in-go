package hierholzers

import "container/heap"

// Use hierholzer's algorithm to find the Eulerian path - greedy DFS. DFS the graph until there is
// a node stuck and add this node to the front of the result list. Then backtracking to find next node
// stuck.
// Reference:
// https://www.geeksforgeeks.org/eulerian-path-and-circuit/
// https://www.geeksforgeeks.org/hierholzers-algorithm-directed-graph/
// https://en.wikipedia.org/wiki/Eulerian_path
// Time: O(Elog(degree))
// Space: O(V+E)
func findItinerary(tickets [][]string) []string {
	res := []string{}
	// construct the graph
	graph := map[string]*priorityQueue{}
	for _, ticket := range tickets {
		toCities, ok := graph[ticket[0]]
		if ok {
			heap.Push(toCities, ticket[1])
		} else {
			graph[ticket[0]] = &priorityQueue{ticket[1]}
		}
	}

	// greedy dfs
	dfs(graph, "JFK", &res)
	return res
}

func dfs(graph map[string]*priorityQueue, fromCity string, res *[]string) {
	toCities, ok := graph[fromCity]
	for ok && toCities.Len() > 0 {
		dfs(graph, heap.Pop(toCities).(string), res)
	}
	*res = append([]string{fromCity}, (*res)...)
}

type priorityQueue []string

func (q priorityQueue) Len() int {
	return len(q)
}
func (q priorityQueue) Less(i, j int) bool {
	return q[i] < q[j]
}
func (q priorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *priorityQueue) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}
func (q *priorityQueue) Push(i interface{}) {
	*q = append(*q, i.(string))
}
