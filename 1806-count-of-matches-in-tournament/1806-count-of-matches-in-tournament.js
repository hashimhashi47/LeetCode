var numberOfMatches = function (n) {
    let a = []
    let b = n
    while (b > 1) {
        if (b % 2 == 0) {
            a.push(b / 2)
            b = (b / 2)
        } else {
            a.push((b - 1) / 2)
            b = (b - 1) / 2 + 1
        }
    }
    return a.reduce((a, b) => a + b, 0)
};