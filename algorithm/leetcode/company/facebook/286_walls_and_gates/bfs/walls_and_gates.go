package bfs

// Use BFS. Start with gates and put those index to queue then we bfs the graph and
// every time we update a room value, we put it into queue so that we can update others.
// All position should only be visit once.
// Time: O(m*n)
// Space: O(m*n)
func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}
	m := len(rooms)
	n := len(rooms[0])
	q := &queue{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rooms[i][j] == 0 {
				q.Offer(Index{i, j})
			}
		}
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for q.Len() > 0 {
		index := q.Poll()
		for _, direction := range directions {
			i := index.i + direction[0]
			j := index.j + direction[1]
			if i < 0 || i >= m || j < 0 || j >= n || rooms[i][j] == -1 ||
				rooms[i][j] <= rooms[index.i][index.j]+1 {
				continue
			}
			rooms[i][j] = rooms[index.i][index.j] + 1
			q.Offer(Index{i, j})
		}
	}
	return
}

type Index struct {
	i int
	j int
}

type queue []Index

func (q *queue) Len() int {
	return len(*q)
}
func (q *queue) Offer(i Index) {
	*q = append(*q, i)
}
func (q *queue) Poll() Index {
	i := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return i
}
