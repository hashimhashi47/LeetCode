func prefixCount(words []string, pref string) int {
    	result := 0
	for i := 0; i < len(words); i++ {
		isit := strings.HasPrefix(words[i], pref)
		if isit {
			result++
		}
	}
	return result
}