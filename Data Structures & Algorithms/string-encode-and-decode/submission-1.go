type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder

	for _, str := range strs {
		b.WriteString(strconv.Itoa(len(str)))
		b.WriteString("#")
		b.WriteString(str)
	}

	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)

	for i := 0; i < len(encoded); {
		//get the number of chars to read
		j := i
		for encoded[j] != '#' {
			j++
		}
		
		length, _ := strconv.Atoi(encoded[i:j])

		//skip #
		start := j + 1
		end := start + length

		//add string
		res = append(res, encoded[start:end])
		i = end
	}


	return res
}
