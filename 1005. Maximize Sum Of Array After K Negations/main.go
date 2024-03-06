package main

import (
	"math"
	"sort"
)

// "strconv"

/*
Given an integer array nums and an integer k, modify the array in the following way:

choose an index i and replace nums[i] with -nums[i].
You should apply this process exactly k times.
You may choose the same index i multiple times.

Return the largest possible sum of the array after modifying it in this way.

Example 1:

Input: nums = [4,2,3], k = 1
Output: 5
Explanation: Choose index 1 and nums becomes [4,-2,3].
Example 2:
1-1 1
Input: nums = [3,-1,0,2], k = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].
Example 3:

Input: nums = [2,-3,-1,5,-4], k = 2
Output: 13
Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].
Constraints:

1 <= nums.length <= 104
-100 <= nums[i] <= 100
1 <= k <= 104
*/
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	indexZero := 0
	minvalue := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			indexZero = i
			break
		}
		if minvalue > -nums[i] {
			indexZero = i - 1
			break
		}
		minvalue = nums[i]
	}
	path := 0
	for k != 0 {
		if path == 0 && indexZero != 0 {
			for ; path <= indexZero; path++ {
				nums[path] = -nums[path]
				k--
				if k == 0 {
					break
				}
			}
		} else {
			sort.Ints(nums)
			nums[0] = -nums[0]
			k--
		}

	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}

func largestSumAfterKNegations(nums []int, k int) int {
	for k > 0 {
		minNum, idx := math.MaxInt, 0
		for i, num := range nums {
			if num < minNum {
				minNum = num
				idx = i
			}
		}
		if minNum == 0 {
			break
		}
		nums[idx] = -minNum
		k--
	}
	return sumNum(nums)
}

func sumNum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
func largestSumAfterKNegations(nums []int, K int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if K > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			K--
		}
	}

	if K%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}
func main() {
	// input := []int{0, 6, 9, -3, -1, 3, -8}
	input := []int{2, -3, -1, 5, -4}
	largestSumAfterKNegations(input, 2)
}
