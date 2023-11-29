package main

import (
	"unicode"
)

// "strconv"
// "strings"
// "math"

/*
A phrase is a palindrome if,
after converting all uppercase letters
into lowercase letters and
removing all non-alphanumeric characters,
it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s,
return true if it is a palindrome,
or false otherwis
*/
func isPalindrome(s string) bool {
	// a := 'a'
	// z := 'z'
	arr := []byte{}
	for i := range s {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			continue
		}
		p := s[i]
		if p < 'a' || p > 'z' {
			p = p + 32
		}
		arr = append(arr, p)
	}

	l, r := 0, len(arr)-1

	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// func isPalindrome(s string) bool {
// 	clearedStr := strings.Map(func(r rune) rune {
// 		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
// 			return -1
// 		}
// 		return unicode.ToLower(r)
// 	}, s)
// 	strArr := []rune(clearedStr)

// 	var leftPtr, rightPtr rune
// 	for i := 0; i < len(strArr); i++ {
// 		leftPtr = strArr[i]
// 		rightPtr = strArr[len(strArr)-1-i]

//			if i < len(strArr)-1-i {
//				if leftPtr != rightPtr {
//					return false
//				}
//			}
//		}
//		return true
//	}
func main() {
	// isPalindrome("A man, a plan, a canal: Panama")
	isPalindrome("0P")
}
