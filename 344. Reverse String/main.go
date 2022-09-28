package main

import (
	"fmt"
	// "strconv"
)

/*
easy
two pointer
done

Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/
func reverseString(s []byte) {

	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-(i+1)] = s[len(s)-i-1], s[i]
	}
}

// func reverseString(s []byte) {

// 	for i := 0; i < len(s)/2; i++ {
// 		tmp := s[i]
// 		s[i],s[len(s)-i-1] = s[len(s)-i-1],s[i]
// 		s[i] = s[len(s)-i-1]
// 		s[len(s)-i-1] = tmp
// 		fmt.Println(s[i])
// 	}
// }
func main() {
	words := "Hannah"
	// words := []byte{"H", "a", "n", "n", "a", "h"}

	// words := ["bella","label","roller"]
	reverseString([]byte(words))

}
