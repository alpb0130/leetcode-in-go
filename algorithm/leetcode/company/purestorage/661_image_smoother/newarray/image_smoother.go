package newarray

// Time: O(m*n)
// Space: O(1)
func imageSmoother(M [][]int) [][]int {
	if len(M) == 0 || len(M[0]) == 0 {
		return [][]int{}
	}
	res := make([][]int, len(M))
	for i := range res {
		res[i] = make([]int, len(M[0]))
	}
	dir := [][]int{{-1, 0, 1}, {-1, 0, 1}}
	sum := 0
	cnt := 0
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if isValid(M, i+dir[0][k], j+dir[1][l]) {
						sum += M[i+dir[0][k]][j+dir[1][l]]
						cnt++
					}
				}
			}
			res[i][j] = sum / cnt
			sum = 0
			cnt = 0
		}
	}
	return res
}

func isValid(M [][]int, i, j int) bool {
	return i >= 0 && i < len(M) && j >= 0 && j < len(M[0])
}
