func runningSum(nums []int) []int {
	b := 0
	a := []int{}
	for i := 0; i < len(nums); i++ {
		b = b + nums[i]
		a = append(a, b)
	}
	return a
}