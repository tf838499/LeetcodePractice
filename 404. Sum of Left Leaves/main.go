package main

import (
	"fmt"
	// "strconv"
	// "math"
)

/*
Given the root of a binary tree,
return the sum of all left leaves.

A leaf is a node with no children.
A left leaf is a leaf that is the left child of another node.

Input: root = [3,9,20,null,null,15,7]
Output: 24
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func sumOfLeftLeaves(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }

//     if root.Left == nil {
//         return sumOfLeftLeaves(root.Right)
//     }

//     if root.Left.Left == nil && root.Left.Right == nil {
//         return root.Left.Val + sumOfLeftLeaves(root.Right)
//     }

//     return sumOfLeftLeaves(root.Left)+ sumOfLeftLeaves(root.Right)
// }
func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if root.Left == nil {
		return sumOfLeftLeaves(root.Right)
	}
	if root.Left.Left == nil && root.Left.Right == nil {
		return sumOfLeftLeaves(root.Right) + root.Left.Val
		// return ans + root.Left.Val
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
func main() {
	node6 := &TreeNode{Val: 10}
	node5 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 15}
	node3 := &TreeNode{Val: 20, Left: node4, Right: node5}
	node2 := &TreeNode{Val: 9, Left: node6}
	node1 := &TreeNode{Val: 3, Left: node2, Right: node3}
	// a := []int{1, 1, 1, 1, 1, 1, 1, 1}
	p := sumOfLeftLeaves(node1)
	fmt.Println(p)
}
