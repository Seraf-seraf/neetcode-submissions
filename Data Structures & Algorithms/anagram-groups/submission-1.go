func groupAnagrams(strs []string) [][]string {
	mem := make(map[[26]int][]string)

	for _, v := range strs {
		varKey := [26]int{}
		for _, l := range v {
			idx := int(l-'a')
			varKey[idx]++
		}

		mem[varKey] = append(mem[varKey], v)
	}

	res := make([][]string, 0, len(mem))
	for _, v := range mem {
		res = append(res, v)
	}

	return res
}
