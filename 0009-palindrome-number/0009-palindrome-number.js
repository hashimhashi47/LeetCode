/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    const a= x.toString()
    const b= a.split('').reverse().join('')
    return b===a
};
