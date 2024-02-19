package main

import "sort"

// "fmt"
// "strconv"

/*
Given a collection of numbers, nums,
that might contain duplicates,
return all possible unique permutations in any order.



Example 1:

Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]
Example 2:
	1 1 2
	11 2
	112
	1 12
	1 12
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	var comb func([]int, []int)
	comb = func(currentComb, candidate []int) {
		if len(candidate) == 0 {
			dst := make([]int, len(currentComb))
			copy(dst, currentComb)
			result = append(result, dst)
		}
		visited := map[int]bool{}
		for i := 0; i < len(candidate); i++ {
			if visited[candidate[i]] == false {
				visited[candidate[i]] = true
			} else {
				continue
			}
			comb(
				append(currentComb, candidate[i]),
				append(append([]int{}, candidate[:i]...), candidate[i+1:]...))

		}
	}
	sort.Ints(nums)
	comb([]int{}, nums)
	return result
}

// func permuteUnique(nums []int) [][]int {
// 	result := [][]int{}
// 	var comb func(current []int, num []int)
// 	// visited := make(map[int]bool)
// 	comb = func(current, num []int) {
// 		if len(num) == 0 {
// 			dst := make([]int, len(current))
// 			copy(dst, current)
// 			result = append(result, dst)
// 		}
// 		// visited = map[int]bool{}
// 		for i := 0; i < len(num); i++ {
// 			// if visited[i] {
// 			// 	continue
// 			// }
// 			if i != 0 && num[i] == num[i-1] {
// 				continue
// 			}

// 			comb(append(current, num[i]),
// 				append(append([]int{}, append(num[:i])...), num[i+1:]...))

// 		}
// 	}
// 	sort.Ints(nums)
// 	comb([]int{}, nums)
// 	return result
// }

// func permuteUnique(nums []int) [][]int {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	sort.Ints(nums)
// 	res := make([][]int, 0)
// 	visited := make(map[int]bool)
// 	list := make([]int, 0)
// 	backTrack(nums, visited, list, &res)
// 	return res
// }

// func backTrack(nums []int, visited map[int]bool, list []int, res *[][]int) {
// 	//back track goal, length of list equals to length of nums
// 	if len(list) == len(nums) {
// 		temp := make([]int, len(list))
// 		copy(temp, list) //copy list into temp, so it won't impact subsequent process on list
// 		*res = append(*res, temp)
// 		return
// 	}

// 	//options are each elements in nums slice
// 	for i := 0; i < len(nums); i++ {
// 		if visited[i] {
// 			continue //constrain, ignore visited element
// 		}
// 		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
// 			continue //constrain, if there is a duplication, the first element need to be visited first
// 		}
// 		list = append(list, nums[i])
// 		visited[i] = true
// 		backTrack(nums, visited, list, res)
// 		list = list[:len(list)-1]
// 		visited[i] = false
// 	}
// }
func main() {
	permuteUnique([]int{1, 1, 2})
	permuteUnique([]int{1, 2, 3})
	permuteUnique([]int{3, 3, 0, 3})
}
