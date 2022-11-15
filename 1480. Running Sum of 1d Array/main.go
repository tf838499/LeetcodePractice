package main

import (
// "fmt"s
// "strconv"
)

/*
easy
array
prefix sum
done

Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
*/
func runningSum(nums []int) []int {
	for i, j := range nums {
		if i == 0 {
			continue
		}
		nums[i] = nums[i-1] + j
	}

	return nums
}
func add(n []int, c chan []int) {
	ans := []int{}
	for j, i := range n {
		if len(ans) != 0 {
			ans = append(ans, ans[j-1]+i)
			// ans = append(ans, i)
		} else {
			ans = append(ans, i)
		}
	}
	c <- ans
}
func main() {
	num := []int{1, 2, 3, 4, 5}
	runningSum(num)
}
