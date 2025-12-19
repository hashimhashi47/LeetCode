func firstUniqChar(s string) int {
	value := make(map[rune]int)
	for _, ch := range s {
		value[ch]++
	}
	for i, ch := range s {
		if value[ch] == 1 {
			return i
		}
	}
	return -1
}