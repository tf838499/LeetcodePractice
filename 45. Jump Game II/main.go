package main

// "fmt"
// "strconv"

/*
You are given a 0-indexed array of integers nums of length n.
You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump
from index i. In other words,
if you are at nums[i], you can jump to any nums[i + j] where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1].
The test cases are generated such that you can reach nums[n - 1].



Example 1:

5-2 3

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [2,3,0,1,4]
Output: 2
*/
// 记录步骤规则：每超过上一次可达最大范围，需要跳跃一次，次数+1
// 记录位置：i == lastDistance + 1
// func jump(nums []int) int {
// 	// 根据题目规则，初始位置为nums[0]
// 	lastDistance := 0 // 上一次覆盖范围
// 	curDistance := 0  // 当前覆盖范围（可达最大范围）
// 	minStep := 0      // 记录最少跳跃次数

// 	for i := 0; i < len(nums); i++ {
// 		if i == lastDistance+1 { // 在上一次可达范围+1的位置，记录步骤
// 			minStep++                  // 跳跃次数+1
// 			lastDistance = curDistance // 记录时才可以更新
// 		}
// 		curDistance = max(nums[i]+i, curDistance) // 更新当前可达的最大范围
// 	}
// 	return minStep
// }
// func max(a, b int) int {
// 	if a >= b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
func jump(nums []int) int {

	count := 0
	max := len(nums) - 1

	for max != 0 {
		for i := 0; i < max; i++ {
			if i+nums[i] >= max {
				max = i
				count++
				break
			}
		}
	}
	return count
}

// func jump(nums []int) int {
// 	if len(nums) == 1 {
// 		return 0
// 	}
// 	min := 1
// 	Isnew := false
// 	nowStep := nums[0] - 1
// 	for i := 1; i < len(nums); i++ {
// 		if nowStep == len(nums)-i-1 {
// 			return min
// 		}
// 		nowStep, Isnew = max(nowStep, nums[i])
// 		if nowStep == 0 {
// 			return 0
// 		}
// 		if Isnew {
// 			min++
// 		}
// 		nowStep--
// 	}
// 	return min
// }
// func max(a, b int) (int, bool) {
// 	if a >= b {
// 		return a, false
// 	} else {
// 		return b, true
// 	}
// }
func main() {
	input := []int{1, 1, 1, 1, 1, 0}
	jump(input)
}
