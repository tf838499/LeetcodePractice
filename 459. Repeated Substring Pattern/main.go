package main

import "strings"

// "strconv"

/*
easy
string
kmp
doing

Given a string s, check if it can be constructed by taking a substring of it
and appending multiple copies of the substring together.

Input: s = "abab"
Output: true
Explanation: It is the substring "ab" twice.

Input: s = "aba"
Output: false

cdcdvasvcdcdvasv
Input: s = "abcabcabcabc"
Output: true
Explanation: It is the substring "abc" four times or the substring "abcabc" twice.
*/

// func repeatedSubstringPattern(s string) bool {
// loop:
// 	// i: repeat times
// 	//  case abcabc or abcabcabc , Confirm how many can be split sub(find common factor),
// 	// bc s len is even number if one char is differect or miss can not
// 	for i := 2; i <= len(s); i++ {
// 		if len(s)%i != 0 {
// 			continue
// 		}
// 		subSize := len(s) / i
// 		for j := 0; j < subSize; j++ {
// 			for k := 1; k < i; k++ {
// 				fmt.Println(string(s[subSize*k+j]))
// 				if s[subSize*k+j] != s[j] {
// 					continue loop // similar breakpoint ,bc don't keep j loop
// 				}
// 			}
// 		}
// 		return true
// 	}
// 	return false
// }

// 前缀表（不减一）的代码实现
// func repeatedSubstringPattern(s string) bool {
// 	n := len(s)
// 	if n == 0 {
// 		return false
// 	}
// 	j := 0
// 	next := make([]int, n)
// 	next[0] = j
// 	for i := 1; i < n; i++ {
// 		for j > 0 && s[i] != s[j] {
// 			j = next[j-1]
// 		}
// 		if s[i] == s[j] {
// 			j++
// 		}
// 		next[i] = j
// 	}
// 	// next[n-1]  最长相同前后缀的长度
// 	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
// 		return true
// 	}
// 	return false
// }

// // 这里使用了前缀表统一减一的实现方式
// func repeatedSubstringPattern(s string) bool {
// 	n := len(s)
// 	if n == 0 {
// 		return false
// 	}
// 	next := make([]int, n)
// 	j := -1
// 	next[0] = j
// 	for i := 1; i < n; i++ {
// 		for j >= 0 && s[i] != s[j+1] {
// 			j = next[j]
// 		}
// 		if s[i] == s[j+1] {
// 			j++
// 		}
// 		next[i] = j
// 	}
// 	// next[n-1]+1 最长相同前后缀的长度
// 	if next[n-1] != -1 && n%(n-(next[n-1]+1)) == 0 {
// 		return true
// 	}
// 	return false
// }
func repeatedSubstringPattern(s string) bool {

	size := len(s)
	sFold := s[1:size] + s[0:size-1]

	return strings.Contains(sFold, s)
}
func main() {
	// s := "bb"
	// repeatsedSubstringPattern(s)
	s := "abcabc"
	repeatedSubstringPattern(s)
	// s = "abcabcabc"
	// repeatedSubstringPattern(s)
	// s = "abcdeabcdeabcdeabcdeabcdeabcdeabcde"
	// repeatedSubstringPattern(s)
}
