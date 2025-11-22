func findWordsContaining(words []string, x byte) []int {
	c := []int{}
	for i := 0; i < len(words); i++ {
		a := words[i]
		for _, v := range a {
			if byte(v) == x {
				c = append(c, i)
				break
			}
		}
	}
	return c
}