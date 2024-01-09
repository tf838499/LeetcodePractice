package main

// "fmt"
// "strconv"

/*
Given an array of distinct integers candidates and a target integer target,
return a list of all unique combinations of candidates
where the chosen numbers sum to target.
You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the frequency of at least one of the chosen numbers is different.

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
	var result [][]int
	// current := []int{}
	var comb func(cand, curComb []int, ind, sumtarget int)
	comb = func(cand, curComb []int, ind, sumtarget int) {
		if sumtarget == 0 {
			dst := make([]int, len(curComb))
			copy(dst, curComb)
			result = append(result, dst)
			return
		}
		if sumtarget < 0 {
			return
		}
		for i := ind; i < len(cand); i++ {
			curComb = append(curComb, cand[i])
			comb(cand, curComb, i, sumtarget-cand[i])
			// curComb = curComb[1:]
			curComb = curComb[:len(curComb)-1]
		}

	}
	comb(candidates, []int{}, 0, target)
	return result
}

// 3
//3 3
//3	5
//3
// func combinationSum(candidates []int, target int) [][]int {
// 	var ans [][]int

// 	if len(candidates) == 0 {
// 		return ans
// 	}

// 	bt(&ans, make([]int, 0), candidates, 0, target)

// 	return ans
// }

// func bt(ans *[][]int, tmp, nums []int, idx, tgt int) {
// 	if tgt < 0 || idx > len(nums) {
// 		return
// 	}

// 	if tgt == 0 {
// 		cpyTmp := make([]int, len(tmp))
// 		copy(cpyTmp, tmp)
// 		*ans = append(*ans, cpyTmp)
// 	}

// 	for i := idx; i < len(nums); i++ {
// 		tmp = append(tmp, nums[i])
// 		bt(ans, tmp, nums, i, tgt-nums[i])
// 		tmp = tmp[:len(tmp)-1]
// 	}
// }
func main() {

	inputArray := []int{3, 5, 8}
	combinationSum(inputArray, 11)
}
