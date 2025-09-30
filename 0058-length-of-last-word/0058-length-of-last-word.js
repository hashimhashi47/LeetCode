/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
      let a=s.trim().split(" ")
    let b=a[a.length-1]
    return b.length
};