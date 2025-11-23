func separateDigits(nums []int) []int {
	var a string
	x := []int{}
	for _, v := range nums {
		a = strconv.Itoa(v)
		for _, b := range a {
			k, _ := strconv.Atoi(string(b))
			x = append(x, k)
		}
	}
	return x
}