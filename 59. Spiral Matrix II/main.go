package main

// "fmt"s
// "strconv"

/*
Medium
tree
Dynamic Programming
Divide and conquer
doing


Given a positive integer n,
generate an n x n matrix filled with
elements from 1 to n2 in spiral order.
Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]
*/
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	i, j, d := 0, 0, 0
	dis := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for k := 1; k <= n*n; k++ {
		res[i][j] = k
		di, dj := dis[d%4][0], dis[d%4][1]
		// if
		if i+di >= n || i+di < 0 || j+dj >= n || j+dj < 0 || res[i+di][j+dj] != 0 {
			d++
			di, dj = dis[d%4][0], dis[d%4][1]
		}
		i, j = i+di, j+dj
	}

	return res
}

// func generateMatrix(n int) [][]int {
// 	res := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		res[i] = make([]int, n)
// 	}

// 	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// 	i, j, d := 0, 0, 0
// 	for k := 1; k <= n*n; k++ {
// 		res[i][j] = k

// 		di, dj := dirs[d%4][0], dirs[d%4][1]
// 		if i+di < 0 || i+di >= n || j+dj < 0 || j+dj >= n || res[i+di][j+dj] != 0 {
// 			d++
// 			di, dj = dirs[d%4][0], dirs[d%4][1]
// 		}

// 		i, j = i+di, j+dj
// 	}

// 	return res
// }

func main() {

	n := 4
	generateMatrix(n)
}
