package main

import (
// "fmt"
// "strconv"
)

/*
You are given two integer arrays nums1 and nums2,
sorted in non-decreasing order, and two integers m and n,
representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function,
but instead be stored inside the array nums1. To accommodate this,
nums1 has a length of m + n,
where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.


Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.


*/
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	slow, fast, nums_pin := 0, 0, 0
// 	for fast < m+n && slow < m+n {
// 		if nums1[fast] != 0 {
// 			// for
// 			for fast != 0 && nums1[fast] < nums1[fast-1] {
// 				nums1[fast], nums1[slow] = nums1[slow], nums1[fast]
// 				slow++
// 			}
// 			fast++
// 		} else if nums1[slow] >= nums2[nums_pin] {
// 			nums1[slow], nums1[fast] = nums1[fast], nums1[slow]
// 			nums1[slow], nums2[nums_pin] = nums2[nums_pin], nums1[slow]
// 			slow++
// 			nums_pin++
// 		} else if nums1[slow] < nums2[nums_pin] {
// 			slow++
// 			if nums1[slow] == 0 {
// 				nums1[slow], nums2[nums_pin] = nums2[nums_pin], nums1[slow]
// 				nums_pin++
// 				fast++
// 			}
// 			// slow++

// 		}

// 	}
// 	fmt.Println("aa")
// }
func merge(nums1 []int, m int, nums2 []int, n int) {
	point := m + n - 1
	m--
	n--
	for n >= 0 || m >= 0 {
		if n < 0 {
			nums1[point] = nums1[m]
			m--
		} else if m < 0 {
			nums1[point] = nums2[n]
			n--
		} else if nums1[m] > nums2[n] {
			nums1[point] = nums1[m]
			m--
		} else {
			nums1[point] = nums2[n]
			n--
		}
		point--
	}

}
func merge(nums1 []int, m int, nums2 []int, n int) {
	cur := n + m - 1
	cur1, cur2 := m-1, n-1

	for cur1 >= 0 || cur2 >= 0 {
		switch {
		case cur1 < 0:
			nums1[cur] = nums2[cur2]
			cur--
			cur2--
		case cur2 < 0:
			nums1[cur] = nums1[cur1]
			cur--
			cur1--
		default:
			if nums1[cur1] > nums2[cur2] {
				nums1[cur] = nums1[cur1]
				cur1--
			} else {
				nums1[cur] = nums2[cur2]
				cur2--
			}
			cur--
		}
	}
}

// 456
// 123
// 0564
// 123
// 1564
// 023
// 10645
// 10
func main() {
	num1 := []int{4, 5, 6, 0, 0, 0}
	// num1 := []int{1, 2, 3, 0, 0, 0}
	// num1 := []int{0}
	num2 := []int{1, 2, 3}
	// num2 := []int{1}
	merge(num1, 3, num2, 3)

}
