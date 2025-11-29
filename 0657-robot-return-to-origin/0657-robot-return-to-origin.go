func judgeCircle(moves string) bool {
	LR := 0
	UD := 0
	for _, v := range moves {
		if string(v) == "L" {
			LR++
		}
		if string(v) == "R" {
			LR--
		}
		if string(v) == "U" {
			UD++
		}
		if string(v) == "D" {
			UD--
		}
	}
	if LR == 0 && UD == 0 {
		return true
	}
	return false
}