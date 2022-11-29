package main

import (
// "fmt"
// "strconv"
)

/*
easy
hash map
array
two pointer
done

Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution,
and you may not use the same element twice.

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Input: nums = [3,2,4], target = 6
Output: [1,2]

*/
// func twoSum(nums []int, target int) []int {
// 	t := make(map[int]int)

// 	for index, val := range nums {
// 		remainder := target - val
// 		// if index_of ,val := t[remainder]
// 		if idx_of_dual, dual_exist := t[remainder]; dual_exist {

// 			// find two numbers with sum = target
// 			return []int{idx_of_dual, index}

// 		} else {
// 			t[val] = index
// 		}
// 	}
// 	return []int{-1, -1}
// }
func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return nil
	}

	//  i < len(nums)-1 ensures we never go out of bounds for the j index
	for i := 0; i < len(nums)-1; i++ {
		// j will always be greater than i while comparing, so we start
		// at i+1
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
func main() {
	// num := []int{2, 7, 11, 15}
	// tar := 9
	// num := []int{3, 2, 4}
	// tar := 6
	num := []int{3, 3}
	tar := 6
	twoSum(num, tar)
}
