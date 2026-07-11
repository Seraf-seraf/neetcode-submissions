type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := strings.Builder{}

	for _, str := range strs {
		l := len(str)
		lStr := strconv.Itoa(l)

		res.WriteString(lStr+"#"+str)
	}

	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0

	for i < len(encoded) {
		j := i

		for string(encoded[j]) != "#" {
			j++
		}

		l, _ := strconv.Atoi(encoded[i:j])
		start := j+1
		end := start+l

		res = append(res, encoded[start:end])
		i = end
	}

	return res
}
