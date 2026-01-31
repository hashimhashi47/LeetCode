func addStrings(num1 string, num2 string) string {
    a := new(big.Int)
	b := new(big.Int)
	a.SetString(num1, 10)
	b.SetString(num2, 10)
	sum := new(big.Int)
	sum.Add(a,b)
	return sum.String()
}