package main

import "fmt"

// "fmt"s
// "strconv"

/*
You are given an integer array cost where cost[i]
is the cost of ith step on a staircase.
Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0,
or the step with index 1.

Return the minimum cost to reach the top of the floor.

Example 1:

Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.
Example 2:

Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.
*/
// func minCostClimbingStairs(cost []int) int {
// 	coststep := make([]int, len(cost))
// 	coststep[0] = cost[0]
// 	coststep[1] = cost[1]
// 	for i := 2; i < len(cost); i++ {
// 		if coststep[i-1] < coststep[i-2] {
// 			coststep[i] = coststep[i-1] + cost[i]
// 		} else {
// 			coststep[i] = coststep[i-2] + cost[i]
// 		}
// 	}
// 	if coststep[len(coststep)-1] < coststep[len(coststep)-2] {
// 		return coststep[len(coststep)-1]
// 	} else {
// 		return coststep[len(coststep)-2]
// 	}
// }
func minCostClimbingStairs(cost []int) int {
	// coststep := make([]int, len(cost))
	coststep1 := 0
	coststep2 := 0
	coststepi := 0
	for i := 2; i <= len(cost); i++ {
		if coststep2+cost[i-1] < coststep1+cost[i-2] {
			coststepi = coststep2 + cost[i-1]
		} else {
			coststepi = coststep1 + cost[i-2]
		}
		coststep1 = coststep2
		coststep2 = coststepi
	}
	return coststep2
}
func main() {
	// input := []int{10, 15, 20}
	input := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	P := minCostClimbingStairs(input)
	fmt.Println(P)
}
