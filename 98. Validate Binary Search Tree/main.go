package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

/*
Given the root of a binary tree,
determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left
subtree
of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res, v := true, root.Val
	doValidateBST(&res, root.Left, &v, nil)
	doValidateBST(&res, root.Right, nil, &v)
	return res
}

func doValidateBST(res *bool, node *TreeNode, upperBound, lowerBound *int) {
	if node == nil || !*res {
		return
	}
	if (upperBound != nil && node.Val >= *upperBound) ||
		(lowerBound != nil && node.Val <= *lowerBound) {
		*res = false
		return
	}

	v := node.Val
	doValidateBST(res, node.Left, &v, lowerBound)
	doValidateBST(res, node.Right, upperBound, &v)
}
func main() {
	var Node5 *TreeNode = &TreeNode{
		Val: 6,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 3,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 6,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 4, Right: Node5, Left: Node4,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 5, Right: Node3, Left: Node2,
	}
	p := isValidBST(Node1)
	fmt.Println(p)
}
