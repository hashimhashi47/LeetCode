/**
 * @param {string} word1
 * @param {string} word2
 * @return {string}
 */
var mergeAlternately = function (word1, word2) {
     let d = []
    for (let i = 0; i < word1.length+word2.length; i++) {
        if (d.length % 2 == 0) {
            d.push(word1[i])
        }
        if (d.length % 2 == 1) {
            d.push(word2[i])
        }
    }
return d.join('')
};