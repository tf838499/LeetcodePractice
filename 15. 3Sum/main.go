package main

import (
	"fmt"
	// "strconv"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
*/
func threeSum(nums []int) [][]int {

	var target int
	var ans [][]int
	// scores := make([]int, 0)
	// tmp := false
	// set := make(map[int]bool)
	// ans := make([]int, 0)

	// Sets = n
	for _, p := range nums {
		n := make(map[int]int)
		for _, p := range nums {
			n[p]++
		}
		scores := make([]int, 0)
		dis := target - p
		n[p]--
		scores = append(scores, p)
		reduce := 0
		for dis != 0 || reduce != 0 || len(scores) < 3 {
			if dis == 0 && reduce != 0 {
				dis = dis + (-reduce)
				reduce--
			}
			if n[dis] >= 0 {
				scores = append(scores, dis)
				n[dis]--
				dis = dis - dis
			} else if dis < 0 {
				reduce++
				dis = dis + reduce
			} else if dis > 0 {
				reduce++
				dis = dis - reduce
			}

			fmt.Println(scores)
		}
		ans = append(ans, scores)

	}

	// // fmt.Println(Sets)
	// ans := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// }
	return ans
}

func threeSum(nums []int) [][]int {
	n := len(nums)

	// Sort the given array
	sort.Ints(nums)

	var result [][]int
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		// Skip all duplicates from left
		// num1Idx>0 ensures this check is made only from 2nd element onwards
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1
		num3Idx := n - 1
		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})

				num3Idx--

				// Skip all duplicates from right
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				// Decrement num3Idx to reduce sum value
				num3Idx--
			} else {
				// Increment num2Idx to increase sum value
				num2Idx++
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 2, 1, 0, -1, -4}
	// words := ["bella","label","roller"]
	threeSum(nums)

}
