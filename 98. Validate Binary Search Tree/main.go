package main

import (
	"fmt"
	"math"
	// "strconv"
	// "strings"
)

/*
Given the root of a binary tree,
determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// var prev *TreeNode

// func isValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	res := true
// 	if root.Left != nil {
// 		res = root.Val > root.Left.Val && isValidBST(root.Left)
// 	}
// 	if prev != nil && prev.Val >= root.Val {
// 		return false
// 	}
// 	prev = root
// 	if root.Right != nil && res {
// 		res = root.Val < root.Right.Val && isValidBST(root.Right)
// 	}
// 	return res
// }

func isValidBST(root *TreeNode) bool {

	return validate(root, math.MinInt, math.MaxInt)
}
func validate(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val > lower && root.Val < upper {
		return validate(root.Left, lower, root.Val) && validate(root.Right, root.Val, upper)
	} else {
		return false
	}
}

// func validate(node *TreeNode, lower int, upper int) bool {

// 	if node == nil {

// 		//empty node or empty tree is always valid
// 		return true
// 	}

// 	if (lower < node.Val) && (node.Val < upper) {

// 		// check if all tree nodes follow BST rule
// 		return validate(node.Left, lower, node.Val) && validate(node.Right, node.Val, upper)
// 	} else {

// 		// early reject when we find violation
// 		return false
// 	}

// }
func main() {
	// var Node6 *TreeNode = &TreeNode{
	// 	Val: 27,
	// }
	// var Node5 *TreeNode = &TreeNode{
	// 	Val: 56,
	// }
	// var Node4 *TreeNode = &TreeNode{
	// 	Val: 19, Right: Node6,
	// }
	// var Node3 *TreeNode = &TreeNode{
	// 	Val: 47, Right: Node5,
	// }
	// var Node2 *TreeNode = &TreeNode{
	// 	Val: 26, Left: Node4,
	// }
	// var Node1 *TreeNode = &TreeNode{
	// 	Val: 32, Right: Node3, Left: Node2,
	// }
	var Node1 *TreeNode = &TreeNode{
		Val: 0,
	}
	p := isValidBST(Node1)
	fmt.Println(p)
}

// [32,26,47,19,null,null,56,null,27]
// 	 	32
//   26	   47
// 19  n  n  56
//n 27
