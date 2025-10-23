
var hasSameDigits = function (s) {
    let a = s.split('').map((a) => Number(a))
    while (a.length > 1) {
        let b = []
        for (let i = 0; i < a.length - 1; i++) {
            b.push((a[i] + a[i + 1]) % 10)
        }
        a = b
        if (a.length == 2) {
            break
        }
    }
    return a.reduce((a, b) => a === b)
};