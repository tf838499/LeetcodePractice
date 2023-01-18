package main

import (
// "fmt"
// "strconv"
)

/*
Given an array of distinct integers candidates and a target integer target,
return a list of all unique combinations of candidates where the chosen numbers sum to target.
You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the
frequency
of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
*/
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	// current := []int{}
	combin(candidates, []int{}, target, 0, &ans)
	return ans
}
func combin(candidate, current []int, target, add int, ans *[][]int) {
	if add == target {
		dst := make([]int, len(current))
		copy(dst, current)
		// result = append(result, dst)
		*ans = append(*ans, dst)
		return
	}
	if add > target {
		return
	}
	// idx := 0
	for i := 0; i < len(candidate); i++ {
		current = append(current, candidate[i])
		add = add + candidate[i]
		combin(candidate[i:], current, target, add, ans)
		current = current[:len(current)-1]
		add = add - candidate[i]
	}
}

func main() {

	inputArray := []int{2, 3, 6, 7}
	combinationSum(inputArray, 7)
}
