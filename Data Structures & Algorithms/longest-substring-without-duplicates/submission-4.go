func lengthOfLongestSubstring(s string) int {
	store := make(map[byte]bool)
	l, bestL := 0, 0

	for r := 0; r < len(s); r++ {
		for store[s[r]] {
			delete(store, s[l])
			l++
		}

		store[s[r]] = true
		bestL = max(bestL, r - l + 1)
		//pwwkew
	}

	return bestL
}
