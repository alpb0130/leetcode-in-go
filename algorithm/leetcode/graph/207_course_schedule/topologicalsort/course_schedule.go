package topologicalsort

// Use topological sort and check whether the # of courses processed in queue is
// equal to # courses or not. If there is a cycle in the graph, not all node will be
// processed successfully, which means count will be less than numCourses
// Time: O(V+E)
// Space: O(V+E)
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	indegrees := make([]int, numCourses)
	count := 0
	// O(E)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		indegrees[prerequisite[0]]++
	}
	q := &queue{}
	//O(V)
	for course, indegree := range indegrees {
		if indegree == 0 {
			q.Offer(course)
		}
	}

	for q.Len() != 0 {
		course := q.Poll()
		count++
		for _, toCourse := range graph[course] {
			indegrees[toCourse]--
			if indegrees[toCourse] == 0 {
				q.Offer(toCourse)
			}
		}
	}
	return count == numCourses
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
