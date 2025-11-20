func smallestEvenMultiple(n int) int {
	a := 1
	b := 0
	for i := 0; i <= a; i++ {
		b = n * a
		if b%2 == 0 {
			return b
		}
		a++
	}
	return b
}