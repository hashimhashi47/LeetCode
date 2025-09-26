/**
 * @param {number[][]} accounts
 * @return {number}
 */
var maximumWealth = function(accounts) {
     arr=[]
    for(let i=0; i<accounts.length; i++){
        let a=accounts[i]
       arr.push(a.reduce((a,b)=>a+b)) 
    }
    return Math.max(...arr)
};