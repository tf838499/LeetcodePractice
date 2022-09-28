package main

import (
	// "fmt"
	// "strconv"
	"strings"
)

/*
Medium
string
done

Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters.
The words in s will be separated by at least one space.
Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words.
The returned string should only have a single space separating the words. Do not include any extra spaces.

Input: s = "the sky is blue"
Output: "blue is sky the"
Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
*/
func reverseWords(s string) string {
	sa := strings.Split(s, " ")
	var ans string
	flag := 0
	stringsLen := len(sa)
	for i := 0; i < stringsLen; i++ {

		if sa[stringsLen-(i+1)] != "" {
			if flag == 1 {
				ans = ans + " "
			}
			ans = ans + sa[stringsLen-(i+1)]
			flag = 1
		}
		// if flag == 1 {
		// 	ans = ans + sa[stringsLen-(i+1)] + " "
		// }
		// if sa[stringsLen-(i+1)] != "" && i != stringsLen-1 {
		// 	ans = ans + sa[stringsLen-(i+1)] + " "
		// } else if i == stringsLen-1 && sa[stringsLen-(i+1)] != " " {
		// 	ans = ans + sa[stringsLen-(i+1)]
		// }
	}
	return ans
}

func main() {
	quet := "the sky is blue"
	reverseWords(quet)
}
