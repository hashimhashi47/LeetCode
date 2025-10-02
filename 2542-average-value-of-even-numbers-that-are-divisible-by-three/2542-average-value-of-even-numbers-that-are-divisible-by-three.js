/**
 * @param {number[]} nums
 * @return {number}
 */
var averageValue = function (nums) {

    let a = []

    for (let i = 0; i < nums.length; i++) {
        if (nums[i] % 3 === 0 && nums[i] % 2 === 0) {
            a.push(nums[i])

        }

    }

    let b = a.reduce((a, b) => a + b, 0)

    return a.length > 0 ? Math.floor(b/ a.length) : 0;
};