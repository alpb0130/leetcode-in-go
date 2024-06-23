package bfs

// BFS.
// Time: O(m*n)
// Space: O(m*n)
func hasPath(maze [][]int, start []int, destination []int) bool {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return true
	}
	m := len(maze)
	n := len(maze[0])
	// BFS
	startIndex := start[0]*n + start[1]
	isVisited := map[int]bool{startIndex: true}
	q := &queue{startIndex}
	for q.Len() > 0 {
		p := q.Poll()
		i, j := p/n, p%n
		if i == destination[0] && j == destination[1] {
			return true
		}
		// Move up
		newi, newj := i, j
		for newi-1 >= 0 && maze[newi-1][newj] == 0 {
			newi--
		}
		newIndex := newi*n + newj
		if !isVisited[newIndex] {
			isVisited[newIndex] = true
			q.Offer(newIndex)
		}
		// Move down
		newi, newj = i, j
		for newi+1 < m && maze[newi+1][newj] == 0 {
			newi++
		}
		newIndex = newi*n + newj
		if !isVisited[newIndex] {
			isVisited[newIndex] = true
			q.Offer(newIndex)
		}
		// Move left
		newi, newj = i, j
		for newj-1 >= 0 && maze[newi][newj-1] == 0 {
			newj--
		}
		newIndex = newi*n + newj
		if !isVisited[newIndex] {
			isVisited[newIndex] = true
			q.Offer(newIndex)
		}
		// Move right
		newi, newj = i, j
		for newj+1 < n && maze[newi][newj+1] == 0 {
			newj++
		}
		newIndex = newi*n + newj
		if !isVisited[newIndex] {
			isVisited[newIndex] = true
			q.Offer(newIndex)
		}
	}
	return false
}

type queue []int

func (q *queue) Len() int {
	return len(*q)
}
func (q *queue) Offer(i int) {
	*q = append(*q, i)
}
func (q *queue) Poll() int {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}
