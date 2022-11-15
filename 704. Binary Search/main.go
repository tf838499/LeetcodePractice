package main

import (
// "fmt"s
// "strconv"

)

/*
easy
binary search
done

Given an array of integers nums which is sorted in ascending order,
and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
*/

// func search(nums []int, target int) int {

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == target {
// 			if nums[i] > target {
// 				return -1
// 			}
// 			return i
// 		}
// 	}
// 	return -1
// }
func search(nums []int, target int) int {
	left := 0
	rights := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		mid := (rights-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left++
		} else {
			rights--
		}
	}
	return -1
}
func main() {

	num := []int{-1, 0, 3, 5, 9, 12} // 0
	target := 9
	search(num, target)
}
