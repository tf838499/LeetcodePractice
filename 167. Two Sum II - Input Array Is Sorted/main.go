package main

import (
	"fmt"
	// "strconv"
)

/*
day 2 筆記
除了hash map外，也可以用雙指針做更省記憶體空間
注意題目已經排序好 只需要指針所指相加，跟target比大小 決定左右指針誰要遞增誰要遞減

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
find two numbers such that they add up to a specific target number.
Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
*/
func twoSum(numbers []int, target int) []int {
	tmp := make(map[int]int)

	for index, val := range numbers {
		remainder := target - val
		if indexOfTmp, ok := tmp[remainder]; ok {
			return []int{indexOfTmp + 1, index + 1}
		}
		tmp[val] = index
	}
	return []int{-1, -1}
}
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	sum := numbers[i] + numbers[j]

	for sum != target {
		if sum < target {
			i++
		} else {
			j--
		}

		sum = numbers[i] + numbers[j]
	}

	return []int{i + 1, j + 1}
}
func main() {
	s := []int{2, 7, 11, 15}
	tar := 9
	p := twoSum(s, tar)
	fmt.Println(p)
}
