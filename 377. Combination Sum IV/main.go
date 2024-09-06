package main

// "fmt"
// "strconv"
// "strings"
// "math"

/*
如果求组合数就是外层for循环遍历物品，内层for遍历背包。

如果求排列数就是外层for遍历背包，内层for循环遍历物品。
Given an array of distinct integers nums and a target integer target,
return the number of possible combinations that add up to target.

The test cases are generated so that the answer can fit in a 32-bit integer.

Input: nums = [1,2,3], target = 4
Output: 7
Explanation:
The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
Note that different sequences are counted as different combinations.
*/

// func combinationSum4(nums []int, target int) int {
// 	dp := make([]int, target+1)
// 	dp[0] = 1
// 	for j := 0; j <= target; j++ {
// 		for i := 0; i < len(nums); i++ {
// 			if j >= nums[i] {
// 				dp[j] = dp[j] + dp[j-nums[i]]
// 			}

// 		}
// 	}
// 	return dp[target]
// }
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j >= nums[i] {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}
func main() {
	target := 4
	nums := []int{1, 2, 3}
	combinationSum4(nums, target)
}
