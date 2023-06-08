package main

// "fmt"5
// // "strconv"
// "sort"

/*
easy
two pointer

解: 重點>以排列過(sorted)，並且是非降序排列，中間的數值必比頭尾小! 使用two poinet 比較最前跟最後平方後誰大就放到隊伍尾即可得解。
Given an integer array nums sorted in non-decreasing order,
return an array of the squares of each number sorted in non-decreasing order.

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

*/

// func sortedSquares(nums []int) []int {
// 	numsLen := len(nums)
// 	l, r := 0, numsLen-1
// 	result := make([]int, numsLen)
// 	for i := numsLen - 1; i >= 0; i-- {
// 		lVal := nums[l] * nums[l]
// 		rVal := nums[r] * nums[r]
// 		if rVal > lVal {
// 			result[i] = rVal
// 			r--
// 		} else {
// 			result[i] = lVal
// 			l++
// 		}
// 	}

//		return result
//	}
func sortedSquares(nums []int) []int {

}
func main() {

	// input := []int{-7, -3, 2, 3, 11}
	input := []int{-4, -1, 0, 3, 10}
	sortedSquares(input)
}
