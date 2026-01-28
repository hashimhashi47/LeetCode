func dominantIndex(nums []int) int {
    temp := make([]int, len(nums))
	copy(temp, nums)
	for i := 0; i < len(temp)-1; i++ {
		for j := i + 1; j < len(temp); j++ {
			if temp[i] < temp[j] {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
	}
	if temp[0] >= 2*temp[1] {
		for i, v := range nums {
			if temp[0] == v {
				return i
			}
		}
	}
	return -1
}