package main

// "strconv"

/*
You are given an array of integers stones where stones[i]
is the weight of the ith stone.

We are playing a game with the stones. On each turn,
we choose any two stones and smash them together.
Suppose the stones have weights x and y with x <= y.

The result of this smash is:
If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed,
and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.

Return the smallest possible weight of the left stone.
If there are no stones left, return 0.
*/

// func lastStoneWeightII(stones []int) int {
// 	allwight := 0
// 	for _, num := range stones {
// 		allwight += num
// 	}
// 	target := allwight / 2
// 	dp := make([]int, target+1)
// 	for i := 0; i < len(stones); i++ {
// 		for j := target; j >= stones[i]; j-- {
// 			if dp[j] < dp[j-stones[i]]+stones[i] {
// 				dp[j] = dp[j-stones[i]] + stones[i]
// 			}
// 		}
// 	}

// 	return allwight - dp[target] - dp[target]
// }
func lastStoneWeightII(stones []int) int {
	allwight := 0
	for i := range stones {
		allwight += stones[i]
	}
	target := allwight / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j-stones[i]]+stones[i], dp[j])
		}
	}
	return allwight - dp[target] - dp[target]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	// out:=5
	// input := []int{31,26,33,21,40}
	//out =1
	// 	Explanation:
	// We can combine 2 and 4 to get 2, so the array converts to [2,7,1,8,1] then,
	// we can combine 7 and 8 to get 1, so the array converts to [2,1,1,1] then,
	// we can combine 2 and 1 to get 1, so the array converts to [1,1,1] then,
	// we can combine 1 and 1 to get 0, so the array converts to [1], then that's the optimal value.
	input := []int{2, 7, 4, 1, 8, 1}
	lastStoneWeightII(input)
}
