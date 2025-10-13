/**
 * @param {string[]} word1
 * @param {string[]} word2
 * @return {boolean}
 */
var arrayStringsAreEqual = function (word1, word2) {
    let b = word1.join("").split().sort().toString()
    let c = word2.join("").split().sort().toString()
    return b === c
};