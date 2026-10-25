func checkInclusion(s1 string, s2 string) bool {
	n := len(s1)

	need := [26]int{}
	have := [26]int{}

	for i := 0; i < n; i++ {
		need[s1[i]-'a']++
	}

	for j := 0; j < len(s2); j++ {
		have[s2[j]-'a']++


		if j > n - 1 {
			have[s2[j-n]-'a']--
		} 
				
		if have == need {
			return true
		}
		
	}


	return false
}
