package main

import (
// "fmt"s
// "strconv"

)

/*
Given an array nums with n objects colored red,
white, or blue, sort them in-place so that objects of the same color are adjacent,
with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red,
white, and blue, respectively.

You must solve this problem without using the library's sort function.
*/

// func sortColors(nums []int) {
// 	s, f := 0, 1
// 	for f > s {
// 		if nums[f] >= nums[s] {
// 			f++
// 			s++
// 		}
// 		for nums[f] <= nums[s] {
// 			nums[s], nums[f] = nums[f], nums[s]

// 		}
// 	}
// }
func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	index := 0

	for index <= right {
		switch nums[index] {
		case 0:
			nums[left], nums[index] = nums[index], nums[left]
			left++
			index++
		case 2:
			nums[right], nums[index] = nums[index], nums[right]
			right--
		default:
			index++
		}
	}
}

func main() {
	num := []int{2, 2, 0, 2, 1, 1, 0}
	sortColors(num)
}
