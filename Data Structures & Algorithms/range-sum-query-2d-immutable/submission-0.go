type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows := len(matrix)
	cols := len(matrix[0])

	pr := make([][]int, rows + 1)
	for i := range pr {
		pr[i] = make([]int, cols + 1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			pr[r+1][c+1] = matrix[r][c] + pr[r][c+1] + pr[r+1][c] - pr[r][c]
		}
	}

	return NumMatrix {
		prefix: pr,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	p := this.prefix

	return p[row2+1][col2+1] - p[row1][col2+1] - p[row2+1][col1] + p[row1][col1]
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
