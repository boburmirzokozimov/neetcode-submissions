func lengthOfLongestSubstring(s string) int {
	store := make(map[byte]bool)

	best := 0
	l := 0

	for r := 0; r < len(s); r++ {
		for store[s[r]] {
			delete(store, s[l])
			l++
		}
		store[s[r]] = true
		best = max(best, r-l+1)
	}

//abcabcbb=>0, 1
	return best
}
