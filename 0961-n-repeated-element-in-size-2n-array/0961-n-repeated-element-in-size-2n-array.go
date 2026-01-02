func repeatedNTimes(nums []int) int {
	ch := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		ch[nums[i]]++
	}
	maxCount := 0
	a := 0
	for key, value := range ch {
		if value > maxCount {
			maxCount = value
			a = key
		}
	}
	return a
}