package main

// "strconv"

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
// func isValid(s string) bool {

// 	parens := map[rune]rune{
// 		')': '(',
// 		'}': '{',
// 		']': '[',
// 	}
// 	if len(s)%2 != 0 {
// 		return false
// 	}

// 	stack := make([]int32, 0)

// 	for _, j := range []rune(s) {
// 		if len(stack) == 0 {
// 			stack = append(stack, j)
// 			continue
// 		}

// 		if stack[len(stack)-1] != parens[j] {
// 			stack = append(stack, j)
// 		} else {
// 			stack = stack[0 : len(stack)-1]
// 		}

// 	}

// 	return len(stack) == 0
// }
func isValid(s string) bool {

	stack := []rune{}
	parens := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, i := range s {
		if len(stack) == 0 || parens[i] != stack[len(stack)-1] {
			stack = append(stack, i)
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
