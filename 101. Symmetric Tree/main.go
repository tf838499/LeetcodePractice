package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
Easy
recursive
dfs
tree
done

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Input: root = [1,2,2,3,4,4,3]
Output: true
Input: root = [1,2,2,null,3,null,3]
Output: false
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) (ans bool) {
	if root == nil {
		return true
	}
	ans = compare(root.Left, root.Right)
	return ans
	// return
}
func compare(left *TreeNode, right *TreeNode) (ans bool) {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	ans = compare(left.Left, right.Right) && compare(left.Right, right.Left)
	return ans
}

// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	return checkSymmetry(root.Left, root.Right)
// }

// func checkSymmetry(n1 *TreeNode, n2 *TreeNode) bool {
// 	if n1 == nil || n2 == nil {
// 		return n1 == n2
// 	}

// 	if n1.Val != n2.Val {
// 		return false
// 	}

// 	return checkSymmetry(n1.Left, n2.Right) && checkSymmetry(n2.Left, n1.Right)
// }

// func largestValues(root *TreeNode) (ans []int) {
// 	stack := []*TreeNode{root}
// 	const INT_MAX = int(^uint(0) >> 1)
// 	if root == nil {
// 		return
// 	}
// 	for len(stack) != 0 {
// 		maxVal := ^INT_MAX
// 		for _, i := range stack {
// 			stack = stack[1:]
// 			if i.Left != nil {
// 				stack = append(stack, i.Left)
// 			}
// 			if i.Right != nil {
// 				stack = append(stack, i.Right)
// 			}
// 			if i.Val > maxVal {
// 				maxVal = i.Val
// 			}
// 		}
// 		ans = append(ans, maxVal)
// 	}
// 	return
// }

func main() {
	var Node7 *TreeNode = &TreeNode{
		Val: 3.0,
	}
	// var Node6 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	var Node5 *TreeNode = &TreeNode{
		Val: 4.0,
	}
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 3.0,
	// }
	var Node3 *TreeNode = &TreeNode{
		Val: 2.0, Right: Node7,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Right: Node5,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}

	isSymmetric(Node1)
}
