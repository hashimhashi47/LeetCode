/**
 * @param {number} x
 * @return {number}
 */
var sumOfTheDigitsOfHarshadNumber = function (x) {
    let y=x
let a=y.toString().split('').map(Number).reduce((a,b)=>a+b,0)

    if (a==0){
        return -1
    }

    if(x%a==0){
          return a
    }else{
        return -1
    }
};