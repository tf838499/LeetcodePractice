package main

import (
// "fmt"

)

/*
easy
binary search
done

Given a sorted array of distinct integers and a target value,
return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Input: nums = [1,3,5,6], target = 5
Output: 2
Input: nums = [1,3,5,6], target = 2
Output: 1
Input: nums = [1,3,5,6], target = 7
Output: 4
*/

func searchInsert(nums []int, target int) int {
	var l, r = 0, len(nums)
	for l < r {
		m := (r + l) >> 1
		// m := (r-l)/2 + l
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}

	}
	return l

}

func main() {
	nums := []int{1, 3, 5, 6}
	searchInsert(nums, 2)
}
