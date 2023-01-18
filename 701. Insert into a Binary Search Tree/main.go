package main

import (
	"fmt"
	// "strconv"
)

/*
You are given the root node of a binary search tree (BST)
and a value to insert into the tree.
Return the root node of the BST after the insertion.
It is guaranteed that the new value does not exist in the original BST.

Notice that there may exist multiple valid ways for the insertion,
as long as the tree remains a BST after insertion.
You can return any of them.

*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val && root.Right == nil {
		root.Right = insertIntoBST(root.Right, val)
	} else if root.Val > val && root.Left == nil {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		if root.Val < val {
			root.Right = insertIntoBST(root.Right, val)
		} else {
			root.Left = insertIntoBST(root.Left, val)
		}
	}
	return root
}
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
func main() {

	node5 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 7}
	node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	node1 := &TreeNode{Val: 4, Left: node2, Right: node3}

	p := insertIntoBST(node1, 5)
	fmt.Println(p)
}
