func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ch := make(map[string]int)
	for i := 0; i < len(s); i++ {
		ch[string(s[i])]++
	}
	for j := 0; j < len(t); j++ {
		ch[string(t[j])]--
		if ch[string(t[j])] < 0 {
			return false
		}
	}
	fmt.Println(ch)
	return true
}