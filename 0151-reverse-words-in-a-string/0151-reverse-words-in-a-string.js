
var reverseWords = function (s) {
     let a = s.split(" ")
    return a.filter((b)=>b).reverse().toString().replaceAll(","," ")
};