func recoverOrder(order []int, friends []int) []int {
	a := []int{}
	for i := 0; i < len(order); i++ {
		for _, v := range friends {
			if order[i] == v {
				a = append(a, order[i])
			}
		}

	}
	return a
}