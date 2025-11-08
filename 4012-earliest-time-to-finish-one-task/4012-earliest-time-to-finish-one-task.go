func earliestTime(tasks [][]int) int {
	c := 1<<31 - 1
	for i := 0; i < len(tasks); i++ {
		b := 0
		for j := 0; j < len(tasks[i]); j++ {
			b += tasks[i][j]
		}
		if b < c {
			c = b
		}
	}
	return c
}