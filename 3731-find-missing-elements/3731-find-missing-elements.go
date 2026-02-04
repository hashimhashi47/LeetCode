func findMissingElements(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	b := []int{}
	c := make(map[int]bool)
	for _, num := range nums {
		c[num] = true
	}
	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		if !c[i] {
			b = append(b, i)
		}
	}

	return b
}