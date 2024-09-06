package main

import (
	"fmt"
	// "strconv"
	// "math"
)

/*
Given an integer array nums,
return true if you can partition the array
into two subsets such
that the sum of the elements in both subsets is equal
or false otherwise.

Example 1:

Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
Example 2:

Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
*/
func canPartition(nums []int) bool {
	target := 0
	for i := range nums {
		target += nums[i]
	}
	if target%2 != 0 {
		return false
	}
	target = target / 2

	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j-nums[i]] + nums[i]
			// dp[j] = max(dp[j-nums[i]]+nums[i], dp[j])
		}
	}

	return dp[target] == target
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := []int{2, 2, 1, 1}
	// input := []int{1, 2, 5}
	// input := []int{1, 2, 3, 4, 5, 6, 7}
	p := canPartition(input)
	fmt.Println(p)
}
