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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	// swap(root.Left, root.Right)
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

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
		Val: 7.0,
	}
	var Node6 *TreeNode = &TreeNode{
		Val: 6.0,
	}
	var Node5 *TreeNode = &TreeNode{
		Val: 5.0,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 4.0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 3.0, Left: Node6, Right: Node7,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Left: Node4, Right: Node5,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2, Right: Node3,
	}

	invertTree(Node1)
}
