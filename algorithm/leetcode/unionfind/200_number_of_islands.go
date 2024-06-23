package unionfind

// DFS
// Time: O(n)
// Space: O(n)
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	isVisited := make([][]bool, m)
	numOfIsland := 0
	for i := 0; i < m; i++ {
		isVisited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			isVisited[i][j] = false
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !isVisited[i][j] {
				numOfIsland++
				dfs(grid, isVisited, i, j, m, n)
			}
		}
	}
	return numOfIsland
}

func dfs(grid [][]byte, isVisited [][]bool, i, j, m, n int) {
	isVisited[i][j] = true
	if i > 0 && grid[i-1][j] == '1' && !isVisited[i-1][j] {
		dfs(grid, isVisited, i-1, j, m, n)
	}
	if i < m-1 && grid[i+1][j] == '1' && !isVisited[i+1][j] {
		dfs(grid, isVisited, i+1, j, m, n)
	}
	if j > 0 && grid[i][j-1] == '1' && !isVisited[i][j-1] {
		dfs(grid, isVisited, i, j-1, m, n)
	}
	if j < n-1 && grid[i][j+1] == '1' && !isVisited[i][j+1] {
		dfs(grid, isVisited, i, j+1, m, n)
	}
}

// DFS
// Time: O(n)
// Space: O(1)
func numIslands1(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	numOfIsland := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				numOfIsland++
				dfs1(grid, i, j, m, n)
			}
		}
	}
	return numOfIsland
}

func dfs1(grid [][]byte, i, j, m, n int) {
	grid[i][j] = '2'
	if i > 0 && grid[i-1][j] == '1' {
		dfs1(grid, i-1, j, m, n)
	}
	if i < m-1 && grid[i+1][j] == '1' {
		dfs1(grid, i+1, j, m, n)
	}
	if j > 0 && grid[i][j-1] == '1' {
		dfs1(grid, i, j-1, m, n)
	}
	if j < n-1 && grid[i][j+1] == '1' {
		dfs1(grid, i, j+1, m, n)
	}
}

// BFS
// Time: O(n)
// Space: O(n)
func numIslands2(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	numOfIsland := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				numOfIsland++
				bfs(grid, i, j, m, n)
			}
		}
	}
	return numOfIsland
}

type Index struct {
	x int
	y int
}

func bfs(grid [][]byte, i, j, m, n int) {
	queue := []Index{Index{i, j}}
	grid[i][j] = '2'
	for {
		if len(queue) == 0 {
			break
		}
		index := queue[0]
		queue = queue[1:]
		if index.x > 0 && grid[index.x-1][index.y] == '1' {
			grid[index.x-1][index.y] = '2'
			queue = append(queue, Index{index.x - 1, index.y})
		}
		if index.x < m-1 && grid[index.x+1][index.y] == '1' {
			grid[index.x+1][index.y] = '2'
			queue = append(queue, Index{index.x + 1, index.y})
		}
		if index.y > 0 && grid[index.x][index.y-1] == '1' {
			grid[index.x][index.y-1] = '2'
			queue = append(queue, Index{index.x, index.y - 1})
		}
		if index.y < n-1 && grid[index.x][index.y+1] == '1' {
			grid[index.x][index.y+1] = '2'
			queue = append(queue, Index{index.x, index.y + 1})
		}
	}
}

// Wighted union find
// Time: O(nlogn)
// Space: O(n)
func numIslands3(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	numOfIsland := 0
	fathers := make([]int, m*n)
	size := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fathers[i*n+j] = i*n + j
			size[i*n+j] = 1
		}
	}
	// union
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				if i > 0 && grid[i][j] == grid[i-1][j] {
					union200(fathers, size, i*n+j, (i-1)*n+j)
				}
				if j > 0 && grid[i][j] == grid[i][j-1] {
					union200(fathers, size, i*n+j, i*n+j-1)
				}
			}
		}
	}
	// check how many roots
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && fathers[i*n+j] == i*n+j {
				numOfIsland++
			}
		}
	}

	return numOfIsland
}

func union200(fathers, size []int, i, j int) {
	root1 := find200(fathers, i)
	root2 := find200(fathers, j)
	if root1 == root2 {
		return
	}
	if size[root1] >= size[root2] {
		fathers[root1] = root2
		size[root2] += size[root1]
	} else {
		fathers[root2] = root1
		size[root1] += size[root2]
	}
	return
}

func find200(fathers []int, i int) int {
	for i != fathers[i] {
		i = fathers[i]
	}
	return i
}

// Union find with path compression
// Time: O(nlogn)
// Space: O(n)
func numIslands4(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	numOfIsland := 0
	fathers := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fathers[i*n+j] = i*n + j
		}
	}
	// union
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				if i > 0 && grid[i][j] == grid[i-1][j] {
					union200_1(fathers, i*n+j, (i-1)*n+j)
				}
				if j > 0 && grid[i][j] == grid[i][j-1] {
					union200_1(fathers, i*n+j, i*n+j-1)
				}
			}
		}
	}
	// check how many roots
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && fathers[i*n+j] == i*n+j {
				numOfIsland++
			}
		}
	}

	return numOfIsland
}

func union200_1(fathers []int, i, j int) {
	root1 := find200_1(fathers, i)
	root2 := find200_1(fathers, j)
	if root1 == root2 {
		return
	}

	fathers[root1] = root2
	return
}

func find200_1(fathers []int, i int) int {
	for i != fathers[i] {
		index := fathers[i]
		fathers[i] = fathers[index]
		i = index

	}
	return i
}
