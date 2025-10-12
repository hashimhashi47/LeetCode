/**
 * @param {number} left
 * @param {number} right
 * @return {number[]}
 */
var selfDividingNumbers = function (left, right) {
    let arr = []
    let result = []
    for (let i = left; i <= right; i++) {
        arr.push(i)
    }

    for (let j = 0; j < arr.length; j++) {
        let p = arr[j].toString().split('').map(Number)
        let isSelfDividing = true
        for (let k of p) {
            if (k === 0 || arr[j] % k !== 0) {
                isSelfDividing = false
                break
            }
        }
        if (isSelfDividing) {
            result.push(arr[j])
        }
    }

    return result
};