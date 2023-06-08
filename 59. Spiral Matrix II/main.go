package main

// "fmt"s
// "strconv"

/*
Medium
tree
Dynamic Programming
Divide and conquer
doing


Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]
*/

// func generateMatrix(n int) [][]int {

// 	var result = [][]int{}
// 	var tmp = make([][]int, n)
// 	for row := 0; row < len(tmp); row++ {
// 		tmp[row] = make([]int, n)
// 	}
// 	stard := 1

// 	for i := 0; i < n/2; i++ {
// 		for j := 0; j < n-1; j++ {
// 			tmp[i][j] = stard
// 			stard++
// 		}
// 		for j := 0; j < n; j++ {
// 			tmp[i+j][n-i-1] = stard
// 			stard++
// 		}
// 	}

//		return result
//	}
func generateMatrix(n int) [][]int {
	startx, starty := 0, 0
	var loop int = n / 2
	var center int = n / 2
	count := 1
	offset := 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := startx, starty

		//行数不变 列数在变
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}
		//列数不变是j 行数变
		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		//行数不变 i 列数变 j--
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		//列不变 行变
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}
	if n%2 == 1 {
		res[center][center] = n * n
	}
	return res
}
func main() {

	n := 3
	generateMatrix(n)
}
