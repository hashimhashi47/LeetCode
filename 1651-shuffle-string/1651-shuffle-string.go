func restoreString(s string, indices []int) string {
	a := make(map[int]string)
	b := ""
	for i := 0; i < len(s); i++ {
		a[indices[i]] = string(s[i])
	}
	for i := 0; i < len(s); i++ {
		b += a[i]
	}
	return b

}