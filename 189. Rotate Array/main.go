package main

import (
// "fmt"5
// // "strconv"
// "sort"
)

/*
medium
two pointer

解法: 
解法很多 可以用切片的方式重新合併矩陣 例如移動3格 將3格後放置一矩陣 3個前放置另一矩陣 在合併
抑或是重新賦一新陣列，在依序將舊有排入
Given an array, rotate the array to the right by k steps, where k is non-negative.
tmp = 7


Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/

// func rotate(nums []int, k int) {
// 	fast, slow := k, 0
// 	for ; slow < len(nums); slow++ {
// 		nums[fast], nums[slow] = nums[slow], nums[fast]
// 		fast++
// 		slow++
// 		// if fast > len(nums)-1 {
// 		// 	fast = 0
// 		// } else if slow > len(nums)-1 {
// 		// 	slow = 0
// 		// }
// 	}
// 	fast++

// }

unc rotate(nums []int, k int)  {
    if len(nums) != 0 {
		copy(nums, append(nums[len(nums) - k % len(nums): ], nums[0:len(nums) - k % len(nums)]...))
	}
}
func main() {

	input := []int{1, 2, 3, 4, 5, 6, 7} // 5,6,7,1,2,3,4
	step := 3
	rotate(input, step)
}
