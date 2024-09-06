package main

// "fmt"
// "strconv"

/*
easy
two pointer
done

Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
The relative order of the elements may be changed.

Since it is impossible to change the length of the array in some languages,
you must instead have the result be placed in the first part of the array nums. More formally,
if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result.
It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).
Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).
*/
// func removeElement(nums []int, val int) int {
// 	l, r := 0, len(nums)-1

// 	for l <= r {
// 		if nums[l] == val {
// 			nums[l], nums[r] = nums[r], nums[l]
// 			r--
// 		} else {
// 			l++
// 		}
// 	}
// 	return l
// }

// func removeElement(nums []int, val int) int {
// 	var start = 0
// 	var end = len(nums) - 1
// 	for start <= end {
// 		for end > -1 && nums[end] == val {
// 			end--
// 		}

// 		if end > start && nums[start] == val {
// 			nums[start], nums[end] = nums[end], nums[start]
// 		}
// 		start++
// 	}

// 	return end + 1
// }
func removeElement(nums []int, val int) int {

	size := len(nums)

	// two-pointers
	// otherIdx: index for other elements
	// targetIdx: index for target elements with val

	otherIdx, targetIdx := 0, size-1

	for otherIdx <= targetIdx {

		if nums[otherIdx] == val {

			// swap target elements to the tail side
			nums[otherIdx], nums[targetIdx] = nums[targetIdx], nums[otherIdx]

			targetIdx -= 1
		} else {

			// increase the index when we meet others
			otherIdx += 1
		}

	}

	// length of others is the answer
	return otherIdx

}

// func removeElement(nums []int, val int) int {
// 	if len(nums) == 1 {
// 		if nums[0] == val {
// 			return 0
// 		} else {
// 			return 1
// 		}
// 	}
// 	l, r := 0, len(nums)-1
// 	for l < r {
// 		for l < len(nums) && nums[l] != val {
// 			l++
// 		}
// 		for r > 0 && nums[r] == val {
// 			r--
// 		}
// 		if l < r && nums[l] == val {
// 			nums[l], nums[r] = nums[r], nums[l]
// 		}
// 	}
// 	return l
// }

func main() {
	// hay := "mississippi"
	// nee := "issipi"
	// var num = []int{1}
	// var num = []int{4, 5}
	// var num = []int{0, 1, 2, 2, 3, 0, 4, 2}
	var num = []int{3, 2, 2, 3}
	val := 3
	// words := ["bella","label","roller"]
	removeElement(num, val)
}
