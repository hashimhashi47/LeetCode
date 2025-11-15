func lemonadeChange(bills []int) bool {
    a := 0 
    b := 0 

    for i := 0; i < len(bills); i++ {

        if bills[i] == 5 {
            a++
        }

        if bills[i] == 10 {
            if a == 0 {
                return false
            }
            a--
            b++
        }

        if bills[i] == 20 {
            if b > 0 && a > 0 {
                b--
                a--
            } else if a >= 3 {
                a -= 3
            } else {
                return false
            }
        }
    }

    return true
}