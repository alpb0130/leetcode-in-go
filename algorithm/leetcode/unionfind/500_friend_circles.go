package unionfind

// Union find with path compression
// Time: O(n^2 * logn)
// Space: O(n)
func findCircleNum(M [][]int) int {
	if M == nil || len(M) <= 1 {
		return 1
	}
	num := len(M)
	fathers := make([]int, num)
	for i := 0; i < num; i++ {
		fathers[i] = i
	}
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if M[i][j] == 1 {
				union500(fathers, i, j)
			}
		}
	}
	numOfCircle := 0
	for i := 0; i < num; i++ {
		if fathers[i] == i {
			numOfCircle++
		}
	}
	return numOfCircle
}

func union500(fathers []int, i, j int) {
	root1 := find500(fathers, i)
	root2 := find500(fathers, j)
	if root1 == root2 {
		return
	}
	fathers[root1] = root2
}

func find500(fathers []int, index int) int {
	for index != fathers[index] {
		father := fathers[index]
		fathers[index] = fathers[father]
		index = father
	}
	return index
}

// DFS
// Time: O(n^2)
// Space: O(n)
func findCircleNum1(M [][]int) int {
	if M == nil || len(M) == 0 {
		return 0
	}
	num := len(M)
	isVisited := make([]bool, num)
	numOfCircles := 0
	for i := 0; i < num; i++ {
		if !isVisited[i] {
			numOfCircles++
			dfs500(M, isVisited, num, i)
		}
	}
	return numOfCircles
}

func dfs500(M [][]int, isVisited []bool, num, i int) {
	isVisited[i] = true
	for j := 0; j < num; j++ {
		if M[i][j] == 1 && !isVisited[j] {
			dfs500(M, isVisited, num, j)
		}
	}
}

// DFS
// Time: O(n^2)
// Space: O(n)
func findCircleNum2(M [][]int) int {
	if M == nil || len(M) == 0 {
		return 0
	}
	isVisited := make([]bool, len(M))
	numOfCircles := 0
	for i := 0; i < len(M); i++ {
		if !isVisited[i] {
			numOfCircles++
			bfs500(M, isVisited, i)
		}
	}
	return numOfCircles
}

func bfs500(M [][]int, isVisited []bool, i int) {
	queue := []int{}
	queue = append(queue, i)
	isVisited[i] = true
	for {
		if len(queue) == 0 {
			return
		}
		i := queue[0]
		queue = queue[1:]
		for j := 0; j < len(M); j++ {
			if !isVisited[j] && M[i][j] == 1 {
				queue = append(queue, j)
				isVisited[j] = true
			}
		}
	}
}
