package main

/*
Given an integer array nums of unique elements,
return all possible subsets
(the power set).

The solution set must not contain duplicate subsets.
Return the solution in any order.



Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [0]
Output: [[],[0]]

*/

func subsets(nums []int) [][]int {
	result := [][]int{}
	var comb func(sub []int, index int)
	comb = func(sub []int, index int) {
		if index > len(nums) {
			return
		}
		dst := make([]int, len(sub))
		copy(dst, sub)
		result = append(result, dst)
		for i := index; i < len(nums); i++ {
			sub = append(sub, nums[i])
			comb(sub, i+1)
			sub = sub[:len(sub)-1]
		}
	}
	comb([]int{}, 0)
	return result
}

// func subsets(nums []int) [][]int {
// 	var queue [][]int
// 	queue = append(queue, []int{})

// 	for i := 0; i < len(nums); i++ {
// 		length := len(queue)

// 		for j := 0; j < length; j++ {
// 			item := queue[j]

// 			tmp := make([]int, len(item))
// 			copy(tmp, item)
// 			tmp = append(tmp, nums[i])

// 			queue = append(queue, tmp)
// 		}
// 	}

// 	return queue
// }
// func subsets(nums []int) [][]int {
// 	return append([][]int{{}}, recursiveSubsets(nums)...)
// }

// func recursiveSubsets(nums []int) [][]int {
// 	if len(nums) == 0 {
// 		return nil
// 	}

// 	allSubsets := [][]int{}
// 	for i, v := range nums {
// 		subs := recursiveSubsets(nums[i+1:])
// 		for i := 0; i < len(subs); i++ {
// 			subs[i] = append([]int{v}, subs[i]...)
// 		}
// 		subs = append([][]int{{v}}, subs...)
// 		allSubsets = append(allSubsets, subs...)
// 	}

// 	return allSubsets
// }
func main() {
	// num := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	subsets([]int{1, 2, 3})
}
