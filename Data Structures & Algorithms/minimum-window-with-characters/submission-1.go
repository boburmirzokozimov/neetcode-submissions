func minWindow(s string, t string) string {
    var need [128]int

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	missing := len(t)

	l, bestL, bestLen := 0, 0, len(s) + 1

	for r := 0; r < len(s); r++ {

		if need[s[r]] > 0 {
			missing--
		}
		need[s[r]]--


		for missing == 0 {

			if r - l + 1 < bestLen {
				bestLen = r - l + 1
				bestL = l
			}

			need[s[l]]++
			if need[s[l]] > 0 {
				missing++
			}
			
			l++
		}
	}

	if bestLen > len(s) {
		return ""
	}

	return s[bestL:bestL+bestLen]
}
