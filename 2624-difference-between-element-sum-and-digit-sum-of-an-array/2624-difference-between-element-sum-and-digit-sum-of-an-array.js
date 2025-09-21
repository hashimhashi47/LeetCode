/**
 * @param {number[]} nums
 * @return {number}
 */
var differenceOfSum = function(nums) {
     let b=nums.join("").split("")
     let a =nums.reduce((a,b)=>(a+b))
    let c=b.map(Number).reduce((a,b)=>(a+b))
    return (a-c)
    
};