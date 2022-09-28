package main

import (
	"fmt"
	// "strconv"
)

/*
Easy
stack
done

Given a string s containing just the characters
'(', ')', '{', '}', '[' and ']',
 determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.

Input: s = "()"
Output: true
Input: s = "()[]{}"
Output: true
Input: s = "(]"
Output: false
*/
func isValid(s string) bool {

	parens := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]int32, 0)

	for _, j := range []rune(s) {
		if len(stack) == 0 {
			stack = append(stack, j)
			continue
		}

		if stack[len(stack)-1] != parens[j] {
			stack = append(stack, j)
		} else {
			stack = stack[0 : len(stack)-1]
		}

	}

	return len(stack) == 0
}
func main() {
	foo := "([])"
	isValid(foo)

}
