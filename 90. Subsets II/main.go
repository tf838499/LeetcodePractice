package main

import "sort"

// "fmt"
// "strconv"

/*
Given an integer array nums that may contain duplicates,
return all possible
subsets (the power set).

The solution set must not contain duplicate subsets.
Return the solution in any order.



Example 1:
			2,
Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
Example 2:

Input: nums = [0]
Output: [[],[0]]
*/
func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	var comb func(index int, path []int)
	comb = func(index int, path []int) {
		if index > len(nums) {
			return
		}
		dst := make([]int, len(path))
		copy(dst, path)
		result = append(result, dst)
		for i := index; i < len(nums); i++ {
			if i != index && nums[i] == nums[i-1] {
				continue
				// path = append(path, nums[i])
			}
			path = append(path, nums[i])
			comb(i+1, path)
			path = path[:len(path)-1]
		}
	}
	comb(0, []int{})
	return result
}
func main() {
	subsetsWithDup([]int{1, 2, 2, 3})

}
