package main

// "sort"
// "strconv"

/*
You are given an integer array nums and an integer target.

You want to build an expression out of nums
by adding one of the symbols '+' and '-'
before each integer in nums and then concatenate
all the integers.

For example, if nums = [2, 1],
you can add a '+' before 2 and a '-'
before 1 and concatenate them to build the expression "+2-1".
Return the number of different expressions
that you can build, which evaluates to target.
*/

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if target > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	left := (sum + target) / 2
	dp := make([]int, left+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := left; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[left]
}

// func findTargetSumWays(nums []int, target int) int {
// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	if abs(target) > sum {
// 		return 0
// 	}
// 	if (sum+target)%2 == 1 {
// 		return 0
// 	}
// 	// 计算背包大小
// 	bag := (sum + target) / 2
// 	// 定义dp数组
// 	dp := make([]int, bag+1)
// 	// 初始化
// 	dp[0] = 1
// 	// 遍历顺序
// 	for i := 0; i < len(nums); i++ {
// 		for j := bag; j >= nums[i]; j-- {
// 			//推导公式
// 			dp[j] += dp[j-nums[i]]
// 			//fmt.Println(dp)
// 		}
// 	}
// 	return dp[bag]
// }

// func abs(x int) int {
// 	return int(math.Abs(float64(x)))
// }
func main() {
	// [1,2,3,4,5,6,7,8,9,10,1,1,1,1,1]
	input := []int{1, 1, 1, 1, 1}

	// input := []int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}
	target := 3
	findTargetSumWays(input, target)
}
