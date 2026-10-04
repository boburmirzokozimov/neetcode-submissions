func groupAnagrams(strs []string) [][]string {
	store := make(map[[26]int][]string)
	
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		k := [26]int{}
		for j := 0; j < len(str); j++ {
			idx := str[j] - 'a'
			k[idx]++
		}
		store[k] = append(store[k], str)
	}

	out := make([][]string, 0, len(store))

	for _, val := range store {
		out = append(out, val)
	}

	return out
}
