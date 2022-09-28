package main

import (
	"fmt"
	// "strconv"
)

/*
easy
string
done

Given a string s and an integer k, reverse the first k characters
for every 2k characters counting from the start of the string.

If there are fewer than k characters left, reverse all of them.
If there are less than 2k but greater than or equal to k characters,
then reverse the first k characters and leave the other as original.

Input: s = "abcdefg", k = 2
Output: "bacdfeg"

Input: s = "abcd", k = 2
Output: "bacd"
*/

func reverseStr(s string, k int) string {
	ans := []byte(s)
	// var stri string
	// "cbadef"

	for i := 0; i < len(ans)-1; {
		// ans[i], ans[i+1] = ans[i+1], ans[i]
		// tmp := ans[i : i+k]
		if len(ans) < k {
			k = len(ans)
		} else if len(ans) < k+i {
			k = len(ans) - i
		}
		for j := 0; j < k/2; j++ {
			ans[i+j], ans[i+(k-j-1)] = ans[i+(k-j-1)], ans[i+j]
		}
		i = i + 2*k
		// k = i + k
	}

	// stri = reverseStr(string(ans[2*k:]), k)
	// return stri

	// for i := 0; i < k/2; i++ {
	// 	ans[i], ans[k-(i+1)] = ans[k-(i+1)], ans[i]
	fmt.Println(string(ans[:]))
	// }
	return string(ans[:])
}

func main() {
	words := "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
	// words := "abcdefg"
	// words := []byte{"H", "a", "n", "n", "a", "h"}
	// fmt.Print(len(words))
	// words := ["bella","label","roller"]
	reverseStr(words, 39)

}

// "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
