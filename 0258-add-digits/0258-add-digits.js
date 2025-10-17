/**
 * @param {number} num
 * @return {number}
 */
var addDigits = function (num) {
    let a = num.toString().split('').map((a) => Number(a)).reduce((a, b) => a + b)
    if (a > 9) {
       return addDigits(a)
    } else {
        return a
    }
};