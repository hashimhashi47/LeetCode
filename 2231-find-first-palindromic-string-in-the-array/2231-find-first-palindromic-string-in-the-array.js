/**
 * @param {string[]} words
 * @return {string}
 */
var firstPalindrome = function (words) {
       for (let i=0; i<words.length; i++){
        let a=words[i].split('').reverse().join('')
        if (words[i]===a){
            return a
            break
        }
    }
    return ("")
};