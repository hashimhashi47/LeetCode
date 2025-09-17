/**
 * @param {number[]} nums
 * @return {number[]}
 */
var numberGame = function (nums) {
  let sor = nums.sort((a, b) => a - b)
    let store = []
    for (let i = 0; i < sor.length; i += 2) {
         let pair = sor.slice(i, i + 2);
         let swap=pair.reverse()
         store.push(swap)
    }
    return store.flat()
};