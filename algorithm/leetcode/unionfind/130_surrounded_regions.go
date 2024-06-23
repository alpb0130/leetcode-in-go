package unionfind

// DFS
// Time: O(n)
// Space: O(1)
func solve(board [][]byte) {
	if len(board) <= 1 {
		return
	}
	length := len(board)
	width := len(board[0])
	for i := 0; i < length; i++ {
		dfsHelper(board, i, 0)
		dfsHelper(board, i, width-1)
	}
	for j := 0; j < width; j++ {
		dfsHelper(board, 0, j)
		dfsHelper(board, length-1, j)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == 'o' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfsHelper(board [][]byte, i, j int) {
	if board[i][j] == 'O' {
		board[i][j] = 'o'
		if i-1 >= 0 {
			dfsHelper(board, i-1, j)
		}
		if i+1 < len(board) {
			dfsHelper(board, i+1, j)
		}
		if j-1 >= 0 {
			dfsHelper(board, i, j-1)
		}
		if j+1 < len(board[0]) {
			dfsHelper(board, i, j+1)
		}
	}
	return

}

// Weighted union find
// Time: O(nlogn)
// Space: O(n)
func solve1(board [][]byte) {
	if board == nil || len(board) <= 1 {
		return
	}
	m := len(board)
	n := len(board[0])

	set := make([]int, m*n+1)
	size := make([]int, m*n+1)
	for i := 0; i < m*n+1; i++ {
		set[i] = i
		size[i] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				union(set, size, i*n+j, m*n)
			}
			if j > 0 && board[i][j] == board[i][j-1] {
				union(set, size, i*n+j, i*n+j-1)
			}
			if i > 0 && board[i][j] == board[i-1][j] {
				union(set, size, i*n+j, (i-1)*n+j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !isConnected(set, i*n+j, m*n) {
				board[i][j] = 'X'
			}
		}
	}
}

func isConnected(set []int, a, b int) bool {
	return find(set, a) == find(set, b)
}

func find(set []int, a int) int {
	for a != set[a] {
		a = set[a]
	}
	return set[a]
}

func union(set, size []int, a, b int) {
	rootA := find(set, a)
	rootB := find(set, b)
	if size[rootA] <= size[rootB] {
		set[rootB] = rootA
		size[rootA] = size[rootA] + size[rootB]
	} else {
		set[rootA] = rootB
		size[rootB] = size[rootA] + size[rootB]
	}
}

// Union find with path compression
// Time: O(nlogn)
// Space: O(n)
func solve2(board [][]byte) {
	if board == nil || len(board) <= 1 {
		return
	}
	m := len(board)
	n := len(board[0])

	set := make([]int, m*n+1)
	for i := 0; i < m*n+1; i++ {
		set[i] = i
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				union2(set, i*n+j, m*n)
			}
			if j > 0 && board[i][j] == board[i][j-1] {
				union2(set, i*n+j, i*n+j-1)
			}
			if i > 0 && board[i][j] == board[i-1][j] {
				union2(set, i*n+j, (i-1)*n+j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !isConnected2(set, i*n+j, m*n) {
				board[i][j] = 'X'
			}
		}
	}
}

func isConnected2(set []int, a, b int) bool {
	return find2(set, a) == find2(set, b)
}

func find2(set []int, a int) int {
	for a != set[a] {
		b := set[a]
		set[a] = set[b]
		a = b
	}
	return set[a]
}

func union2(set []int, a, b int) {
	rootA := find2(set, a)
	rootB := find2(set, b)
	set[rootB] = rootA
}

