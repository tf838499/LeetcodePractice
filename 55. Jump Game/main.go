package main

/*
You are given an integer array nums.
You are initially positioned at the array's first index,
and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
*/

// func canJump(nums []int) bool {
// 	step := 0
// 	for i := 0; i < len(nums); {
// 		step--
// 		if step < nums[i] {
// 			step = nums[i]
// 		}
// 		i++
// 		if i >= len(nums) {
// 			return true
// 		}
// 		if step == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
// func canJump(nums []int) bool {
// 	if len(nums) <= 1 {
// 		return true
// 	}
// 	step := 0
// 	for i := 0; i < len(nums); {
// 		step = nums[i]
// 		if step == 0 {
// 			for nums[i] <= step {
// 				step++
// 				i--
// 				if i < 0 {
// 					return false
// 				}
// 			}
// 			if nums[i] > step {
// 				step = nums[i]
// 			}
// 		}
// 		i = i + step
// 		if i+1 >= len(nums) {
// 			return true
// 		}

// 	}
// 	return true
// }
func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
}
func main() {
	// num := []int{2, 3, 1, 1, 4}
	// num := []int{3, 2, 1, 0, 4}
	// num := []int{3, 2, 1, 3, 0, 0, 0, 1}
	num := []int{2, 5, 0, 0}
	// num := []int{2, 0, 0}
	canJump(num)
}
