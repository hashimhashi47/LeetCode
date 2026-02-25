func buildArray(nums []int) []int {
	Array := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		Array[i] = nums[nums[i]]
	}
	return Array
}