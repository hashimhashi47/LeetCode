/**
 * @param {number[]} nums
 * @return {number}
 */
var differenceOfSum = function (nums) {
    let b = nums.join("").split("").map(Number).reduce((a, b) => (a + b))
    let a = nums.reduce((a, b) => (a + b))
    return (a - b)
};