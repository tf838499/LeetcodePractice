package main

import (
	"fmt"
	// "strconv"
)

/*
Given an integer n,
return true if it is a power of three.
Otherwise, return false.

An integer n is a power of three,
if there exists an integer x such that n == 3x.

Input: n = 27
Output: true
Explanation: 27 = 33
27 /3
Input: n = 0
Output: false
Explanation: There is no x where 3x = 0.
*/
func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 {
		return false
	}
	ans := false
	if n%3 == 0 {
		ans = isPowerOfThree(n / 3)
	}
	return ans
}

func main() {

	// words := []byte{"H", "a", "n", "n", "a", "h"}

	// words := ["bella","label","roller"]
	p := isPowerOfThree(27)
	fmt.Println(p)
}
