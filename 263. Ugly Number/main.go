package main

import (
	"fmt"
	// "strconv"
)

/*
day 2 筆記


An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return true if n is an ugly number.

Input: n = 6
Output: true
Explanation: 6 = 2 × 3

Input: n = 1
Output: true
Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.

Input: n = 14
Output: false
Explanation: 14 is not ugly since it includes the prime factor 7.

*/
func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	ugly := make(map[int]int)
	ugly[2] = 2
	ugly[3] = 3
	ugly[5] = 5
	ans := true
	for n != 0 && n != 1 {
		flag := 0
		for _, j := range ugly {

			if n%j == 0 {
				n = n / j
				break
			} else {
				flag++
			}
		}
		if flag == 3 {
			ans = false
			break
		}
	}
	return ans
}

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	if num <= 6 {
		return true
	}

	if num%2 == 0 || num%3 == 0 || num%5 == 0 {
		for num%2 == 0 {
			num = num / 2
		}
		for num%3 == 0 {
			num = num / 3
		}
		for num%5 == 0 {
			num = num / 5
		}
		if num != 1 {
			return false
		}
		return true
	}

	return false
}
func main() {
	s := 0

	p := isUgly(s)
	fmt.Println(p)
}
