package main

import (
	"fmt"
	// "strconv"
)

/*
Given a root node reference of a BST and a key,
delete the node with the given key in the BST.
Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.
*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		var Adjust func(node *TreeNode)
		Adjust = func(root *TreeNode) {
			if root == nil {
				return
			}
			if root.Right == nil {
				root.Right = root.Left
				root.Left = nil
				// if root != nil {
				// 	root.Left = nil
				// }
				return
			}
			tmpa := root.Left
			root = root.Right
			Adjust(root.Right)
			root.Left = tmpa
		}
		tmp := root.Left

		Adjust(root.Right)
		// root
		root = root.Right
		if root != nil {
			root.Left = tmp
		}
		// root.Left = tmp
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

// func deleteNode(root *TreeNode, key int) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	if key < root.Val {
// 		root.Left = deleteNode(root.Left, key)
// 		return root
// 	}
// 	if key > root.Val {
// 		root.Right = deleteNode(root.Right, key)
// 		return root
// 	}
// 	if root.Right == nil {
// 		return root.Left
// 	}
// 	if root.Left == nil {
// 		return root.Right
// 	}
// 	minnode := root.Right
// 	for minnode.Left != nil {
// 		minnode = minnode.Left
// 	}
// 	root.Val = minnode.Val
// 	root.Right = deleteNode1(root.Right)
// 	return root
// }

func deleteNode1(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}
func main() {
	// node6 := &TreeNode{Val: 7}
	// node5 := &TreeNode{Val: 4}
	// node4 := &TreeNode{Val: 2}
	// node3 := &TreeNode{Val: 6, Right: node6}
	// node2 := &TreeNode{Val: 3, Left: node4, Right: node5}
	// node1 := &TreeNode{Val: 5, Left: node2, Right: node3}
	// node1 := &TreeNode{Val: 0}
	node7 := &TreeNode{Val: 80}
	node6 := &TreeNode{Val: 60}
	node5 := &TreeNode{Val: 40}
	// node4 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 70, Left: node6, Right: node7}
	node2 := &TreeNode{Val: 30, Right: node5}
	node1 := &TreeNode{Val: 50, Left: node2, Right: node3}
	// [50,30,70,null,40,60,80]
	p := deleteNode(node1, 50)
	fmt.Println(p)
}
