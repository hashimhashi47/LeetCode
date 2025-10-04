/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
    let concat = nums1.concat(nums2).sort((a, b) => a - b)

    if (concat.length % 2 == 1) {
        let middle = concat[Math.floor(concat.length / 2)]
        return middle
    } else {
        let middle1 = concat[(concat.length / 2 - 1)]
        let middle2 = concat[(concat.length / 2)]
        return (middle1 + middle2) / 2
    }
};