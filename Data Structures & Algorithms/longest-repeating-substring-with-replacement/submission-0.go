func characterReplacement(s string, k int) int {
	store := [26]int{}
	l := 0
	maxCount := 0
	best := 0

	for r := 0; r < len(s); r++ {
		char := s[r] - 'A'
		store[char]++
		maxCount = max(maxCount, store[char])
		curr := r-l+1

		if curr - maxCount > k {
			store[s[l]-'A']--
			l++
		} 
		best = max(best, r-l+1)
	}	
	//XYXXYYY. 2



	return best
}
