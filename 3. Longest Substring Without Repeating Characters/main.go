package main

import (
// "fmt"
// "strconv"
)

/*

Given a string s, find the length of the longest substring without repeating characters.

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	data := []rune(s)
	var slow int = 0
	var fast int = -1
	t := make(map[int32]int)
	var ans = 0
	if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}
	for fast < len(s)-1 {
		fast++
		if fast-(slow)-1 > ans {
			if slow == 0 {
				ans = fast - slow
			} else {
				if data[slow] == data[fast] {
					ans = fast - slow
				} else {
					ans = fast - slow + 1
				}
			}
			// ans = fast - slow
			// if slow == 0 {
			// 	ans = fast - slow
			// } else {
			// 	ans = fast - slow + 1
			// }
		}
		if t[data[fast]] != 0 {
			// fmt.Println(slow < fast)

			slow++
			for slow < fast && data[slow+1] == data[slow] {
				slow++
			}
			// slow++
			// for data[fast] == data[slow] {
			// 	slow++
			// }
		} else {
			t[data[fast]] = t[data[fast]] + 1
		}
		// fast++

	}
	return ans
}

func main() {
	/*
		key string has max ASCII code 128
	*/
	// s := "dvdf"
	// s := "abcabcbb"
	// s := "bbbbb"
	s := "au"
	// s := "pwwkew"
	// s := "tmmzuxt"
	// s := "nfpdmpiy"
	// pwwkew
	lengthOfLongestSubstring(s)
	// fmt.Println(l1)

}
