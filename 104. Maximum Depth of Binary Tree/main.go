package main

import (
// "fmt"
// "strconv"
// "strings"
// "math"
)

/*
easy
recursion
dfs
done


Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* recursive */
func maxDepth(root *TreeNode) int {
	// ans := 0
	if root == nil {
		return 0
	}
	// ans = ans + 1
	n1 := maxDepth(root.Left)
	n2 := maxDepth(root.Right)

	if n1 > n2 {
		return n1 + 1
	}
	return n2 + 1
}

/* stack */
// func maxDepth(root *TreeNode) int {
// 	ans := 0
// 	if root == nil {
// 		return ans
// 	}
// 	stack := []*TreeNode{root}

// 	for len(stack) != 0 {
// 		for _, p := range stack {
// 			stack = stack[1:]
// 			if p.Right != nil {
// 				stack = append(stack, p.Right)
// 			}
// 			if p.Left != nil {
// 				stack = append(stack, p.Left)
// 			}
// 		}
// 		ans++
// 	}
// 	return ans
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

	maxDepth(Node1)
}
