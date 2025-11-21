func sortPeople(names []string, heights []int) []string {
	a := make(map[int]string)
	c := []string{}
	b := make([]int, len(heights))
	copy(b, heights)
	sort.Ints(b)
	for i := 0; i < len(names); i++ {
		a[heights[i]] = names[i]
	}

	for j := len(b) - 1; j >= 0; j-- {
		c = append(c, a[b[j]])
	}

	return c

}