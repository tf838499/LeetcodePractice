package main

// "strconv"

/*
A wiggle sequence is a sequence where the differences between
successive numbers strictly alternate between
positive and negative.
The first difference (if one exists) may be
either positive or negative.
A sequence with one element and a sequence with
two non-equal elements are trivially wiggle sequences.

For example, [1, 7, 4, 9, 2, 5] is a wiggle sequence because
the differences (6, -3, 5, -7, 3) alternate between
positive and negative.
In contrast,
[1, 4, 7, 2, 5] and [1, 7, 4, 5, 5] are not wiggle sequences.
The first is not because its first two differences are positive,
and the second is not because its last difference is zero.
A subsequence is obtained by deleting some elements (possibly zero)
from the original sequence,
leaving the remaining elements in their original order.

Given an integer array nums,
return the length of the longest wiggle subsequence of nums.
*/

/*
Input: nums = [1,7,4,9,2,5]
Output: 6
Explanation:
The entire sequence is a wiggle sequence with
differences (6, -3, 5, -7, 3)
1 -12 5 3 2 -5
Input: nums = [1,17,5,10,13,15,10,5,16,8]
Output: 7
Explanation:
There are several subsequences that achieve this length.
One is [1, 17, 10, 13, 10, 16, 8] with
differences (16, -7, 3, -3, 6, -8).
*/

// func wiggleMaxLength(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	result := 1
// 	// tmp := []int{}
// 	prediff := 0
// 	for i := 1; i < len(nums); i++ {
// 		diff := nums[i] - nums[i-1]
// 		if i == 1 || prediff == 0 {
// 			prediff = diff
// 			if prediff != 0 {
// 				result++
// 			}
// 			// result++
// 			continue
// 		}

// 		if diff > 0 && prediff < 0 || diff < 0 && prediff > 0 {
// 			result++
// 			prediff = diff
// 		}

// 	}
// 	return result
// }
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func wiggleMaxLength(nums []int) int {

	n := len(nums)

	// Base case:
	// length of wiggle sequence, ending in positive difference, negative difference
	positive, negative := 1, 1

	// General cases:
	// scan from second number to last number
	for i := 1; i < n; i++ {

		if nums[i] > nums[i-1] {

			// difference is positive, concatenated with negative prefix wiggle subsequence
			positive = negative + 1

		} else if nums[i] < nums[i-1] {

			// difference is negative, concatenated with positive prefix wiggle subsequence
			negative = positive + 1
		}

	}

	//  compute the longest length of wiggle sequence
	return max(positive, negative)
}
func main() {
	wiggleMaxLength([]int{1, 2, 3, 4, 5, 2})
}
