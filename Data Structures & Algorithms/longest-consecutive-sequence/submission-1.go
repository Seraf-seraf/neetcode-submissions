func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, v := range nums {
		numSet[v] = struct{}{}
	}

	res := 0

	for n, _ := range numSet {
		// Если в сете нет n-1, значит текущее n - старт цепочки
		if _, ok := numSet[n-1]; !ok {
			length := 1

			for {
				if _, ok := numSet[n+length]; ok {
					length++
				} else {
					break
				}
			}

			res = max(res, length)	
		}
	}

	return res
}
