package dfs

// Use dfs.
// 1. Detect cycle in directed graph
// 2. The topological order will be the reversed dfs postorder.
// dfs call.
// Time: O(V+E)
// Space: O(V+E)
func findOrder(numCourses int, prerequisites [][]int) []int {
	res := []int{}
	graph := make([][]int, numCourses)
	isMarked := make([]bool, numCourses)
	onStack := make([]bool, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	for i := 0; i < numCourses; i++ {
		if !isMarked[i] {
			if !dfs(graph, isMarked, onStack, i, &res) {
				return []int{}
			}
		}
	}
	return res
}

func dfs(graph [][]int, isMarked []bool, onStack []bool, course int, res *[]int) bool {
	isMarked[course] = true
	onStack[course] = true
	for _, toCourse := range graph[course] {
		if !isMarked[toCourse] && !dfs(graph, isMarked, onStack, toCourse, res) {
			return false
		} else if onStack[toCourse] {
			return false
		}
	}
	*res = append([]int{course}, (*res)...)
	onStack[course] = false
	return true
}
