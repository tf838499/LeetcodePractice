package main

import (
// "fmt"
// "strconv"
)

/*
easy
hashmap
done
day 1筆記  11/16
可以很簡單 另外使用index 以目前找到不為零的數量，
例如 找到兩個Index = 2 ，在將這兩個數字放回數組中
找完後 index後面的則全改為0，因為已找完不為零的數值
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Input: nums = [0]
Output: [0]

*/
func moveZeroes(nums []int) {
	slow, fast := 0, 1
	numsLen := len(nums)
	if numsLen == 0 || numsLen == 1 {
		return
	}
	for fast < numsLen {
		for nums[fast] == 0 {
			fast++
			if fast >= numsLen {
				return
			}
		}
		for nums[slow] != 0 {
			if slow == fast {
				fast++
			}
			if fast >= numsLen {
				return
			}
			slow++
		}

		nums[fast], nums[slow] = nums[slow], nums[fast]
	}
}
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	insertIndex := 0
	for _, e := range nums {
		if e != 0 {
			nums[insertIndex] = e
			insertIndex++
		}
	}

	for ; insertIndex < len(nums); insertIndex++ {
		nums[insertIndex] = 0
	}
}
func main() {
	s := []int{0, 1, 0, 3, 12}
	// s := []int{1, 2, 3, 4, 5}
	// s := []int{0, 0, 0, 0, 0}
	moveZeroes(s)

}
