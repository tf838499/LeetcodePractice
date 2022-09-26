package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) (ans bool) {
	if root == nil {
		return true
	}
	depth := 0
	// NoChild := 0
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		for _, p := range stack {
			stack = stack[1:]
			if p.Left != nil {
				stack = append(stack, p.Left)
			} else if depth < 2 {
				depth = 0
				depth++
			}
			if p.Right != nil {
				stack = append(stack, p.Right)
			} else {
				depth++
			}
			// return true
		}
		if depth != 0 {
			for _, p := range stack {
				stack = stack[1:]
				if p.Left != nil {
					depth++
				}
				if p.Right != nil {
					depth++
				}
			}
			if depth >= 2 {
				return false
			} else {
				return true
			}
		}
	}
	return ans
}

// func countNodes(root *TreeNode) (ans int) {
// 	if root == nil {
// 		return 0
// 	}
// 	ans = 1 + countNodes(root.Right) + countNodes(root.Left)
// 	return ans
// }

func main() {
	// var Node7 *TreeNode = &TreeNode{
	// 	Val: 7.0,
	// }
	// var Node6 *TreeNode = &TreeNode{
	// 	Val: 6.0,
	// }
	// var Node5 *TreeNode = &TreeNode{
	// 	Val: 5.0,
	// }
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 4.0, Right: Node5,
	// }
	var Node3 *TreeNode = &TreeNode{
		Val: 3.0,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 2.0, Left: Node3,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 1.0, Left: Node2,
	}

	isBalanced(Node1)
}

// [1,2,3,4,5,6,null,8]
