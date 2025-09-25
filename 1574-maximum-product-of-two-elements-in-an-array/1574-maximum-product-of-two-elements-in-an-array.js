/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {
        let arr=[]
    for(let i=0; i<nums.length; i++){
        for(let a= i+1; a<nums.length; a++)
       arr.push((nums[i] - 1) * (nums[a] - 1))
    }
    return Math.max(...arr)
};
