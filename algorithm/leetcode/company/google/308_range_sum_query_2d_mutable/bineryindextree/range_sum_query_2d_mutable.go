package bineryindextree

type NumMatrix struct {
	tree   [][]int
	matrix [][]int
}

// 2D binary index tree
func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	tree := make([][]int, len(matrix)+1)
	for i := range tree {
		tree[i] = make([]int, len(matrix[0])+1)
	}
	matrixCopy := make([][]int, len(matrix))
	for i := range matrixCopy {
		matrixCopy[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			update(tree, i+1, j+1, matrix[i][j])
			matrixCopy[i][j] = matrix[i][j]
		}
	}
	return NumMatrix{tree, matrixCopy}
}

func update(tree [][]int, col, row, delta int) {
	for i := col; i < len(tree); i += i & (-i) {
		for j := row; j < len(tree[i]); j += j & (-j) {
			tree[i][j] += delta
		}
	}
}

func query(tree [][]int, col, row int) int {
	res := 0
	for i := col; i > 0; i -= i & (-i) {
		for j := row; j > 0; j -= j & (-j) {
			res += tree[i][j]
		}
	}
	return res
}

func (this *NumMatrix) Update(row int, col int, val int) {
	update(this.tree, row+1, col+1, val-this.matrix[row][col])
	this.matrix[row][col] = val
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return query(this.tree, row2+1, col2+1) - query(this.tree, row1, col2+1) - query(this.tree, row2+1, col1) + query(this.tree, row1, col1)
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */
