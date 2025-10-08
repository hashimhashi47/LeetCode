/**
 * @param {string} title
 * @return {string}
 */
var capitalizeTitle = function (title) {
    let a = title.toLowerCase().split(" ")
    let result = ""

    for (let i = 0; i < a.length; i++) {
        let b = a[i].split("")
        if (b.length <= 2) {
            result += a[i]  
        } else {
            for (let j = 0; j < b.length; j++) {
                if (j === 0) {
                    result += b[j].toUpperCase()
                } else {
                    result += b[j]
                }
            }
        }

        result += " "
    }

    return result.trim()
};
