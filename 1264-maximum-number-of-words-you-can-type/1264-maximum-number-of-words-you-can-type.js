/**
 * @param {string} text
 * @param {string} brokenLetters
 * @return {number}
 */
var canBeTypedWords = function (text, brokenLetters) {
    let w = text.split(" ")
    let chr = brokenLetters.split("")
    let counter = w.length


    for (let i = 0; i < w.length; i++) {
        for (l of chr) {
            if (w[i].includes(l)) {
                counter--
                break;
            }
        }

    }
    return counter
};
