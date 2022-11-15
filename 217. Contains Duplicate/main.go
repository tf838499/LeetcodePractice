package main

import (
// "fmt"
// "strconv"

)

/*
easy
array
hash map
done

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct

Input: nums = [1,2,3,1]
Output: true
Input: nums = [1,2,3,4]
Output: false
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

*/

func containsDuplicate(nums []int) bool {
	stack := make(map[int]int)
	for _, i := range nums {
		if stack[i] != 0 {
			return true
		}
		stack[i] = 1
	}
	return false
}

func main() {

	num := []int{1, 2, 3, 1}
	containsDuplicate(num)
}
