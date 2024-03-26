package main

// "strconv"

/*
Given an integer n,
break it into the sum of k positive integers,
where k >= 2,
and maximize the product of those integers.

Return the maximum product you can get.



Example 1:

Input: n = 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.
Example 2:

Input: n = 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
Constraints:

2 <= n <= 58
*/
// 2
func integerBreak(n int) int {
	10
	help(n/2, n/2)
}
func help(left, right int) int {
	lef
	if left > 2 {
		value := help(left/2, left/2)
		
	}
	if right > 2 {
		help(right/2, right/2)
	}

}
func main() {
	integerBreak(10)
}
