func canConstruct(ransomNote string, magazine string) bool {
	ch := make(map[string]int)
	for i := 0; i < len(ransomNote); i++ {
		ch[string(ransomNote[i])]++
	}
	for j := 0; j < len(magazine); j++ {
		ch[string(magazine[j])]--
	}
	for _, v := range ch {
		if 0 < v {
			return false
		}
	}

	return true
}