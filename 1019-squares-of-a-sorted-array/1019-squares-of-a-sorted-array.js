/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortedSquares = function (nums) {
    return nums.map((nums) => nums * nums).sort((a, b) => a - b)
};