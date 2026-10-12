func isValidSudoku(board [][]byte) bool {
	rows := make(map[[2]int]bool)
	cols := make(map[[2]int]bool)
	boxes := make(map[[2]int]bool)

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			char := board[r][c]
			if char == '.' {
				continue
			}
			box := (r/3)*3 + c/3
			
			rowKey := [2]int{r, int(char)}
			colKey := [2]int{c, int(char)}
			boxKey := [2]int{box, int(char)}

			if rows[rowKey] || cols[colKey] || boxes[boxKey] {
				return false
			}

			rows[rowKey] = true
			cols[colKey] = true
			boxes[boxKey] = true
		}
	}

	return true
}