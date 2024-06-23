package dfs

// Use dfs to detect cycle in directed graph. Need to pass into a node stack for each
// dfs call.
// Time: O(V+E)
// Space: O(V+E)
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	isMarked := make([]bool, numCourses)
	onStack := make([]bool, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	for i := 0; i < numCourses; i++ {
		if !isMarked[i] {
			// dfs
			if !dfs(graph, isMarked, onStack, i) {
				return false
			}
		}
	}
	return true
}

func dfs(graph [][]int, isMarked []bool, onStack []bool, course int) bool {
	isMarked[course] = true
	onStack[course] = true
	for _, toCourse := range graph[course] {
		if !isMarked[toCourse] && !dfs(graph, isMarked, onStack, toCourse) {
			return false
		} else if onStack[toCourse] {
			return false
		}
	}
	onStack[course] = false
	return true
}
