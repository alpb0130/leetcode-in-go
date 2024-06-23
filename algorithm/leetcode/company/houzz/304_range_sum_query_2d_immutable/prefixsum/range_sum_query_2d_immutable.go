package prefixsum

type NumMatrix struct {
	m [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	m := make([][]int, len(matrix)+1)
	for i := range m {
		m[i] = make([]int, len(matrix[0])+1)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			m[i+1][j+1] = matrix[i][j] + m[i][j+1] + m[i+1][j] - m[i][j]
		}
	}
	return NumMatrix{m}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.m[row2+1][col2+1] - this.m[row1][col2+1] - this.m[row2+1][col1] + this.m[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
