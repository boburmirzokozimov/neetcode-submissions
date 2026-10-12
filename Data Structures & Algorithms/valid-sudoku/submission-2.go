func isValidSudoku(board [][]byte) bool {
	rows := make(map[[2]int]bool)
	cols := make(map[[2]int]bool)
	boards := make(map[[2]int]bool)

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			char := board[r][c]

			if char == '.' {
				continue
			}

			box := (r/3)*3 + c/3

			row := [2]int{r, int(char)}
			col := [2]int{c, int(char)}
			board := [2]int{box, int(char)}

			if rows[row] {
				return false
			}
			if cols[col] {
				return false
			}
			if boards[board] {
				return false
			}
			rows[row] = true
			cols[col] = true
			boards[board] = true
		}
	}

	return true
}