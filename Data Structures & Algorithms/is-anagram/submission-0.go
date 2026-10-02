func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	store := make(map[rune]int)

	for _, v := range s {
		store[v]++
	}

	for _, v := range t {
		_, ok := store[v]
		if !ok {
			return false
		}
		store[v]--
		if store[v] == 0 {
			delete(store, v)
		}
	}

	return true
}
