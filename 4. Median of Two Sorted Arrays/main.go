package main

import (
	// "fmt"
	// "sort"
	"math"
)

// "strconv"

/*


Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

// 	l, r := 0, 0
// 	l_len := len(nums1)
// 	r_len := len(nums2)
// 	total_len := l_len + r_len
// 	res := []int{}
// 	ans := 0.0
// 	for l+r < total_len/2 {
// 		if nums1[l] > nums2[r] {
// 			res = append(res, nums2[r])
// 			r++
// 		} else if l < l_len {
// 			res = append(res, nums1[l])
// 			l++
// 		}
// 	}
// 	if len(res)%2 != 0 {
// 		ans = float64(res[len(res)-1])
// 	} else {
// 		ans = float64(res[len(res)-1]+res[len(res)-2]) / 2
// 	}

// 	return ans
// }
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen == 0 {
		return float64(totalLen)
	}
	halfItem := int(math.Ceil(float64(totalLen) / 2.0))
	multi := totalLen%2 == 0
	maxLen := len(nums1)
	if len(nums2) > maxLen {
		maxLen = len(nums2)
	}
	var item, prevItem, counter, i, j int
	if multi {
		halfItem += 1
	}
	for counter < halfItem {
		prevItem = item
		if (i < len(nums1) && j == len(nums2)) || (i < len(nums1) && j < len(nums2) && nums1[i] < nums2[j]) {
			item = nums1[i]
			i++
		} else {
			item = nums2[j]
			j++
		}
		counter++
	}
	if multi {
		return (float64(item) + float64(prevItem)) / 2
	}
	return float64(item)
}
func main() {

	n1 := []int{1, 2}
	n2 := []int{3, 4}
	// pwwkew
	findMedianSortedArrays(n1, n2)
	// fmt.Println(l1)

}
