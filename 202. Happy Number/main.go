package main

// "fmt"
// "strconv"

/*
Easy
Done

Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay),
or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.
Input: n = 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
Input: n = 2
Output: false
*/
// func isHappy(n int) bool {
// 	set := make(map[int]bool)
// 	var remainder int
// 	var quotient int
// 	for n != 1 {
// 		remainder = n % 10
// 		quotient = n / 10
// 		n = remainder * remainder
// 		for quotient != 0 {
// 			remainder = quotient % 10
// 			quotient = quotient / 10
// 			n += remainder * remainder
// 		}
// 		if set[n] == true {
// 			return false
// 		} else {
// 			set[n] = true
// 		}

//		}
//		return true
//	}
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true
		remainder := n % 10
		quotient := n / 10
		remainder = remainder * remainder
		n = remainder

		for quotient != 0 {

			remainder = quotient % 10
			quotient = quotient / 10
			remainder = remainder * remainder
			n = n + remainder
		}
	}
	return n == 1
}

// func isHappy(n int) bool {
// 	m := make(map[int]bool)
// 	for n != 1 && !m[n] {
// 		n, m[n] = getSum(n), true
// 	}
// 	return n == 1
// }

//	func getSum(n int) int {
//		sum := 0
//		for n > 0 {
//			sum += (n % 10) * (n % 10)
//			n = n / 10
//		}
//		return sum
//	}
func main() {
	// numtest1_1 := 19
	numtest1_2 := 2

	// numtest2_1 := []int{4, 9, 5}
	// numtest2_2 := []int{9, 4, 9, 8, 4}
	// words := ["bella","label","roller"]
	isHappy(numtest1_2)

}
