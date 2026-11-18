func isValid(s string) bool {
    validParenthesesMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}

	for _, char := range s {
		if parantheses, ok := validParenthesesMap[byte(char)]; ok {
			if len(stack) > 0 {
				n := len(stack)
				last := stack[n-1]
				if last != parantheses {
					return false
				}
				stack = stack[:n-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, byte(char))
		}
	}

	return len(stack) == 0
}
