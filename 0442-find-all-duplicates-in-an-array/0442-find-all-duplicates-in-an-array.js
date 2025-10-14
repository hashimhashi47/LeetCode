/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findDuplicates = function (nums) {
    let unique = new Set()
    const result = []
    for (n of nums) {
        unique.has(n) ? result.push(n) : unique.add(n)
    }
 return result
};