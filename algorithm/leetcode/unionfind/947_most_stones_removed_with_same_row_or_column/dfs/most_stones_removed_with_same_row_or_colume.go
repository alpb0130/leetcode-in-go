package dfs

// DFS
// n is the number of stone
// Time: O(n)
// Space: O(n)
func removeStones(stones [][]int) int {
	if stones == nil || len(stones) == 0 {
		return 0
	}

	graph := map[int][]int{}
	isVisited := map[int]bool{}
	for _, stone := range stones {
		xIndex := stone[0]
		yIndex := stone[1] + 10000
		// build graph
		buildGraph(graph, xIndex, yIndex)
		buildGraph(graph, yIndex, xIndex)
	}

	// dfs
	numOfGroup := 0
	for from := range graph {
		if !isVisited[from] {
			numOfGroup++
			dfs(graph, isVisited, from)
		}
	}
	return len(stones) - numOfGroup
}

func buildGraph(graph map[int][]int, x, y int) {
	list, ok := graph[x]
	if !ok {
		graph[x] = []int{y}
	} else {
		graph[x] = append(list, y)
	}
}

func dfs(graph map[int][]int, isVisited map[int]bool, x int) {
	isVisited[x] = true
	tolist := graph[x]
	for _, to := range tolist {
		if !isVisited[to] {
			dfs(graph, isVisited, to)
		}
	}
}
