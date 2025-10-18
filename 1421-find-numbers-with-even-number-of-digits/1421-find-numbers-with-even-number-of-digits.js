/**
 * @param {number[]} nums
 * @return {number}
 */
var findNumbers = function (nums) {
    let b = 0
    for (i of nums) {
        if (i.toString().length % 2 === 0) {
            b++
        }
    }
    return b
};