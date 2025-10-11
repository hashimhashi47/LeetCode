/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function (nums) {
    for (let i = 0; i < nums.length; i++) {
        let u = true;
        for (let j = 0; j < nums.length; j++) {
            if (i !== j && nums[i] === nums[j]) {
                u = false;
                break;
            }
        }
        if (u) {
            return nums[i];
        }
    }
};