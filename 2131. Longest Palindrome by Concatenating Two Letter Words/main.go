package main

import (
	"fmt"
	// "strconv"
)

/*
Medium
string
hash table
greedy
doing

You are given an array of strings words. Each element of words consists of two lowercase English letters.

Create the longest possible palindrome by selecting some elements from words and concatenating them in any order. Each element can be selected at most once.

Return the length of the longest palindrome that you can create. If it is impossible to create any palindrome, return 0.

A palindrome is a string that reads the same forward and backward.

Input: words = ["lc","cl","gg"]
Output: 6
Explanation: One longest palindrome is "lc" + "gg" + "cl" = "lcggcl", of length 6.

["ab","ty","yt","lc","cl","ab"]
Output: 8
Explanation: One longest palindrome is "ty" + "lc" + "cl" + "yt" = "tylcclyt", of length 8.
Note that "lcyttycl" is another longest palindrome that can be created.
words = ["cc","ll","xx"]
Note that "clgglc" is another longest palindrome that can be created.
"ab","ty","yt","ll","cc"
tyllllccyt
*/

func longestPalindrome(words []string) int {

	n := make(map[string]int)
	var ans int

	fmt.Println(n)
	tmp := ""
	for _, str := range words {
		p := []byte(str)
		var reverse string
		endIndex := len(p) - 1

		if p[endIndex] == p[0] {
			if tmp == "" {
				tmp = str
				ans = ans + 2
			}
			if n[str] == 0 {
				n[str]++
			} else if tmp != str && n[str] > 1 {
				ans = ans + 4
			} else {
				ans = ans + 2
			}
			continue
		}

		for i := endIndex; i >= 0; i-- {
			reverse = reverse + string(p[i])
		}

		if n[reverse] == 0 {
			n[str]++
		} else {
			ans = ans + 4
			// n[str]--s
			n[reverse]--
		}

	}

	return ans
	// for i,j in words
}

// ["ab","ty","yt","lc","cl","ab"]
// ddaabbdd ddddaaaa aa
func main() {
	foo := []string{"ll", "lb", "bb", "bx", "xx", "lx", "xx", "lx", "ll", "xb", "bx", "lb", "bb", "lb", "bl", "bb", "bx", "xl", "lb", "xx"}
	longestPalindrome(foo)
}
