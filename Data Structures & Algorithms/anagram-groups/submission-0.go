func groupAnagrams(strs []string) [][]string {
	mem := make(map[string][]string)

	for _, v := range strs {
		b := []byte(v)
		sort.Slice(b, func(i int, j int) bool {
			return b[i] < b[j]
		})
		mem[string(b)] = append(mem[string(b)], v)
	}

	res := make([][]string, 0, len(mem))
	for _, v := range mem {
		res = append(res, v)
	}

	return res
}
