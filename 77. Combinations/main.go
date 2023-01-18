package main

import (
// "fmt"
// "strconv"
)

/*
Given two integers n and k,
return all possible combinations of k numbers chosen from the range [1, n].

You may return the answer in any order.

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered,
i.e., [1,2] and [2,1] are considered to be the same combination.
*/

func combine(n int, k int) [][]int {

	result := make([][]int, 0)

	//-------------------------------------------
	var comb func(start int, curComb []int)

	comb = func(start int, curComb []int) {

		// Base case
		if len(curComb) == k {
			// make a copy of current combination
			dst := make([]int, k)
			copy(dst, curComb)
			result = append(result, dst)
			return
		}

		// General cases:
		for i := start; i <= n; i++ {
			curComb = append(curComb, i) // 1.取1 3.取2, 5.取3
			comb(i+1, curComb)           // 2.開始深度尋找 2,3,4  4.第二次會變3,4 6.變4
			curComb = curComb[:len(curComb)-1]
		}
		return
	}
	//-------------------------------------------
	comb(1, make([]int, 0))
	return result
}

func main() {
	// num := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	combine(4, 2)
}
