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
func generate(numRows int) [][]int {
	ans := [][]int{}
	tmp := []int{1}
	ans = append(ans, tmp)
	if numRows == 1 {
		return ans
	}
	ans = append(ans, gen(tmp, numRows)...)
	return ans
}
func gen(row []int, numRows int) [][]int {
	ans := []int{1}
	for i := 1; i < len(row); i++ {
		ans = append(ans, row[i-1]+row[i])
	}
	ans = append(ans, 1)
	result := [][]int{ans}
	if len(ans) >= numRows {
		return result
	}
	n := gen(ans, numRows)
	result = append(result, n...)

	return result
}
func main() {

	// digits := []int{1, 2, 3, 9, 9}
	generate(1)
}
