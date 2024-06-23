package dfs

import "math"

// Get the minimum distance from start to dest in the maze
// Basically we can maintain the min distance to all positions and we
// use dfs to visit all possible path. If the current dist is smaller than
// previous distance, we can keep going and update the distance.
// Time: O()
// Space: O()
func shortestDistance(maze [][]int, start []int, destination []int) int {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return -1
	}
	distance := make([][]int, len(maze))
	for i := range distance {
		distance[i] = make([]int, len(maze[0]))
		for j := range distance[i] {
			distance[i][j] = math.MaxInt32
		}
	}
	distance[start[0]][start[1]] = 0
	// dfs and update the distance
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dfs(maze, start, distance, dir)
	if distance[destination[0]][destination[1]] == math.MaxInt32 {
		return -1
	}
	return distance[destination[0]][destination[1]]
}

func dfs(maze [][]int, start []int, distance [][]int, dir [][]int) {
	dist := distance[start[0]][start[1]]
	for _, d := range dir {
		i, j := start[0], start[1]
		step := 0
		for i+d[0] >= 0 && i+d[0] < len(maze) && j+d[1] >= 0 && j+d[1] < len(maze[0]) && maze[i+d[0]][j+d[1]] != 1 {
			i += d[0]
			j += d[1]
			step++
		}
		if i != start[0] || j != start[1] {
			if dist+step < distance[i][j] {
				distance[i][j] = dist + step
				dfs(maze, []int{i, j}, distance, dir)
			}
		}
	}
}
