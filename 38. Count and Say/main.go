package main

import (
	// "fmt"
	"strconv"
)

/*
Medium
Dynamic Programming
Recursion
sliding windows
doing

The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1),
which is then converted into a different digit string.
To determine how you "say" a digit string, split it into the minimal number of substrings such that each substring contains exactly one unique digit.
Then for each substring, say the number of digits, then say the digit. Finally, concatenate every said digit.

For example, the saying and conversion for digit string "3322251":

Input: n = 1
Output: "1"
Explanation: This is the base case.

Input: n = 4
Output: "1211"
Explanation:
countAndSay(1) = "1"								1
countAndSay(2) = say "1" = one 1 = "11"				11
countAndSay(3) = say "11" = two 1's = "21"			111
countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"	1 11 21 1211
1211  1個1，1個2，2個1  111221
*/
// func countAndSay(n int) string{
// 	if n == 1{
// 		return "1"
// 	}

// }
func countAndSay(n int) string {
	tmp := []byte{}
	if n == 1 {
		return "1"
	}
	for i := 0; i < n; i++ {
		if len(tmp) == 0 {
			tmp = append(tmp, []byte(strconv.Itoa(1))...)

			continue
		}

		num := 1
		tmp1 := []byte{}
		for j := 0; j <= len(tmp)-1; j++ {
			if j+1 >= len(tmp) {
				tmp1 = append(tmp1, []byte(strconv.Itoa(num))...)
				tmp1 = append(tmp1, tmp[j])
				break
			}
			if tmp[j] != tmp[j+1] {
				tmp1 = append(tmp1, []byte(strconv.Itoa(num))...)
				tmp1 = append(tmp1, tmp[j])
				num = 1
			} else {
				num++
			}

		}
		// rep:
		tmp = tmp1
	}

	return string(tmp)
}

// func countAndSay(n int) string {
// 	if n == 1 {
// 		return "1"
// 	}
// 	previous := countAndSay(n - 1)
// 	result, length := []byte{}, 1
// 	for i := 0; i < len(previous)-1; i++ {
// 		if previous[i] != previous[i+1] {
// 			result = append(result, []byte(strconv.Itoa(length))...)
// 			result = append(result, previous[i])
// 			length = 1
// 		} else {
// 			length++
// 		}
// 	}
// 	result = append(result, []byte(strconv.Itoa(length))...)
// 	result = append(result, previous[len(previous)-1])
// 	return string(result)
// }
func main() {
	// hay := "mississippi"
	// nee := "issipi"
	// var num = []int{1}
	// var num = []int{4, 5}
	n := 4

	// words := ["bella","label","roller"]
	countAndSay(n)
}
