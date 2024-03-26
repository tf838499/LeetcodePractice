package main

import "fmt"

// "fmt"
// "strconv"
// "strings"

/*
Given an integer n,
return the number of structurally unique BST's
(binary search trees)
which has exactly n nodes of unique values from 1 to n.
*/
// 3
// func numTrees(n int) int {
// 	// n=2
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	// 2 2
// 	return numTrees(n-1) + numTrees(n-1) + 1
// }
func numTrees(n int) int {
	dp := make([]int, n+1)    // 创建一个长度为n+1的数组，初始化为0
	dp[0] = 1                 // 当n为0时，只有一种情况，即空树，所以dp[0] = 1
	for i := 1; i <= n; i++ { // 遍历从1到n的每个数字
		for j := 1; j <= i; j++ { // 对于每个数字i，计算以i为根节点的二叉搜索树的数量
			dp[i] += dp[j-1] * dp[i-j] //返回以1到n为节点的二叉搜索树的总数量

		}
	}
	return dp[n]
}

func main() {
	P := numTrees(4)
	fmt.Println(P)
}
