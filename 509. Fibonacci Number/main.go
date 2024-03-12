package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
The Fibonacci numbers, commonly denoted F(n) form a sequence,
called the Fibonacci sequence,
such that each number is the sum of the two preceding ones,
starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.

Example 1:

Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
*/
// func fib(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 0 {
// 		return 0
// 	}
// 	result := fib(n-1) + fib(n-2)
// 	return result
// }
// func fib(n int) int {
// 	cache := make([]int, 2+n)

// 	cache[0] = 0
// 	cache[1] = 1

// 	for i := 2; i <= n; i++ {
// 		cache[i] = cache[i-1] + cache[i-2]
// 	}

// 	return cache[n]
// }
var cache = make(map[int]int)

func fib(n int) int {
	dparray := make([]int, n+2)
	dparray[0] = 0
	dparray[1] = 1
	for i := 2; i <= n; i++ {
		dparray[i] = dparray[i-1] + dparray[i-2]
	}
	return dparray[n]
}
func main() {
	fib(3)
}
