func majorityElement(nums []int) int {
	ch := make(map[int]int)
	a := 0
	for i := 0; i < len(nums); i++ {
		ch[nums[i]]++
		if len(nums)/2 < ch[nums[i]] {
			a = nums[i]
		}
	}
    return a
}