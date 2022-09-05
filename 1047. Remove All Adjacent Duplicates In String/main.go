package main

import (
// "fmt"
// "strconv"
)

/*
You are given a string s consisting of lowercase English letters.
A duplicate removal consists of choosing two adjacent and equal letters and removing them.

We repeatedly make duplicate removals on s until we no longer can.

Return the final string after all such duplicate removals have been made.
It can be proven that the answer is unique.

Input: s = "abbaca"
Output: "ca"
Explanation: For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal,
and this is the only possible move.
The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
Input: s = "azxxzy"
Output: "ay"
*/
func removeDuplicates(s string) string {
	// ans := ""
	stack := make([]int32, 0)
	// fmt.Print(s[0])
	for _, i := range []byte(s) {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		if stack[len(stack)-1] == i {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, i)
		}
	}
	// for i, _ := range stack {
	// 	ans = ans + string(stack[i])
	// }
	return string(stack)
}

func main() {
	foo := "abbaca"
	removeDuplicates(foo)
}
