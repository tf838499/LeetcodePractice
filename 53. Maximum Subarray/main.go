package main

// "fmt"s
// "strconv"

/*
Medium
tree
Dynamic Programming
Divide and conquer
doing


** 注意 小於零的 越接近零 越大
Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Input: nums = [1]
Output: 1

Input: nums = [5,4,-1,7,8]
Output: 23
*/

// func maxSubArray(nums []int) int {
// 	slow := 0
// 	// fast := 0
// 	MaxTotal := 0
// 	tmp := 0
// 	for i := 0; i < len(nums); i++ {
// 		tmp = tmp + nums[i]
// 		if MaxTotal < tmp {
// 			MaxTotal = tmp
// 		} else if tmp < MaxTotal {
// 			tmp = tmp - nums[slow]
// 			slow++
// 		}

// 	}
// 	return MaxTotal
// }
// func maxSubArray(nums []int) int {
// 	n := len(nums)
// 	// 这里的dp[i] 表示，最大的连续子数组和，包含num[i] 元素
// 	dp := make([]int, n)
// 	// 初始化，由于dp 状态转移方程依赖dp[0]
// 	dp[0] = nums[0]
// 	// 初始化最大的和
// 	mx := nums[0]
// 	for i := 1; i < n; i++ {
// 		// 这里的状态转移方程就是：求最大和
// 		// 会面临2种情况，一个是带前面的和，一个是不带前面的和
// 		//num := []int{5, 4, -1, 7, 8}
// 		dp[i] = max(dp[i-1]+nums[i], nums[i])
// 		mx = max(mx, dp[i])
// 	}
// 	return mx
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func maxSubArray(nums []int) int {
// 	max, sum := nums[0], nums[0]
// 	for _, v := range nums[1:] {
// 		if sum < 0 {
// 			sum = v
// 		} else {
// 			sum += v
// 		}
// 		if max < sum {
// 			max = sum
// 		}
// 	}
// 	return max
// }

// func maxSubArray(nums []int) int {
// 	max := nums[0]
// 	count := 0

// 	for i := 0; i < len(nums); i++ {
// 		count += nums[i]
// 		if count > max {
// 			max = count
// 		}
// 		if count < 0 {
// 			count = 0
// 		}
// 	}
// 	return max
// }
func maxSubArray(nums []int) int {
	dt := make([]int, len(nums))
	dt[0] = nums[0]
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		dt[i] = max(dt[i-1]+nums[i], nums[i])
		mx = max(dt[i], mx)
	}
	return mx
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {

	// num := []int{5, 4, -1, 7, 8}
	num := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSubArray(num)
}
