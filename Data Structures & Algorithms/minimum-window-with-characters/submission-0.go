func minWindow(s string, t string) string {
    var need [128]int

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	missing := len(t)
	l, bestL, bestLen := 0, 0, len(s)+1

	for r := 0; r < len(s); r++ {
		ch := s[r]
		if need[ch] > 0 {
			missing--
		}
		need[ch]--

		for missing == 0 {

			if r-l+1 < bestLen {
				bestL, bestLen = l, r-l+1
			}

			lCh := s[l]
			need[lCh]++

			if need[lCh] > 0 {
				missing++
			}
			
			l++
		}
	}
	if bestLen > len(s) { 
		return ""
	}
	return s[bestL : bestL+bestLen]
}
