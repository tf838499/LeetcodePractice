package main

import (
	"fmt"
	// "strconv"
)

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
// 	t := s + s
// 	t = t[1 : len(t)-1]
// 	for i := 0; i < len(t); i++ {
// 		if i+len(s) > len(t) {
// 			break
// 		}
// 		if s == t[i:i+len(s)] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func repeatedSubstringPattern(s string) bool {

//     // helper: returns all indexes of a rune in string
//     // use it to find indexes of starting char
//     indexesOf := func(r rune) []int{
//         idxs := []int{}
//         for i, c := range s { if c == r { idxs = append(idxs, i)} }
//         return idxs
//     }

//     //helper: returns true if s1 and s2 are equal
//     areSame := func(s1, s2 string) bool {
//         if len(s1) != len(s2) {return false}
//         i := 0
//         for ; i < len(s1); i++ {
//             if s1[i] != s2[i] {break}
//         }
//         return i == len(s1)
//     }

//     // helper: checks if we have a repeating substring of length lss
//     hasSubstringOfLen := func(x int) bool {
//         ls := len(s)
//         if (ls % x) != 0 {return false}

//         s0 := s[:x]
//         for i := x; i < ls; i += x {
//             si := s[i:i+x]
//             if !areSame(s0, si) {return false}
//         }

//         return true
//     }

//     rs := []rune(s)
//     c := rs[0]
//     idxs := indexesOf(c)
//     // fmt.Println(idxs)
//     // skip the first index as it is 0 for starting char
//     for _, ix := range idxs[1:] {
//         if hasSubstringOfLen(ix) { return true }
//     }
//     return false
// }

func repeatedSubstringPattern(s string) bool {
loop:
	// i: repeat times
	//  case abcabc or abcabcabc , Confirm how many can be split sub(find common factor),
	// bc s len is even number if one char is differect or miss can not
	for i := 2; i <= len(s); i++ {
		if len(s)%i != 0 {
			continue
		}
		subSize := len(s) / i
		for j := 0; j < subSize; j++ {
			for k := 1; k < i; k++ {
				fmt.Println(string(s[subSize*k+j]))
				if s[subSize*k+j] != s[j] {
					continue loop // similar breakpoint ,bc don't keep j loop
				}
			}
		}
		return true
	}
	return false
}

func main() {
	// s := "bb"
	// repeatsedSubstringPattern(s)
	s := "bba"
	repeatedSubstringPattern(s)
	// s = "abcabcabc"
	// repeatedSubstringPattern(s)
	s = "abcdeabcdeabcdeabcdeabcdeabcdeabcde"
	repeatedSubstringPattern(s)
}
