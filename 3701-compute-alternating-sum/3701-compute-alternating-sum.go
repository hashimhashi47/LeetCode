func alternatingSum(nums []int) int {
    sum := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			sum += nums[i]
		}
		if i%2 != 0 {
			sum -= nums[i]
		}
	}
	return sum
}