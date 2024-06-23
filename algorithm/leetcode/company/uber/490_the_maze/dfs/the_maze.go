package dfs

// DFS.
// Time: O(m*n)
// Space: O(m*n)
func hasPath(maze [][]int, start []int, destination []int) bool {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return true
	}
	// DFS
	return dfs(maze, start, destination, map[int]bool{})
}

func dfs(maze [][]int, start []int, end []int, isVisited map[int]bool) bool {
	index := start[0]*len(maze[0]) + start[1]
	if isVisited[index] {
		return false
	}
	if start[0] == end[0] && start[1] == end[1] {
		return true
	}
	isVisited[index] = true
	// Move up
	i, j := start[0], start[1]
	for i-1 >= 0 && maze[i-1][j] != 1 {
		i--
	}
	if dfs(maze, []int{i, j}, end, isVisited) {
		return true
	}
	// Move down
	i, j = start[0], start[1]
	for i+1 < len(maze) && maze[i+1][j] != 1 {
		i++
	}
	if dfs(maze, []int{i, j}, end, isVisited) {
		return true
	}
	// Move left
	i, j = start[0], start[1]
	for j-1 >= 0 && maze[i][j-1] != 1 {
		j--
	}
	if dfs(maze, []int{i, j}, end, isVisited) {
		return true
	}
	// Move right
	i, j = start[0], start[1]
	for j+1 < len(maze[0]) && maze[i][j+1] != 1 {
		j++
	}
	if dfs(maze, []int{i, j}, end, isVisited) {
		return true
	}
	//isVisited[index] = false
	return false
}
