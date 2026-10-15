func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := 0, 0
	
	var res strings.Builder

	for l1 < len(word1) && l2 < len(word2) {
		res.WriteByte(word1[l1])
		res.WriteByte(word2[l2])
		l1++
		l2++
	}
	if l1 < len(word1) {
		res.WriteString(word1[l1:])
	}
	if l2 < len(word2) {
		res.WriteString(word2[l2:])
	}

	return res.String()
}
