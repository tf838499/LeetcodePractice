package main

import (
	"fmt"
	// "strconv"
)

/*
easy
hashmap
done
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[string]int)
	n := make(map[string]int)
	a := []rune(s)
	b := []rune(t)
	// str := strings.Split("HELLO", "")
	for i := 0; i < len(s); i++ {
		m[string(a[i])] = m[string(a[i])] + 1
		n[string(b[i])] = n[string(b[i])] + 1
	}
	for i := 0; i < len(a); i++ {
		if m[string(a[i])] != n[string(a[i])] {
			return false
		}
	}
	return true
}

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	mS := make(map[rune]int)
// 	mT := make(map[rune]int)

// 	for _, v := range s {
// 		mS[v]++
// 	}

// 	for _, v := range t {
// 		mT[v]++
// 	}

// 	for k, v := range mS {
// 		if v != mT[k] {
// 			return false
// 		}
// 	}

// 	return true
// }

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,

ypically using all the original letters exactly once.

anagram : 同字母異序字,重組字,變位字

Input: s = "anagram", t = "nagaram"
Input: s = "rat", t = "car"

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
*/
func main() {
	// s := "anagram"
	// t := "nagaram"

	s := "rat"
	t := "car"
	fmt.Println(len(s))
	fmt.Println(len(t))
	fmt.Println(isAnagram(s, t))

}
