package main

// "fmt"s
// "strconv"

/*
You are given an m x n integer array grid.
There is a robot initially located at the top-left corner
(i.e., grid[0][0]).
The robot tries to move to the bottom-right corner
(i.e., grid[m - 1][n - 1]).
The robot can only move either down or right at any point in time.

An obstacle and space are marked as 1 or 0
respectively in grid.
A path that the robot takes cannot include
any square that is an obstacle.

Return the number of possible unique paths
that the robot can take to reach the bottom-right corner.

The testcases are generated so that the answer
will be less than or equal to 2 * 109.

Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
Output: 2
Explanation: There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
*/

// func uniquePathsWithObstacles(obstacleGrid [][]int) int {
// 	m, n := len(obstacleGrid), len(obstacleGrid[0])
// 	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {
// 		return 0
// 	}
// 	dst := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		dst[i] = make([]int, n)
// 	}
// 	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
// 		dst[i][0] = 1
// 	}
// 	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
// 		dst[0][i] = 1
// 	}
// 	for i := 1; i < len(obstacleGrid); i++ {
// 		for j := 1; j < len(obstacleGrid[i]); j++ {
// 			if obstacleGrid[i-1][j] == 1 {
// 				dst[i-1][j] = 0
// 			}
// 			if obstacleGrid[i][j-1] == 1 {
// 				dst[i][j-1] = 0
// 			}
// 			dst[i][j] = dst[i-1][j] + dst[i][j-1]
// 		}
// 	}
// 	return dst[m-1][n-1]
// }
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	previous := make([]int, n)
	current := make([]int, n)
	previous[0] = 1

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			current[0] = 0
		} else {
			current[0] = previous[0]
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				current[j] = 0
			} else {
				current[j] = current[j-1] + previous[j]
			}
		}
		previous, current = current, previous
	}
	return previous[n-1]
}
func main() {
	grid := make([][]int, 3)
	for i := 0; i < 3; i++ {
		grid[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid[i][j] = 0
		}
	}
	// digits := []int{1, 2, 3, 9, 9}
	grid[1][1] = 1
	uniquePathsWithObstacles(grid)
}
