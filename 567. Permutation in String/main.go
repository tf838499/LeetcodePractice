package main

import (
	"fmt"
	// "strconv"
	// "strings"
	// "math"
)

/*
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
*/

// func checkInclusion(s1 string, s2 string) bool {
// 	s1_permutation := make(map[byte]int)
// 	// Input: s1 = "ab", s2 = "eidboaoo"
// 	s1_byte := []byte(s1)
// 	s2_byte := []byte(s2)
// 	for _, i := range s1_byte {
// 		s1_permutation[i]++
// 	}

// 	slow, fast := 0, len(s1_byte)
// 	for fast <= len(s2_byte) {
// 		if s1_permutation[s2_byte[slow]] != 0 {
// 			result := match(s1_permutation, s2_byte[slow:fast])
// 			if result {
// 				return result
// 			}
// 		}
// 		slow++
// 		fast = slow + len(s1_byte)
// 	}

// 	return false
// }
// func match(byte_s1 map[byte]int, byte_s2 []byte) bool {
// 	byte_s1_tmp := make(map[byte]int)
// 	for i, j := range byte_s1 {
// 		byte_s1_tmp[i] = j
// 	}
// 	for _, i := range byte_s2 {
// 		if byte_s1_tmp[i] == 0 {
// 			return false
// 		}
// 		byte_s1_tmp[i]--
// 	}
// 	return true
// }

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	v1 := make([]int, 26)
	v2 := make([]int, 26)
	for i := range s1 {
		v1[s1[i]-'a']++
		v2[s2[i]-'a']++
	}
	if match(v1, v2) {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		v2[s2[i]-'a']++
		v2[s2[i-len(s1)]-'a']--
		if s2[i] != s2[i-len(s1)] && match(v1, v2) {
			return true
		}
	}
	return false
}

func match(s1, s2 []int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
func main() {
	s1, s2 := "ab", "eidbaooo"
	// s1, s2 := "adc", "dcda"
	a := checkInclusion(s1, s2)
	fmt.Println(a)
}

// func checkInclusion(s1 string, s2 string) bool {
// 	if len(s2) < len(s1) {
// 		return false
// 	}
// 	if len(s1) == 0 {
// 		return true
// 	}
// 	v1 := make([]int, 26)
// 	v2 := make([]int, 26)
// 	for i := range s1 {
// 		v1[s1[i]-'a']++
// 		v2[s2[i]-'a']++
// 	}
// 	if match(v1, v2) {
// 		return true
// 	}
// 	for i := len(s1); i < len(s2); i++ {
// 		v2[s2[i]-'a']++
// 		v2[s2[i-len(s1)]-'a']--
// 		if s2[i] != s2[i-len(s1)] && match(v1, v2) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func match(s1, s2 []int) bool {
// 	for i := 0; i < 26; i++ {
// 		if s1[i] != s2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
