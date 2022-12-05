package main

import (
// "fmt"s
// "strconv"
)

/*
Given an array of integers nums sorted in non-decreasing order,
find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Input: nums = [], target = 0
Output: [-1,-1]
*/
// func findIdx(nums []int, target int, firstIdx bool) int {
//     lo, hi, idx := 0, len(nums)-1, -1
//     for lo <= hi {
//         mid := (lo+hi)>>1
//         if target == nums[mid] {
//             idx = mid
//             if firstIdx {
//                 hi = mid - 1
//             } else {
//                 lo = mid + 1
//             }
//         } else if target < nums[mid] {
//             hi = mid - 1
//         } else {
//             lo = mid + 1
//         }
//     }
//     return idx
// }

// func searchRange(nums []int, target int) []int {
//     return []int{findIdx(nums, target, true), findIdx(nums, target, false)}
// }
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}
	lenNums := len(nums) - 1
	l, r := 0, lenNums
	L, R := 0, lenNums
	idx, Idx := -1, -1
	for l <= r || L <= R {
		mid := (r-l)/2 + l
		if nums[mid] == target && l <= r {
			idx = mid
			l = (mid-l)/2 + l + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}

		Mid := (R-L)/2 + L
		if nums[Mid] == target && L <= R {
			Idx = Mid
			R = (R-Mid)/2 + Mid - 1
		} else if nums[Mid] < target {
			L = Mid + 1
		} else {
			R = Mid - 1
		}

	}

	return []int{Idx, idx}
}
func main() {
	// num := []int{5, 5, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 9}
	// num := []int{1}
	// num := []int{1, 3}
	num := []int{2, 2}
	searchRange(num, 2)
}
