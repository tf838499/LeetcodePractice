package main

import (
// "fmt"
// "strconv"
)

/*

Given n pairs of parentheses,
write a function to generate all combinations
of well-formed parentheses.
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Input: n = 1
Output: ["()"]
*/
// func generateParenthesis(n int) []string {

// }
func generateParenthesis(n int) []string {
	var res []string

	helper(n, 0, 0, "", &res)

	return res
}
func helper(n int, leftCount int, RightCount int, currentStr string, result *[]string) {
	if leftCount == n && RightCount == n {
		*result = append(*result, currentStr)
		return
	}
	if leftCount < n {
		helper(n, leftCount+1, RightCount, currentStr+"(", result)
	}
	if RightCount < leftCount {
		helper(n, leftCount, RightCount+1, currentStr+")", result)
	}
}

// func helper(n int, openCount int, closeCount int, currentStr string, result *[]string) {
// 	if openCount == n && closeCount == n {
// 		*result = append(*result, currentStr)
// 		return
// 	}

// 	if openCount < n {
// 		helper(n, openCount+1, closeCount, currentStr+"(", result)
// 	}

// 	if openCount > closeCount {
// 		helper(n, openCount, closeCount+1, currentStr+")", result)
// 	}
// }

func main() {
	generateParenthesis(3)
}
