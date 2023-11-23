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
	for i := 0; i < numRows; i++ {
		tmp := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i || i-1 == 0 {
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, ans[i-1][j-1]+ans[i-1][j])
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 1; j < i; j++ {
			row[j] = ans[i-1][j] + ans[i-1][j-1]
		}
		row[0], row[i] = 1, 1
		ans[i] = row
	}
	return ans
}
func main() {

	// digits := []int{1, 2, 3, 9, 9}
	generate(9)
}
