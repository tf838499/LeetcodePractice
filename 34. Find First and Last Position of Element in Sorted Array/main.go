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
func findIdx(nums []int, target int, firstIdx bool) int {
    lo, hi, idx := 0, len(nums)-1, -1
    for lo <= hi {
        mid := (lo+hi)>>1
        if target == nums[mid] {
            idx = mid 
            if firstIdx {
                hi = mid - 1
            } else {
                lo = mid + 1
            }
        } else if target < nums[mid] {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    return idx
}


func searchRange(nums []int, target int) []int {
    return []int{findIdx(nums, target, true), findIdx(nums, target, false)}
}
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	// l := 0
	// l_r, r_l := len(nums)/2, len(nums)/2
	l_r, r_l := len(nums)-1, 0
	// l_r := len(nums) - 1
	// r := len(nums) - 1
	// r_l := 0
	for l < r {
		if nums[l] < target {
			l = (l_r-l)/2 + l
		} else if nums[l] > target {
			l_r = l
			l = l_r - (l_r)/2

		}

		if nums[r] > target {
			r = (r-r_l)/2 + r_l
		} else if nums[r] < target {
			r_l = r
			r = (len(nums)-1-r_l)/2 + r_l
		}
		if nums[r] == target && nums[r+1] != target && {
			return []int{l, r}
		}
	}

	return []int{-1, -1}
}
func main() {
	// num := []int{5, 5, 6, 6, 7, 7, 7, 7, 7, 7, 7, 9, 9}
	num := []int{7, 8}
	searchRange(num, 8)
}
