package main

// "fmt"
// "strconv"

/*

Find all valid combinations of k numbers
that sum up to n such
that the following conditions are true:

Only numbers 1 through 9 are used.
Each number is used at most once.
Return a list of all possible valid combinations.
The list must not contain the same combination twice,
and the combinations may be returned in any order.
Example 1:

Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.
Example 2:

Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.
Example 3:

Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9],
the smallest sum we can get is 1+2+3+4 = 10 and
since 10 > 1, there are no valid combination.
*/

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	var comb func(curComb []int, sum int)
	comb = func(curComb []int, sum int) {
		if len(curComb) > k {
			return
		}
		if len(curComb) == k && sum == 0 {
			dst := make([]int, k)
			copy(dst, curComb)
			result = append(result, dst)
		}
		start := 1
		if len(curComb) != 0 {
			start = curComb[len(curComb)-1] + 1
		}
		for i := start; i <= n; i++ {
			if sum < 0 || i > 9 {
				return
			}
			curComb = append(curComb, i)
			comb(curComb, sum-i)
			curComb = curComb[:len(curComb)-1]
		}
	}
	comb(make([]int, 0), n)
	return result
}
func combinationSum3(k int, n int) [][]int {
	var result [][]int

	combination, visit := make([]int, k), make([]bool, 10)

	var backtrack func(int, int)
	backtrack = func(index int, target int) {
		if target < 0 {
			return
		}

		if index == k {
			if target == 0 {
				copiedCombination := make([]int, k)
				copy(copiedCombination, combination)
				result = append(result, copiedCombination)
			}

			return
		}

		begin := 1

		if index > 0 {
			begin = combination[index-1] + 1
		}

		for num := begin; num <= 9; num++ {
			if visit[num] == false {
				visit[num] = true
				combination[index] = num
				backtrack(index+1, target-num)
				visit[num] = false
			}
		}
	}

	backtrack(0, n)

	return result
}
func main() {

	// num := []int{1, 2, 3, 1}
	combinationSum3(2, 18)
}
