func numJewelsInStones(jewels string, stones string) int {
	a := 0
	for i := 0; i < len(stones); i++ {
		for _,v:=range jewels{
			if string(v)==string(stones[i]){
				a++
			}
		}
	}
	return a
}