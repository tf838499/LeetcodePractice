package main

// "fmt"
// "strconv"

/*
Given an array nums of distinct integers,
return all the possible permutations.
You can return the answer in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]
*/
func permute(nums []int) [][]int {
	result := [][]int{}
	var comb func([]int, []int)
	comb = func(candidate, currComb []int) {
		if len(candidate) == 0 {
			dst := make([]int, len(currComb))
			copy(dst, currComb)
			result = append(result, dst)
		}
		for i := 0; i < len(candidate); i++ {
			// currComb = append(currComb, candidate[i])
			comb(append(append([]int{}, candidate[:i]...), candidate[i+1:]...), append(currComb, candidate[i]))
		}
	}
	comb(nums, []int{})
	return result
}

// func permute(nums []int) [][]int {
// 	result := [][]int{}
// 	var comb func([]int, []int)
// 	comb = func(num, cand []int) {
// 		if len(num) == 0 {
// 			dst := make([]int, len(cand))
// 			copy(dst, cand)
// 			result = append(result, dst)
// 			return
// 		}
// 		// for i := 0; i < len(num); i++ {
// 		// 	comb(append(append([]int{}, num[:i]...), num[i+1:]...), append(cand, num[i]))
// 		// }
// 		for i := 0; i < len(num); i++ {

// 			cand = append(cand, num[i])
// 			nextnum := append([]int{}, num[:i]...)
// 			nextnum = append(nextnum, num[i+1:]...)
// 			// append(append([]int{}, left[:idx]...), left[idx+1:]...)
// 			comb(nextnum, cand)
// 			cand = cand[:len(cand)-1]
// 		}
// 	}
// 	comb(nums, []int{})
// 	return result
// }

func main() {
	// permute([]int{5, 4, 6, 2})
	permute([]int{0, 1})
	// permute([]int{1, 2, 3})
}
