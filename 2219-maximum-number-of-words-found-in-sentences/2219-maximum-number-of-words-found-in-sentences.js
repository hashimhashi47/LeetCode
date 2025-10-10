/**
 * @param {string[]} sentences
 * @return {number}
 */
var mostWordsFound = function (sentences) {
    arr = []
    for (let i = 0; i < sentences.length; i++) {
        a = 0
        let b = sentences[i].split(" ")
        for (let j = 0; j < b.length; j++) {
            a++
        }
        arr.push(a)
    }
    return Math.max(...arr)
};