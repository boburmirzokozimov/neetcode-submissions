func maxArea(h []int) int {
	sum := 0
	l := 0
	r := len(h)-1

	for l < r {
		w := r - l
		bestH := min(h[l], h[r])
		sum = max(sum, bestH*w)
		if h[l] > h[r] {
			r--
		} else {
			l++
		}
	}



	return sum
}
