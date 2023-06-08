package main

import "sort"

// "strconv"

/*
Medium
hash map
doing

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
// func threeSum(nums []int) [][]int {
// 	ans := [][]int{}
// 	sort.Ints(nums)

// 	for index, value := range nums {

// 		l, r := index+1, len(nums)-1
// 		target := value
// 		for l < r {
// 			n2, n3 := nums[l], nums[r]
// 			sum := n2 + n3
// 			if target+sum == 0 {
// 				ans = append(ans, []int{value, nums[l], nums[r]})
// 				l++
// 				r--
// 				for l < r && nums[l] == nums[l-1] {
// 					l++
// 				}
// 				for l < r && nums[r] == nums[r+1] {
// 					r--
// 				}
// 			} else if target+sum < 0 {
// 				l++
// 			} else if target+sum > 0 {
// 				r--
// 			}
// 		}
// 	}
// 	return ans
// }

func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue //To prevent the repeat
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return results
}

//	func threeSum(nums []int) [][]int {
//		sort.Ints(nums)
//		res := [][]int{}
//		// 找出a + b + c = 0
//		// a = nums[i], b = nums[left], c = nums[right]
//		for i := 0; i < len(nums)-2; i++ {
//			// 排序之后如果第一个元素已经大于零，那么无论如何组合都不可能凑成三元组，直接返回结果就可以了
//			n1 := nums[i]
//			if n1 > 0 {
//				break
//			}
//			// 去重a
//			if i > 0 && n1 == nums[i-1] {
//				continue
//			}
//			l, r := i+1, len(nums)-1
//			for l < r {
//				n2, n3 := nums[l], nums[r]
//				if n1+n2+n3 == 0 {
//					res = append(res, []int{n1, n2, n3})
//					// 去重逻辑应该放在找到一个三元组之后，对b 和 c去重
//					for l < r && nums[l] == n2 {
//						l++
//					}
//					for l < r && nums[r] == n3 {
//						r--
//					}
//				} else if n1+n2+n3 < 0 {
//					l++
//				} else {
//					r--
//				}
//			}
//		}
//		return res
//	}
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	threeSum(nums)

}
