func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])

	left, right := 0, row*col - 1

	for left <= right {
		mid := left + (right-left)/2

		r := mid / col
		c := mid % col

		val := matrix[r][c]
		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
