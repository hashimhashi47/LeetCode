/**
 * @param {number} n
 * @return {number}
 */
var subtractProductAndSum = function (n) {
    let a = n.toString().split("")
    let sum = 0
    let r = []
    for (i of a) {
        let m = Number(i)
        r.push(m)
    }
    let plus = r.reduce((a, c) => a + c, 0)
    let multiply = r.reduce((a, c) => a * c)
    return (multiply - plus)

};