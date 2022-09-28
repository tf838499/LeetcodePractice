package main

import (
// "fmt"
// "strconv"
)

/*
easy
hashmap
done

Given two integer arrays nums1 and nums2,
return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/
func intersection(nums1 []int, nums2 []int) []int {
	numlen := len(nums1)
	var tmp []int
	ans := []int{}
	if numlen > len(nums2) {
		tmp = nums1
		nums1 = nums2
		nums2 = tmp
	}
	set := make(map[int]bool)
	for _, o := range nums2 {
		if set[o] == true {
			continue
		} else {
			set[o] = true
		}
	}
	for _, o := range nums1 {

		if set[o] == true {
			ans = append(ans, o)
			set[o] = false
		}
	}
	return ans
}

// func intersection(nums1 []int, nums2 []int) []int {
//     mem := make([]int, 1001)
//     for _, n := range nums1 {
//         if mem[n] == 0 {
//             mem[n]++
//         }
//     }
//     for _, n := range nums2 {
//         if mem[n] == 1 {
//             mem[n]++
//         }
//     }
//     result := []int{}
//     for i, sum := range mem {
//         if sum > 1 {
//             result = append(result, i)
//         }
//     }
//     return result
// }
func main() {
	numtest1_1 := []int{1, 2, 2, 1}
	numtest1_2 := []int{2, 2}

	// numtest2_1 := []int{4, 9, 5}
	// numtest2_2 := []int{9, 4, 9, 8, 4}
	// words := ["bella","label","roller"]
	intersection(numtest1_1, numtest1_2)

}
