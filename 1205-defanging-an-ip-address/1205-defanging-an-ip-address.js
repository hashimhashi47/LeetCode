/**
 * @param {string} address
 * @return {string}
 */
var defangIPaddr = function (address) {
    for (let char of address) {
        if (char === ".") {
            return address.replaceAll(".", "[.]")
        }
    }
};