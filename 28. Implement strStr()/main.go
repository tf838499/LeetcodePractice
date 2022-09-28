package main

import (
// "fmt"
// "strconv"
)

/*
Medium
string
done


Implement strStr().

Given two strings needle and haystack,
return the index of the first occurrence of needle in haystack,
or -1 if needle is not part of haystack.
Clarification:

What should we return when needle is an empty string?
This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string.
This is consistent to C's strstr() and Java's indexOf().

Input: haystack = "hello", needle = "ll"
Output: 2
Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	slow := -1
	var fast int
	if needleLen > len(haystack) {
		return slow
	}

	for i := 0; i < len(haystack)-(needleLen-1); i++ {
		if needle[0] == haystack[i] {
			slow, fast = i, 0
			for fast < needleLen && haystack[i+fast] == needle[fast] {
				fast++
			}
			if fast == needleLen {
				break
			} else {
				slow = -1
			}
		}
	}
	return slow
}

// func strStr(haystack string, needle string) int {
// 	if needle == "" {
// 		return 0
// 	}
// 	length := len(needle)
// 	for i := 0; i < len(haystack)-len(needle)+1; i++ {
// 		if haystack[i:i+length] == needle {
// 			return i
// 		}
// 	}
// 	return -1
// }

func main() {
	// hay := "mississippi"
	// nee := "issipi"
	hay := "aa"
	nee := "aa"

	// words := ["bella","label","roller"]
	strStr(hay, nee)
}
