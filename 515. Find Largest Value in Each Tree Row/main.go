package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
Given the root of a binary tree,
return an array of the largest value in each row of the tree (0-indexed).

Input: root = [1,3,2,5,3,null,9]
Output: [1,3,9]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) (ans []int) {
	stack := []*TreeNode{root}
	const INT_MAX = int(^uint(0) >> 1)
	if root == nil {
		return
	}
	for len(stack) != 0 {
		maxVal := ^INT_MAX
		for _, i := range stack {
			stack = stack[1:]
			if i.Left != nil {
				stack = append(stack, i.Left)
			}
			if i.Right != nil {
				stack = append(stack, i.Right)
			}
			if i.Val > maxVal {
				maxVal = i.Val
			}
		}
		ans = append(ans, maxVal)
	}
	return
}

// func rightSideView(root *TreeNode) (ans []int) {
// 	stack := []*TreeNode{root}
// 	if root == nil {
// 		return
// 	}
// 	// ans = append(ans, root.Val)
// 	for len(stack) != 0 {
// 		// ans = append(ans, root.Val)
// 		tmp := []int{}
// 		for _, i := range stack {
// 			stack = stack[1:]
// 			if i.Left != nil {
// 				stack = append(stack, i.Left)
// 			}
// 			if i.Right != nil {
// 				stack = append(stack, i.Right)
// 			}
// 			tmp = append(tmp, i.Val)
// 		}
// 		ans = append(ans, tmp[len(tmp)-1])
// 	}
// 	return
// }

// func rightSideView(root *TreeNode) []int {
// 	var result []int
// 	populateRightSideView(root, &result, 0)
// 	return result
// }

// func populateRightSideView(root *TreeNode, result *[]int, level int) {
// 	if root == nil {
// 		return
// 	}
// 	if level == len(*result) {
// 		*result = append(*result, root.Val)
// 	}
// 	populateRightSideView(root.Right, result, level+1)
// 	populateRightSideView(root.Left, result, level+1)
// }
func main() {

	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 4.0,
	// }
	var Node3 *TreeNode = &TreeNode{
		Val: 3.0,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: -1.0,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}

	largestValues(Node1)
}
