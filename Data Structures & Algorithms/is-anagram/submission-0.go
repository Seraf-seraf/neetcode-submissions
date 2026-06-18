func isAnagram(s string, t string) bool {
	mem := make(map[byte]int)

	for k, _ := range s {
		mem[s[k]]++
	}

	for k, _ := range t {
		mem[t[k]]--
		if mem[t[k]] < 0 {
			return false
		}
	}

	for _, v := range mem {
		if v > 0 {
			return false
		}
	}

	return true
}
