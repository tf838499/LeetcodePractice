package main

// "fmt"s
// "strconv"

/*
You are given a large integer represented as an integer array digits,
where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order.
The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.


Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] >= 10 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
				return digits
			}
		} else {
			return digits
		}
	}
	return digits
}
func plusOne(digits []int) []int {

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0

	}
	digits = append([]int{1}, digits...)
	return digits
}
func main() {

	digits := []int{1, 2, 3, 9, 9}
	plusOne(digits)
}
