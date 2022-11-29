package main

import (
	// "fmt"
	"math"
	// "strings"
	// "math"
)

/*

Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer; in other words,
it is the product of some integer with itself.
For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.

81 / 3
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
*/

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func main() {
	n := 12
	numSquares(n)
}
