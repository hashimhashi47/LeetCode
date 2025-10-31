/**
 * @param {number[]} nums
 * @return {number[]}
 */
var getSneakyNumbers = function(nums) {
     let array=[]
    for (let i = 0; i<nums.length; i++){
        for (let j = i+1; j<nums.length; j++){
            if (nums[i]==nums[j]){
                array.push(nums[i])
            }
        }
    }
    return array
};