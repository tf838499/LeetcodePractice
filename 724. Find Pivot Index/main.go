package main

import (
// "fmt"s
// "strconv"

)

/*
easy
prefix sum
Dynamic Programming
two pointers
array
doing

Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.
Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11

*/
// func pivotIndex(nums []int) int {
// 	// flag := 0
// 	sum := 0
// 	for i := range nums {
// 		sum += nums[i]
// 	}
// 	left := 0
// 	for i := len(nums) - 1; i > 0; i-- {

// 		// tmp := sum - left - nums[i-1]
// 		// tmp := nums[i]
// 		sum = sum - left
// 		// sum = sum + nums[i]
// 		if sum == left {
// 			return i
// 		}
// 		sum = sum + nums[i]
// 		left += nums[i]
// 	}
// 	return -1
// }

func accumulation(nums []int) int {

	summation := 0

	for _, num := range nums {
		summation += num
	}

	return summation
}

func pivotIndex(nums []int) int {

	// Initialization
	// Left hand side be empty
	// Right hand side holds all weights
	totalWeightOnLeft := 0
	totalWeightOnRight := accumulation(nums)

	for idx, currentWeight := range nums {

		totalWeightOnRight -= currentWeight

		if totalWeightOnLeft == totalWeightOnRight {
			// balance is met on both sides
			return idx
		}

		totalWeightOnLeft += currentWeight
	}

	return -1

}
func main() {

	num := []int{-1, -1, -1, -1, -1, -1} // 0
	pivotIndex(num)
}
