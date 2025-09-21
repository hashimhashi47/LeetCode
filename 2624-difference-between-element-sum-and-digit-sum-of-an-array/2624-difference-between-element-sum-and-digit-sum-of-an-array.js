/**
 * @param {number[]} nums
 * @return {number}
 */
var differenceOfSum = function(nums) {
     let a =nums.reduce((a,b)=>(a+b))
     let b=nums.join("").split("")
    let c=b.map(Number).reduce((a,b)=>(a+b))
    return (a-c)
    
};