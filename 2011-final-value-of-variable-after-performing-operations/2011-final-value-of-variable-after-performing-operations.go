func finalValueAfterOperations(operations []string) int {
	a := 0
	for i := 0; i < len(operations); i++ {
		fmt.Println(operations[i])
		if "++X" == operations[i] || "X++" == operations[i] {
			a += 1
		}
		if "--X" == operations[i] || "X--" == operations[i] {
			a -= 1
		}
	}
	return a
}