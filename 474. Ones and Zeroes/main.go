package main

// "sort"
// "strconv"

/*
You are given an array of binary strings strs and two integers m and n.

Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.

A set x is a subset of a set y if all elements of x are also elements of y.

Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
Output: 4
Explanation: The largest subset with at most 5 0's and 3 1's is {"10", "0001", "1", "0"}, so the answer is 4.
Other valid but smaller subsets include {"0001", "1"} and {"10", "1", "0"}.
{"111001"} is an invalid subset because it contains 4 1's, greater than the maximum of 3.
*/

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range strs {
		zero := 0
		one := 0
		for p := range strs[i] {
			if strs[i][p] == '0' {
				zero++
			} else {
				one++
			}
		}

		for j := m; j >= zero; j-- {
			for k := n; k >= one; k-- {
				dp[j][k] = max(dp[j-zero][k-one]+1, dp[j][k])
			}
		}

	}
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// [1,2,3,4,5,6,7,8,9,10,1,1,1,1,1]
	input := []string{"10", "0001", "111001", "1", "0"}
	mInput := 5
	nInput := 3
	// input := []int{1000}
	findMaxForm(input, mInput, nInput)
}
