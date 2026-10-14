func validPalindrome(s string) bool {
	l, r := 0, len(s) - 1
	deletable := 1

	for l < r {
		if s[l] != s[r] {
			if deletable == 1 {
				valid := isValid(s, l + 1, r) || isValid(s, l, r - 1)
				if valid {
					l++
					r--
					continue
				}
				deletable--
			}
			return false
		}
		l++
		r--
	}

	return true
}

func isValid(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
