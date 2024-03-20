package main

// "fmt"s
// "strconv"

/*
There is a robot on an m x n grid.
The robot is initially located at the top-left corner
(i.e., grid[0][0]).
The robot tries to move to the bottom-right corner
(i.e., grid[m - 1][n - 1]).
The robot can only move either down or right at any point in time.

Given the two integers m and n,
return the number of possible unique paths
that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be
less than or equal to 2 * 109.

Input: m = 3, n = 7
Output: 28
Example 2:

Input: m = 3, n = 3
Output: 3
Explanation: From the top-left corner,
there are a total of 3 ways to reach the bottom-right corner:
1. Right  Down  Down
4. Down   Right Down
5. Down   Down  Right


1. Right  Right  Down  Down
2. Right  Down  Right  Down
3. Right  Down  Down   Right
4. Down   Right Right  Down
5. Down   Right Down   Right
6. Down   Down  Right  Right


*/
// func uniquePaths(m int, n int) int {
// 	depth := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		depth[i] = make([]int, n)
// 		depth[i][0] = 1
// 	}
// 	for j := 0; j < n; j++ {
// 		depth[0][j] = 1
// 	}

// 	for i := 1; i < m; i++ {
// 		for j := 1; j < n; j++ {
// 			depth[i][j] = depth[i-1][j] + depth[i][j-1]
// 		}
// 	}
// 	return depth[m-1][n-1]
// }
func uniquePaths(m int, n int) int {
	ans := 1
	for i := 1; i <= m-1; i++ {
		ans = ans * (n - 1 + i) / i
	}
	return ans
}
func main() {

	// digits := []int{1, 2, 3, 9, 9}

	uniquePaths(3, 2)
}
