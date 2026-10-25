func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr) - 1

	for r - l + 1 > k {
		if abs(arr[r] - x) >= abs(arr[l] - x) {
			r--
		} else {
			l++
		}
	}

	return arr[l:r+1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
