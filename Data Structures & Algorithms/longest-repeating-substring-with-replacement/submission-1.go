func characterReplacement(s string, k int) int {
	store := make(map[byte]int)
	maxF := 0
	l := 0
	longest := 0

	for i := 0; i < len(s); i++ {
		store[s[i]]++

		maxF = max(maxF, store[s[i]])

		for (i - l + 1) - maxF > k {
			store[s[l]]--
			l++
		}


		longest = max(longest, i - l + 1)
	}

	return longest
}
