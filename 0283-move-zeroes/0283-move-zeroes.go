func moveZeroes(nums []int)  {
   a := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					a = nums[j]
					nums[j] = 0
					nums[i] = a
					break
				}
			}
		} 
    }
}