package main

import (
	"fmt"
	// "strconv"
)

/*
Given the root of a Binary Search Tree (BST),
convert it to a Greater Tree such that every key of
the original BST is changed to the original key plus
the sum of all keys greater than the original key in BST.

As a reminder, a binary search tree is a tree that satisfies these constraints:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func convertBST(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	rightValue := 0
// 	right := convertBST(root.Right)
// 	if right != nil {
// 		rightValue = right.Val
// 	}
// 	root.Val = root.Val + rightValue
// 	if root.Left != nil {
// 		root.Left.Val = root.Left.Val + root.Val
// 	}

// 	left := convertBST(root.Left)
// 	// return left

// 	return root

// }

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var travel func(*TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Right)
		root.Val = root.Val + sum
		sum = root.Val
		travel(root.Left)
	}
	travel(root)
	return root
}

func main() {

	// root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
	var Node9 *TreeNode = &TreeNode{
		Val: 8,
	}
	var Node8 *TreeNode = &TreeNode{
		Val: 3,
	}
	var Node7 *TreeNode = &TreeNode{
		Val: 7, Right: Node9,
	}
	var Node6 *TreeNode = &TreeNode{
		Val: 5,
	}
	var Node5 *TreeNode = &TreeNode{
		Val: 2, Right: Node8,
	}
	var Node4 *TreeNode = &TreeNode{
		Val: 0,
	}
	var Node3 *TreeNode = &TreeNode{
		Val: 6, Left: Node6, Right: Node7,
	}
	var Node2 *TreeNode = &TreeNode{
		Val: 1, Left: Node4, Right: Node5,
	}
	var Node1 *TreeNode = &TreeNode{
		Val: 4, Left: Node2, Right: Node3,
	}
	// [50,30,70,null,40,60,80]
	p := convertBST(Node1)
	fmt.Println(p)
}
