package main

import "sort"

// "fmt"
// "strconv"

/*
Given a collection of candidate numbers (candidates) and
a target number (target),
find all unique combinations in candidates
where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.



Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
Example 2:
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

*/
func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	var comb func(nowcomb []int, targe int, index int)
	comb = func(nowcomb []int, targe, index int) {
		if targe < 0 {
			return
		}
		if targe == 0 {
			dst := make([]int, len(nowcomb))
			copy(dst, nowcomb)
			result = append(result, dst)
		}
		for i := index; i < len(candidates); i++ {
			if i != index && candidates[i] == candidates[i-1] {
				continue
			}
			nowcomb = append(nowcomb, candidates[i])
			if i == 0 || candidates[i-1] != candidates[i] {
				comb(nowcomb, targe-candidates[i], i+1)
			}
			nowcomb = nowcomb[:len(nowcomb)-1]
		}
	}
	comb([]int{}, target, 0)
	return result
}

func main() {

	inputArray := []int{10, 1, 2, 7, 6, 1, 5}
	combinationSum2(inputArray, 8)
}
