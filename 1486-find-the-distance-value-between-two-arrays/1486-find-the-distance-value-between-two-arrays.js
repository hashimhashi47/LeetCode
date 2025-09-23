/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @param {number} d
 * @return {number}
 */
var findTheDistanceValue = function (arr1, arr2, d) {
    let c=0
   for(let i=0; i<arr1.length; i++){
        for(j=0; j<arr2.length; j++){
            let count=Math.abs(arr1[i]-arr2[j])
            if(count<=d){
                   c++  
                   break;
            }
            }
        }
        return arr1.length-c
    };