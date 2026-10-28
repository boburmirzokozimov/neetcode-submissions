func checkInclusion(s1 string, s2 string) bool {
	count1 := [28]int{}
	count2 := [28]int{}

	for i := 0; i < len(s1); i++ {
		char := s1[i]-'a'

		count1[char]++
	}

	l := 0
	for r := 0; r < len(s2); r++ {
		char := s2[r] - 'a'


		if r - l + 1 > len(s1) {
			charL := s2[l] - 'a'
			count2[charL]--
			l++
		}

		count2[char]++
		
		if count1 == count2 {
			return true
		}
	}


	return false
}
