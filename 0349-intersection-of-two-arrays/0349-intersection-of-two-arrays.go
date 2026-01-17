func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	ch := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				ch[nums1[i]]++
			}
		}
	}

	for i, v := range ch {
		if v > 0 {
			result = append(result, i)
		}
	}
	return result
}