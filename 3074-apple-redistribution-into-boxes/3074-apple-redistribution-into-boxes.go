func minimumBoxes(apple []int, capacity []int) int {
	total := 0
	output := 0
	for i := 0; i < len(apple); i++ {
		total += apple[i]
	}
	sort.Ints(capacity)
	for j := len(capacity) - 1; j >= 0; j-- {
		total = total - capacity[j]
		output++
		if total <= 0 {
			break
		}
	}
	return output
}