func heightChecker(heights []int) int {
	b := 0
	newarr := make([]int, len(heights))
	copy(newarr, heights)
	sort.Ints(newarr)
	for i := 0; i < len(heights); i++ {
		if heights[i] != newarr[i] {
			b++
		}
	}
	return b
}