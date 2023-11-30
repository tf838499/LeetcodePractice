package main

// "fmt"s
// "strconv"

/*
Given an array nums with n objects colored red,
white, or blue, sort them in-place
so that objects of the same color are adjacent,
with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red,
white, and blue, respectively.

You must solve this problem without using the library's sort function.
*/

// func sortColors(nums []int) {
// 	if len(nums) < 2 {
// 		return
// 	}
// 	left, right := 0, 1
// 	tmp := 0
// 	for right < len(nums) {
// 		tmp = right
// 		for nums[left] > nums[right] && left >= 0 {
// 			nums[left], nums[right] = nums[right], nums[left]
// 			if left > 0 {
// 				left--
// 			}
// 			right--
// 		}
// 		right = tmp + 1
// 		left = tmp
// 	}
// }
func sortColors(nums []int) {
	for step := 0; step < len(nums)-1; step++ {
		pos := step
		for i := step + 1; i < len(nums); i++ {
			if nums[pos] > nums[i] {
				pos = i
			}
		}

		nums[pos], nums[step] = nums[step], nums[pos]
	}
}
func main() {
	num := []int{2, 2, 0, 2, 1, 1, 0}
	// 0 2
	// 022
	sortColors(num)
}
