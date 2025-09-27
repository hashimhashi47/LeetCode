/**
 * @param {number[]} seats
 * @param {number[]} students
 * @return {number}
 */
var minMovesToSeat = function (seats, students) {
    let a = []
    seats.sort((a, b) => a - b)
    students.sort((a, b) => a - b)
    for (let i = 0; i < seats.length; i++) {
        a.push(Math.abs(seats[i] - students[i]));
    }
    return a.reduce((a, b) => a + b)
};