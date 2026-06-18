func twoSum(nums []int, target int) []int {
    mem := make(map[int]int)

	for i, num := range nums {
		complement := target-num

		if _, ok := mem[complement]; !ok {
			mem[num] = i
			continue
		}

		return []int{mem[complement], i}
	}

	return []int{}
}
