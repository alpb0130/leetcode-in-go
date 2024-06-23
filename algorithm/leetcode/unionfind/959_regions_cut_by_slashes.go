package unionfind

// Weighted union find with path compression
// Time: O(n^2*alpha(n^2))
// Space: O(n^2)
func regionsBySlashes(grid []string) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	n := len(grid)
	connections := make([]int, 2*n*n)
	size := make([]int, 2*n*n)
	for i, _ := range connections {
		connections[i] = i
		size[i] = 1
	}

	// iterate and union depend on the slash
	for i, row := range grid {
		for j, slash := range row {
			// the square
			if slash == ' ' {
				union_959(connections, size, i*n*2+j*2, i*n*2+j*2+1)
			}
			// join with left
			if j > 0 {
				union_959(connections, size, i*n*2+j*2, i*n*2+j*2-1)
			}
			// join with up
			if i > 0 {
				upperIndex := (i-1)*n*2 + j*2
				if grid[i-1][j] == '/' {
					upperIndex = (i-1)*n*2 + j*2 + 1
				}
				lowerIndex := i*n*2 + j*2
				if grid[i][j] == '\\' {
					lowerIndex = i*n*2 + j*2 + 1
				}
				union_959(connections, size, lowerIndex, upperIndex)
			}
		}
	}

	numOfRegions := 0
	for i, root := range connections {
		if i == root {
			numOfRegions++
		}
	}
	return numOfRegions
}

func union_959(connections, size []int, i, j int) {
	rootI := find_959(connections, i)
	rootJ := find_959(connections, j)
	if rootI == rootJ {
		return
	}
	if size[rootI] < size[rootJ] {
		connections[rootI] = rootJ
		size[rootJ] = size[rootI] + size[rootJ]
	} else {
		connections[rootJ] = rootI
		size[rootI] = size[rootI] + size[rootJ]
	}
}

func find_959(connections []int, i int) int {
	if i != connections[i] {
		connections[i] = find_959(connections, connections[i])
	}
	return connections[i]
}
