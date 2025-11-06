func countPairs(nums []int, k int) int {
	a := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if i*j%k == 0 {
					a++
				}
			}

		}

	}
	return a
}