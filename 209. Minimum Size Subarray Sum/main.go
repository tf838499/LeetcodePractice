package main

import (
// "fmt"
// "strconv"
// "math"
)

/*
Given an array of positive integers nums and a positive integer target,
return the minimal length of a subarray
whose sum is greater than or equal to target.
If there is no such subarray, return 0 instead

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Input: target = 4, nums = [1,4,4]
Output: 1
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/

func minSubArrayLen(target int, nums []int) int {
	windows := []int{}
	sum := 0
	ans := 100000000000000000
	//2, 3, 1, 2, 4, 3
	for _, val := range nums {
		if target >= sum {
			sum = sum + val
			windows = append(windows, val)
		}
		for target <= sum {
			if sum >= target && ans > len(windows) {
				ans = len(windows)
			}
			sum = sum - windows[0]
			windows = windows[1:]

		}

	}
	if ans == 100000000000000000 {
		ans = 0
	}
	return ans
}

// func minSubArrayLen(s int, nums []int) int {
//     var sum int
//     var left, right int
//     minLen := len(nums)+1

//     for right < len(nums) {
//         for right < len(nums) && sum < s {
//             sum += nums[right]
//             right++
//         }
//         for left <= right && sum >= s {
//             minLen = min(minLen, right-left)
//             sum -= nums[left]
//             left++
//         }
//     }

//     if minLen > len(nums) {
//         return 0
//     }
//     return minLen
// }

// func min(x, y int) int {
//     if x < y {
//         return x
//     }
//     return y
// }
func main() {
	a := []int{1, 2, 3, 4, 5}
	// a := []int{1, 1, 1, 1, 1, 1, 1, 1}
	minSubArrayLen(15, a)
}
