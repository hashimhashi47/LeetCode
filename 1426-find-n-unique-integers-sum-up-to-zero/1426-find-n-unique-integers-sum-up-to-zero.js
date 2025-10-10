/**
 * @param {number} n
 * @return {number[]}
 */
var sumZero = function (n) {
    arr = []
    for (let i = 1; i < n / 2 + .5; i++) {
        arr.push(i, -i)
    }

    if (n % 2 !== 0) {
        arr.push(0)
    }
    return arr
};