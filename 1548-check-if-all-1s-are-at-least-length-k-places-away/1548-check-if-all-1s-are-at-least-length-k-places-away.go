func kLengthApart(nums []int, k int) bool {
	a := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if a != -1 && i-a-1 < k {
				return false
			}
			a = i
		}
	}
	return true
}