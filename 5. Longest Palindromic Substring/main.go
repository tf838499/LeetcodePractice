package main

import "fmt"

// "fmt"
// "strconv"

/*

Given a string s, return the longest
palindromic

substring
 in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
*/

// func longestPalindrome(s string) string {
// 	right, left := len(s)-1, 0
// 	ansL, ansR := 0, 0
// 	for right != left {
// 		if s[right] == s[left] {
// 			ansL = left
// 			ansR = right
// 			left++
// 			right--
// 		} else {
// 			right--
// 		}
// 	}
// 	if right == left {
// 		return s[ansL : ansR+1]
// 	}
// 	return s
// }
func longestPalindrome(s string) string {
	ans := ""
	for start := 0; start < len(s); start++ {
		for end := len(s) - 1; end >= start; end-- {
			if s[start] == s[end] {
				if isPalindrome(s[start : end+1]) {
					if len(ans) < len(s[start:end+1]) {
						ans = s[start : end+1]
					}
				}
			}
		}
	}

	return ans
}
func isPalindrome(s string) bool {
	end := len(s) - 1
	for i := 0; i < end; i++ {
		if s[i] == s[end] {
			end--
		} else {
			return false
		}
	}
	return true
}
func main() {
	s := "aacabdkacaa"
	// s := "ac"
	ans := longestPalindrome(s)
	fmt.Println(ans)
}
