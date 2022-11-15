package main

import (
// "fmt"s
// "strconv"

)

/*
Medium
tree
Dynamic Programming
Divide and conquer
doing


** 注意 小於零的 越接近零 越大
Given an integer array nums, find the
subarray
which has the largest sum and return its sum.

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Input: nums = [1]
Output: 1

Input: nums = [5,4,-1,7,8]
Output: 23
*/

func maxSubArray(nums []int) int {
	max := ^int(^uint(0) >> 1)
	tmp := 0
	slow := 0
	fast := 0
	for fast < len(nums) && slow < len(nums) {

		tmp = tmp + nums[fast]

		if tmp > max {
			max = tmp
		}
		for tmp < 0 {
			tmp -= nums[slow]
			slow++
			continue
		}
		fast++

	}
	return max
}

// func maxSubArray(nums []int) int {
//     return findSubArray(nums, 0, len(nums) - 1)
// }

// func findCrossSubArray(nums []int, left int, right int ) int {

//     mid := left + (right - left) / 2
//     leftSum := math.MinInt32
//     rightSum := math.MinInt32
//     sum := 0

//     for i:=mid; i >= left; i-- {
//         sum += nums[i]
//         if sum > leftSum {
//             leftSum = sum
//         }
//     }

//     sum = 0
//     for i:=mid+1; i<= right; i++ {
//         sum += nums[i]
//         if sum > rightSum {
//             rightSum = sum
//         }
//     }

//     return leftSum + rightSum
// }

// func findSubArray(nums []int, left int, right int ) int {
//     if left == right {
//         return nums[left]
//     }

//     mid := left + (right - left) / 2

//     a := findSubArray(nums, left, mid)
//     b := findSubArray(nums, mid+1, right)
//     c := findCrossSubArray(nums, left, right)

//     if a > b {
//         if a > c {
//             return a
//         } else {
//             return c
//         }
//     } else {
//         if b > c {
//             return b
//         } else {
//             return c
//         }
//     }
// }
func main() {

	num := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSubArray(num)
}
