func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	
	for i := 0; i < len(nums); i++ {
		n := 1

		for j, num := range nums {
			if j == i {
				continue
			}

			n *= num
		}

		res[i] = n
	}

	return res
}
