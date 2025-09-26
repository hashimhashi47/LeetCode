/**
 * @param {number[]} nums
 * @return {number[]}
 */
var runningSum = function(nums) {
        let a=[]
       let b=0
    for(let i=0; i<nums.length; i++){
        b=b+nums[i]
        a.push(b)
    }
    return a
};