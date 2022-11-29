package main

import (
	"fmt"
	// "strconv"
)

/*
easy
hashmap
done



Given a string s,
reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.


Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Input: s = "God Ding"
Output: "doG gniD"
*/
func reverseWords(s string) string {

	runeS := []byte(s)

	record := []int{}
	blank := rune(' ')
	for j, i := range s {
		if i == blank {
			record = append(record, j)
		}
	}
	// record[len(runeS)] = 1
	record = append(record, len(s))
	l := 0
	for i := range record {
		// fmt.Println(j)
		j := record[i]
		r := j - 1
		for l < r {
			runeS[l], runeS[r] = runeS[r], runeS[l]
			l++
			r--
		}
		l = j + 1
	}
	// fmt.Println(string(runeS))
	return string(runeS)
}

// func reverseWords(s string) string {
// 	i, j := 0, 0
// 	res := make([]byte, len(s))

// 	for j < len(s) {
// 		if s[j] == ' ' {
// 			if i == j {
// 				res[j] = s[j]
// 			} else {
// 				reverseCopy(s, res, i, j-1)
// 				res[j] = s[j]
// 			}
// 			i, j = j+1, j+1
// 			continue
// 		}
// 		j++
// 	}

// 	if i < j {
// 		reverseCopy(s, res, i, j-1)
// 	}
// 	return string(res)
// }

// func reverseCopy(s string, res []byte, start, end int) {
// 	for start < end {
// 		res[start], res[end] = s[end], s[start]
// 		start, end = start+1, end-1
// 	}
// 	if start == end {
// 		res[start] = s[start]
// 	}
// }
func main() {
	s := "Let's take LeetCode contest"
	// s := "as"
	// s := "I love you"
	p := reverseWords(s)
	fmt.Println(p)
}
