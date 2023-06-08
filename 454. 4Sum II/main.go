package main

// "strconv"

/*
Medium

doing

Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the number of tuples (i, j, k, l) such that:

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
Output: 2
Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for _, i := range nums1 {
		for _, j := range nums2 {
			m[i+j]++

		}
	}
	ans := 0
	for _, i := range nums3 {
		for _, j := range nums4 {

			_, ok := m[-(i + j)]
			if ok {
				ans++
			}
		}
	}
	return ans
}

// func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
// 	oneTwoSums := map[int]int{}
// 	for _, one := range nums1 {
// 		for _, two := range nums2 {
// 			oneTwoSums[one+two]++
// 		}
// 	}

// 	tuples := 0
// 	for _, three := range nums3 {
// 		for _, four := range nums4 {
// 			tuples += oneTwoSums[-(three + four)]
// 		}
// 	}

//		return tuples
//	}
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fourSumCount(nums1, nums2, nums3, nums4)
}
