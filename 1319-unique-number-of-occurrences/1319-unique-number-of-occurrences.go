func uniqueOccurrences(arr []int) bool {
	a := make(map[int]int)
	b := make(map[int]bool)
	for _, v := range arr {
		a[v]++
	}
	for _, d := range a {
		if b[d] {
			return false
		}
		b[d] = true
	}
	return true
}