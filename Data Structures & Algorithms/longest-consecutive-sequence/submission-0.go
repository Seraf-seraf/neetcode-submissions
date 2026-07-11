func longestConsecutive(nums []int) int {
	mem := make(map[int][]int)
	sort.Ints(nums)
	seen := make(map[int]struct{})
	j := 0
	for _, v := range nums {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			nums[j] = v
			j++
		}
	}

	nums = nums[:j:j]

	for _, n := range nums {
		if _, ok := mem[n]; !ok {
			mem[n] = []int{n}
		}

		if _, ok := mem[n-1]; ok {
			q := mem[n-1]
			q = append(q, n)

			mem[n-1] = nil

			mem[n] = q
		}
	}

	maxLen := 0

	for _, v := range mem {
		maxLen = max(maxLen, len(v))
	}

	return maxLen
}
