func canPlaceFlowers(flowerbed []int, n int) bool {
	a := 0
	b := 0
	if flowerbed[0] == 0 {
		a++
	}
	for i := 0; i < len(flowerbed); i++ {

		if flowerbed[i] == 0 {
			a++
		} else {
			a = 0
		}
		if a == 3 {
			b++
			a = 1
		}

	}

	if a == 2 {
		b++
	}

	return b >= n
}