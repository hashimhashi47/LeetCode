func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i++ {
		a := nums[0]
		for j := 0; j < len(nums); j++ {
			if a > nums[j] {
				a = nums[j]
			}
		}
		index := slices.Index(nums, a)
		nums[index] = a * multiplier
		fmt.Println(nums)
	}

	return nums
}