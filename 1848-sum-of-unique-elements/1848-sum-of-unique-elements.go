func sumOfUnique(nums []int) int {
    a := make(map[int]int)
	b := 0
	for _, v := range nums {
		a[v]++
	}

	for _, v := range nums {
		if a[v] == 1 {
			b += v
		}
	}
	return b
}