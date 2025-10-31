func isSubsequence(s string, t string) bool {
	a := 0
	b := 0
	for _, v := range s {
		for i := b; i < len(t); i++ {
			if v == rune(t[i]) {
				a++
                b=i+1
				break
			}
		
		}
	}
	if a == len(s) {
		return true
	}
    return false
}