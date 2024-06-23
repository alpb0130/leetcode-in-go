package topologicalsort

// Use topological sort. Always enqueue the course with in-degree 0. If a course has been
// processed, need to update the in-degree map. If there is a cycle, not all nodes will
// be processed and then the length of result list will be less than num of courses.
// Time: O(V+E)
// Space: O(V+E)
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	indegrees := make([]int, numCourses)
	res := []int{}
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
		res = append(res, course)
		for _, toCourse := range graph[course] {
			indegrees[toCourse]--
			if indegrees[toCourse] == 0 {
				q.Offer(toCourse)
			}
		}
	}
	if len(res) == numCourses {
		return res
	} else {
		return []int{}
	}
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
